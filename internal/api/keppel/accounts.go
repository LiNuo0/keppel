/******************************************************************************
*
*  Copyright 2018-2020 SAP SE
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
*
******************************************************************************/

package keppelv1

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sapcc/go-api-declarations/cadf"
	"github.com/sapcc/go-bits/audittools"
	"github.com/sapcc/go-bits/errext"
	"github.com/sapcc/go-bits/httpapi"
	"github.com/sapcc/go-bits/respondwith"
	"github.com/sapcc/go-bits/sqlext"

	"github.com/sapcc/keppel/internal/auth"
	peerclient "github.com/sapcc/keppel/internal/client/peer"
	"github.com/sapcc/keppel/internal/keppel"
	"github.com/sapcc/keppel/internal/models"
)

////////////////////////////////////////////////////////////////////////////////
// data types

// Account represents an account in the API.
type Account struct {
	Name              string                    `json:"name"`
	AuthTenantID      string                    `json:"auth_tenant_id"`
	GCPolicies        []keppel.GCPolicy         `json:"gc_policies,omitempty"`
	InMaintenance     bool                      `json:"in_maintenance"`
	Metadata          map[string]string         `json:"metadata"`
	RBACPolicies      []keppel.RBACPolicy       `json:"rbac_policies"`
	ReplicationPolicy *keppel.ReplicationPolicy `json:"replication,omitempty"`
	ValidationPolicy  *keppel.ValidationPolicy  `json:"validation,omitempty"`
	PlatformFilter    models.PlatformFilter     `json:"platform_filter,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// data conversion/validation functions

func (a *API) renderAccount(dbAccount models.Account) (Account, error) {
	gcPolicies, err := keppel.ParseGCPolicies(dbAccount)
	if err != nil {
		return Account{}, err
	}
	rbacPolicies, err := keppel.ParseRBACPolicies(dbAccount)
	if err != nil {
		return Account{}, err
	}
	if rbacPolicies == nil {
		// do not render "null" in this field
		rbacPolicies = []keppel.RBACPolicy{}
	}

	metadata := make(map[string]string)
	if dbAccount.MetadataJSON != "" {
		err := json.Unmarshal([]byte(dbAccount.MetadataJSON), &metadata)
		if err != nil {
			return Account{}, fmt.Errorf("malformed metadata JSON: %q", dbAccount.MetadataJSON)
		}
	}

	return Account{
		Name:              dbAccount.Name,
		AuthTenantID:      dbAccount.AuthTenantID,
		GCPolicies:        gcPolicies,
		InMaintenance:     dbAccount.InMaintenance,
		Metadata:          metadata,
		RBACPolicies:      rbacPolicies,
		ReplicationPolicy: renderReplicationPolicy(dbAccount),
		ValidationPolicy:  renderValidationPolicy(dbAccount),
		PlatformFilter:    dbAccount.PlatformFilter,
	}, nil
}

func renderReplicationPolicy(dbAccount models.Account) *keppel.ReplicationPolicy {
	if dbAccount.UpstreamPeerHostName != "" {
		return &keppel.ReplicationPolicy{
			Strategy:             "on_first_use",
			UpstreamPeerHostName: dbAccount.UpstreamPeerHostName,
		}
	}

	if dbAccount.ExternalPeerURL != "" {
		return &keppel.ReplicationPolicy{
			Strategy: "from_external_on_first_use",
			ExternalPeer: keppel.ReplicationExternalPeerSpec{
				URL:      dbAccount.ExternalPeerURL,
				UserName: dbAccount.ExternalPeerUserName,
				//NOTE: Password is omitted here for security reasons
			},
		}
	}

	return nil
}

func renderValidationPolicy(dbAccount models.Account) *keppel.ValidationPolicy {
	if dbAccount.RequiredLabels == "" {
		return nil
	}

	return &keppel.ValidationPolicy{
		RequiredLabels: dbAccount.SplitRequiredLabels(),
	}
}

////////////////////////////////////////////////////////////////////////////////
// handlers

func (a *API) handleGetAccounts(w http.ResponseWriter, r *http.Request) {
	httpapi.IdentifyEndpoint(r, "/keppel/v1/accounts")
	var accounts []models.Account
	_, err := a.db.Select(&accounts, "SELECT * FROM accounts ORDER BY name")
	if respondwith.ErrorText(w, err) {
		return
	}
	scopes := accountScopes(keppel.CanViewAccount, accounts...)

	authz := a.authenticateRequest(w, r, scopes)
	if authz == nil {
		return
	}
	if authz.UserIdentity.UserType() == keppel.AnonymousUser {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// restrict accounts to those visible in the current scope
	var accountsFiltered []models.Account
	for idx, account := range accounts {
		if authz.ScopeSet.Contains(*scopes[idx]) {
			accountsFiltered = append(accountsFiltered, account)
		}
	}
	// ensure that this serializes as a list, not as null
	if len(accountsFiltered) == 0 {
		accountsFiltered = []models.Account{}
	}

	// render accounts to JSON
	accountsRendered := make([]Account, len(accountsFiltered))
	for idx, account := range accountsFiltered {
		accountsRendered[idx], err = a.renderAccount(account)
		if respondwith.ErrorText(w, err) {
			return
		}
	}
	respondwith.JSON(w, http.StatusOK, map[string]any{"accounts": accountsRendered})
}

func (a *API) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	httpapi.IdentifyEndpoint(r, "/keppel/v1/accounts/:account")
	authz := a.authenticateRequest(w, r, accountScopeFromRequest(r, keppel.CanViewAccount))
	if authz == nil {
		return
	}
	account := a.findAccountFromRequest(w, r, authz)
	if account == nil {
		return
	}

	accountRendered, err := a.renderAccount(*account)
	if respondwith.ErrorText(w, err) {
		return
	}
	respondwith.JSON(w, http.StatusOK, map[string]any{"account": accountRendered})
}

var looksLikeAPIVersionRx = regexp.MustCompile(`^v[0-9][1-9]*$`)

func (a *API) handlePutAccount(w http.ResponseWriter, r *http.Request) {
	httpapi.IdentifyEndpoint(r, "/keppel/v1/accounts/:account")
	// decode request body
	var req struct {
		Account Account `json:"account"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "request body is not valid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	// we do not allow to set name in the request body ...
	if req.Account.Name != "" {
		http.Error(w, `malformed attribute "account.name" in request body is no allowed here`, http.StatusUnprocessableEntity)
		return
	}
	// ... transfer it here into the struct, to make the below code simpler
	req.Account.Name = mux.Vars(r)["account"]

	if err := a.authDriver.ValidateTenantID(req.Account.AuthTenantID); err != nil {
		http.Error(w, `malformed attribute "account.auth_tenant_id" in request body: `+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// reserve identifiers for internal pseudo-accounts and anything that might
	// appear like the first path element of a legal endpoint path on any of our
	// APIs (we will soon start recognizing image-like URLs such as
	// keppel.example.org/account/repo and offer redirection to a suitable UI;
	// this requires the account name to not overlap with API endpoint paths)
	if strings.HasPrefix(req.Account.Name, "keppel") {
		http.Error(w, `account names with the prefix "keppel" are reserved for internal use`, http.StatusUnprocessableEntity)
		return
	}
	if looksLikeAPIVersionRx.MatchString(req.Account.Name) {
		http.Error(w, `account names that look like API versions are reserved for internal use`, http.StatusUnprocessableEntity)
		return
	}

	for _, policy := range req.Account.GCPolicies {
		err := policy.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
	}

	for idx, policy := range req.Account.RBACPolicies {
		err := policy.ValidateAndNormalize()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		req.Account.RBACPolicies[idx] = policy
	}

	metadataJSONStr := ""
	if len(req.Account.Metadata) > 0 {
		metadataJSON, _ := json.Marshal(req.Account.Metadata)
		metadataJSONStr = string(metadataJSON)
	}

	gcPoliciesJSONStr := "[]"
	if len(req.Account.GCPolicies) > 0 {
		gcPoliciesJSON, _ := json.Marshal(req.Account.GCPolicies)
		gcPoliciesJSONStr = string(gcPoliciesJSON)
	}

	rbacPoliciesJSONStr := ""
	if len(req.Account.RBACPolicies) > 0 {
		rbacPoliciesJSON, _ := json.Marshal(req.Account.RBACPolicies)
		rbacPoliciesJSONStr = string(rbacPoliciesJSON)
	}

	accountToCreate := models.Account{
		Name:                     req.Account.Name,
		AuthTenantID:             req.Account.AuthTenantID,
		InMaintenance:            req.Account.InMaintenance,
		MetadataJSON:             metadataJSONStr,
		GCPoliciesJSON:           gcPoliciesJSONStr,
		RBACPoliciesJSON:         rbacPoliciesJSONStr,
		SecurityScanPoliciesJSON: "[]",
	}

	// validate replication policy
	if req.Account.ReplicationPolicy != nil {
		rp := *req.Account.ReplicationPolicy

		httpStatus, err := rp.ApplyToAccount(a.db, &accountToCreate)
		if err != nil {
			http.Error(w, err.Error(), httpStatus)
			return
		}
		//NOTE: There are some delayed checks below which require the existing account to be loaded from the DB first.
	}

	// validate validation policy
	if req.Account.ValidationPolicy != nil {
		vp := *req.Account.ValidationPolicy
		for _, label := range vp.RequiredLabels {
			if strings.Contains(label, ",") {
				http.Error(w, fmt.Sprintf(`invalid label name: %q`, label), http.StatusUnprocessableEntity)
				return
			}
		}

		accountToCreate.RequiredLabels = vp.JoinRequiredLabels()
	}

	// validate platform filter
	if req.Account.PlatformFilter != nil {
		if req.Account.ReplicationPolicy == nil {
			http.Error(w, `platform filter is only allowed on replica accounts`, http.StatusUnprocessableEntity)
			return
		}
		accountToCreate.PlatformFilter = req.Account.PlatformFilter
	}

	// check permission to create account
	authz := a.authenticateRequest(w, r, authTenantScope(keppel.CanChangeAccount, accountToCreate.AuthTenantID))
	if authz == nil {
		return
	}

	// check if account already exists
	account, err := keppel.FindAccount(a.db, req.Account.Name)
	if respondwith.ErrorText(w, err) {
		return
	}
	if account != nil && account.AuthTenantID != req.Account.AuthTenantID {
		http.Error(w, `account name already in use by a different tenant`, http.StatusConflict)
		return
	}

	// late replication policy validations (could not do these earlier because we
	// did not have `account` yet)
	if req.Account.ReplicationPolicy != nil {
		rp := *req.Account.ReplicationPolicy

		if rp.Strategy == "from_external_on_first_use" {
			// for new accounts, we need either full credentials or none
			if account == nil {
				if (rp.ExternalPeer.UserName == "") != (rp.ExternalPeer.Password == "") {
					http.Error(w, `need either both username and password or neither for "from_external_on_first_use" replication`, http.StatusUnprocessableEntity)
					return
				}
			}

			// for existing accounts, having only a username is acceptable if it's
			// unchanged (this case occurs when a client GETs the account, changes
			// something unrelated to replication, and PUTs the result; the password is
			// redacted in GET)
			if account != nil && rp.ExternalPeer.UserName != "" && rp.ExternalPeer.Password == "" {
				if rp.ExternalPeer.UserName == account.ExternalPeerUserName {
					rp.ExternalPeer.Password = account.ExternalPeerPassword // to pass the equality checks below
				} else {
					http.Error(w, `cannot change username for "from_external_on_first_use" replication without also changing password`, http.StatusUnprocessableEntity)
					return
				}
			}
		}
	}

	// replication strategy may not be changed after account creation
	if account != nil && req.Account.ReplicationPolicy != nil && !replicationPoliciesFunctionallyEqual(req.Account.ReplicationPolicy, renderReplicationPolicy(*account)) {
		http.Error(w, `cannot change replication policy on existing account`, http.StatusConflict)
		return
	}
	if account != nil && req.Account.PlatformFilter != nil && !reflect.DeepEqual(req.Account.PlatformFilter, account.PlatformFilter) {
		http.Error(w, `cannot change platform filter on existing account`, http.StatusConflict)
		return
	}

	// late RBAC policy validations (could not do these earlier because we did not
	// have `account` yet)
	isExternalReplica := req.Account.ReplicationPolicy != nil && req.Account.ReplicationPolicy.ExternalPeer.URL != ""
	if account != nil {
		isExternalReplica = account.ExternalPeerURL != ""
	}
	for _, policy := range req.Account.RBACPolicies {
		if slices.Contains(policy.Permissions, keppel.GrantsAnonymousFirstPull) && !isExternalReplica {
			http.Error(w, `RBAC policy with "anonymous_first_pull" may only be for external replica accounts`, http.StatusUnprocessableEntity)
			return
		}
	}

	// create account if required
	if account == nil {
		// sublease tokens are only relevant when creating replica accounts
		subleaseTokenSecret := ""
		if accountToCreate.UpstreamPeerHostName != "" {
			subleaseToken, err := SubleaseTokenFromRequest(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			subleaseTokenSecret = subleaseToken.Secret
		}

		// check permission to claim account name (this only happens here because
		// it's only relevant for account creations, not for updates)
		claimResult, err := a.fd.ClaimAccountName(r.Context(), accountToCreate, subleaseTokenSecret)
		switch claimResult {
		case keppel.ClaimSucceeded:
			// nothing to do
		case keppel.ClaimFailed:
			// user error
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		case keppel.ClaimErrored:
			// server error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy PlatformFilter when creating an account with the Replication Policy on_first_use
		if req.Account.ReplicationPolicy != nil {
			rp := *req.Account.ReplicationPolicy
			if rp.Strategy == "on_first_use" {
				var peer models.Peer
				err := a.db.SelectOne(&peer, `SELECT * FROM peers WHERE hostname = $1`, rp.UpstreamPeerHostName)
				if errors.Is(err, sql.ErrNoRows) {
					http.Error(w, fmt.Sprintf(`unknown peer registry: %q`, rp.UpstreamPeerHostName), http.StatusUnprocessableEntity)
					return
				}
				if respondwith.ErrorText(w, err) {
					return
				}

				viewScope := auth.Scope{
					ResourceType: "keppel_account",
					ResourceName: accountToCreate.Name,
					Actions:      []string{"view"},
				}
				client, err := peerclient.New(r.Context(), a.cfg, peer, viewScope)
				if respondwith.ErrorText(w, err) {
					return
				}

				var upstreamAccount Account
				err = client.GetForeignAccountConfigurationInto(r.Context(), &upstreamAccount, accountToCreate.Name)
				if respondwith.ErrorText(w, err) {
					return
				}

				if req.Account.PlatformFilter == nil {
					accountToCreate.PlatformFilter = upstreamAccount.PlatformFilter
				} else if !reflect.DeepEqual(req.Account.PlatformFilter, upstreamAccount.PlatformFilter) {
					// check if the peer PlatformFilter matches the primary account PlatformFilter
					jsonPlatformFilter, _ := json.Marshal(req.Account.PlatformFilter)
					jsonFilter, _ := json.Marshal(upstreamAccount.PlatformFilter)
					msg := fmt.Sprintf("peer account filter needs to match primary account filter: primary account %s, peer account %s ", jsonPlatformFilter, jsonFilter)
					http.Error(w, msg, http.StatusConflict)
					return
				}
			}
		}

		err = a.sd.CanSetupAccount(accountToCreate)
		if err != nil {
			msg := "cannot set up backing storage for this account: " + err.Error()
			http.Error(w, msg, http.StatusConflict)
			return
		}

		tx, err := a.db.Begin()
		if respondwith.ErrorText(w, err) {
			return
		}
		defer sqlext.RollbackUnlessCommitted(tx)

		account = &accountToCreate
		err = tx.Insert(account)
		if respondwith.ErrorText(w, err) {
			return
		}

		// commit the changes
		err = tx.Commit()
		if respondwith.ErrorText(w, err) {
			return
		}
		if userInfo := authz.UserIdentity.UserInfo(); userInfo != nil {
			a.auditor.Record(audittools.EventParameters{
				Time:       time.Now(),
				Request:    r,
				User:       userInfo,
				ReasonCode: http.StatusOK,
				Action:     cadf.CreateAction,
				Target:     AuditAccount{Account: *account},
			})
		}
	} else {
		// account != nil: update if necessary
		needsUpdate := false
		needsAudit := false
		if account.InMaintenance != accountToCreate.InMaintenance {
			account.InMaintenance = accountToCreate.InMaintenance
			needsUpdate = true
		}
		if account.MetadataJSON != accountToCreate.MetadataJSON {
			account.MetadataJSON = accountToCreate.MetadataJSON
			needsUpdate = true
		}
		if account.GCPoliciesJSON != accountToCreate.GCPoliciesJSON {
			account.GCPoliciesJSON = accountToCreate.GCPoliciesJSON
			needsUpdate = true
			needsAudit = true
		}
		if account.RBACPoliciesJSON != accountToCreate.RBACPoliciesJSON {
			account.RBACPoliciesJSON = accountToCreate.RBACPoliciesJSON
			needsUpdate = true
			needsAudit = true
		}
		if account.RequiredLabels != accountToCreate.RequiredLabels {
			account.RequiredLabels = accountToCreate.RequiredLabels
			needsUpdate = true
		}
		if account.ExternalPeerUserName != accountToCreate.ExternalPeerUserName {
			account.ExternalPeerUserName = accountToCreate.ExternalPeerUserName
			needsUpdate = true
		}
		if account.ExternalPeerPassword != accountToCreate.ExternalPeerPassword {
			account.ExternalPeerPassword = accountToCreate.ExternalPeerPassword
			needsUpdate = true
		}
		if needsUpdate {
			_, err := a.db.Update(account)
			if respondwith.ErrorText(w, err) {
				return
			}
		}
		if needsAudit {
			if userInfo := authz.UserIdentity.UserInfo(); userInfo != nil {
				a.auditor.Record(audittools.EventParameters{
					Time:       time.Now(),
					Request:    r,
					User:       userInfo,
					ReasonCode: http.StatusOK,
					Action:     cadf.UpdateAction,
					Target:     AuditAccount{Account: *account},
				})
			}
		}
	}

	accountRendered, err := a.renderAccount(*account)
	if respondwith.ErrorText(w, err) {
		return
	}
	respondwith.JSON(w, http.StatusOK, map[string]any{"account": accountRendered})
}

// Like reflect.DeepEqual, but ignores some fields that are allowed to be
// updated after account creation.
func replicationPoliciesFunctionallyEqual(lhs, rhs *keppel.ReplicationPolicy) bool {
	// one nil and one non-nil is not equal
	if (lhs == nil) != (rhs == nil) {
		return false
	}
	// two nil's are equal
	if lhs == nil {
		return true
	}

	// ignore pull credentials (the user shall be able to change these after account creation)
	lhsClone := *lhs
	rhsClone := *rhs
	lhsClone.ExternalPeer.UserName = ""
	lhsClone.ExternalPeer.Password = ""
	rhsClone.ExternalPeer.UserName = ""
	rhsClone.ExternalPeer.Password = ""
	return reflect.DeepEqual(lhsClone, rhsClone)
}

type deleteAccountRemainingManifest struct {
	RepositoryName string `json:"repository"`
	Digest         string `json:"digest"`
}

type deleteAccountRemainingManifests struct {
	Count uint64                           `json:"count"`
	Next  []deleteAccountRemainingManifest `json:"next"`
}

type deleteAccountRemainingBlobs struct {
	Count uint64 `json:"count"`
}

type deleteAccountResponse struct {
	RemainingManifests *deleteAccountRemainingManifests `json:"remaining_manifests,omitempty"`
	RemainingBlobs     *deleteAccountRemainingBlobs     `json:"remaining_blobs,omitempty"`
	Error              string                           `json:"error,omitempty"`
}

func (a *API) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	httpapi.IdentifyEndpoint(r, "/keppel/v1/accounts/:account")
	authz := a.authenticateRequest(w, r, accountScopeFromRequest(r, keppel.CanChangeAccount))
	if authz == nil {
		return
	}
	account := a.findAccountFromRequest(w, r, authz)
	if account == nil {
		return
	}

	resp, err := a.deleteAccount(r.Context(), *account)
	if respondwith.ErrorText(w, err) {
		return
	}
	if resp == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		respondwith.JSON(w, http.StatusConflict, resp)
	}
}

var (
	deleteAccountFindManifestsQuery = sqlext.SimplifyWhitespace(`
		SELECT r.name, m.digest
			FROM manifests m
			JOIN repos r ON m.repo_id = r.id
			JOIN accounts a ON a.name = r.account_name
			LEFT OUTER JOIN manifest_manifest_refs mmr ON mmr.repo_id = r.id AND m.digest = mmr.child_digest
		 WHERE a.name = $1 AND parent_digest IS NULL
		 LIMIT 10
	`)
	deleteAccountCountManifestsQuery = sqlext.SimplifyWhitespace(`
		SELECT COUNT(m.digest)
			FROM manifests m
			JOIN repos r ON m.repo_id = r.id
			JOIN accounts a ON a.name = r.account_name
		 WHERE a.name = $1
	`)
	deleteAccountReposQuery                   = `DELETE FROM repos WHERE account_name = $1`
	deleteAccountCountBlobsQuery              = `SELECT COUNT(id) FROM blobs WHERE account_name = $1`
	deleteAccountScheduleBlobSweepQuery       = `UPDATE accounts SET next_blob_sweep_at = $2 WHERE name = $1`
	deleteAccountMarkAllBlobsForDeletionQuery = `UPDATE blobs SET can_be_deleted_at = $2 WHERE account_name = $1`
)

func (a *API) deleteAccount(ctx context.Context, account models.Account) (*deleteAccountResponse, error) {
	if !account.InMaintenance {
		return &deleteAccountResponse{
			Error: "account must be set in maintenance first",
		}, nil
	}

	// can only delete account when user has deleted all manifests from it
	var nextManifests []deleteAccountRemainingManifest
	err := sqlext.ForeachRow(a.db, deleteAccountFindManifestsQuery, []any{account.Name},
		func(rows *sql.Rows) error {
			var m deleteAccountRemainingManifest
			err := rows.Scan(&m.RepositoryName, &m.Digest)
			nextManifests = append(nextManifests, m)
			return err
		},
	)
	if err != nil {
		return nil, err
	}
	if len(nextManifests) > 0 {
		manifestCount, err := a.db.SelectInt(deleteAccountCountManifestsQuery, account.Name)
		return &deleteAccountResponse{
			RemainingManifests: &deleteAccountRemainingManifests{
				Count: uint64(manifestCount),
				Next:  nextManifests,
			},
		}, err
	}

	// delete all repos (and therefore, all blob mounts), so that blob sweeping
	// can immediately take place
	_, err = a.db.Exec(deleteAccountReposQuery, account.Name)
	if err != nil {
		return nil, err
	}

	// can only delete account when all blobs have been deleted
	blobCount, err := a.db.SelectInt(deleteAccountCountBlobsQuery, account.Name)
	if err != nil {
		return nil, err
	}
	if blobCount > 0 {
		// make sure that blob sweep runs immediately
		_, err := a.db.Exec(deleteAccountMarkAllBlobsForDeletionQuery, account.Name, time.Now())
		if err != nil {
			return nil, err
		}
		_, err = a.db.Exec(deleteAccountScheduleBlobSweepQuery, account.Name, time.Now())
		if err != nil {
			return nil, err
		}
		return &deleteAccountResponse{
			RemainingBlobs: &deleteAccountRemainingBlobs{Count: uint64(blobCount)},
		}, nil
	}

	// start deleting the account in a transaction
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}
	defer sqlext.RollbackUnlessCommitted(tx)
	_, err = tx.Delete(&account)
	if err != nil {
		return nil, err
	}

	// before committing the transaction, confirm account deletion with the
	// storage driver and the federation driver
	err = a.sd.CleanupAccount(account)
	if err != nil {
		return &deleteAccountResponse{Error: err.Error()}, nil
	}
	err = a.fd.ForfeitAccountName(ctx, account)
	if err != nil {
		return &deleteAccountResponse{Error: err.Error()}, nil
	}

	return nil, tx.Commit()
}

func (a *API) handlePostAccountSublease(w http.ResponseWriter, r *http.Request) {
	httpapi.IdentifyEndpoint(r, "/keppel/v1/accounts/:account/sublease")
	authz := a.authenticateRequest(w, r, accountScopeFromRequest(r, keppel.CanChangeAccount))
	if authz == nil {
		return
	}
	account := a.findAccountFromRequest(w, r, authz)
	if account == nil {
		return
	}

	if account.UpstreamPeerHostName != "" {
		http.Error(w, "operation not allowed for replica accounts", http.StatusBadRequest)
		return
	}

	st := SubleaseToken{
		AccountName:     account.Name,
		PrimaryHostname: a.cfg.APIPublicHostname,
	}

	var err error
	st.Secret, err = a.fd.IssueSubleaseTokenSecret(r.Context(), *account)
	if respondwith.ErrorText(w, err) {
		return
	}

	// only serialize SubleaseToken if it contains a secret at all
	var serialized string
	if st.Secret == "" {
		serialized = ""
	} else {
		serialized = st.Serialize()
	}

	respondwith.JSON(w, http.StatusOK, map[string]any{"sublease_token": serialized})
}

func (a *API) handleGetSecurityScanPolicies(w http.ResponseWriter, r *http.Request) {
	httpapi.IdentifyEndpoint(r, "/keppel/v1/accounts/:account/security_scan_policies")
	authz := a.authenticateRequest(w, r, accountScopeFromRequest(r, keppel.CanViewAccount))
	if authz == nil {
		return
	}
	account := a.findAccountFromRequest(w, r, authz)
	if account == nil {
		return
	}

	respondwith.JSON(w, http.StatusOK, map[string]any{"policies": json.RawMessage(account.SecurityScanPoliciesJSON)})
}

func (a *API) handlePutSecurityScanPolicies(w http.ResponseWriter, r *http.Request) {
	httpapi.IdentifyEndpoint(r, "/keppel/v1/accounts/:account/security_scan_policies")
	authz := a.authenticateRequest(w, r, accountScopeFromRequest(r, keppel.CanChangeAccount))
	if authz == nil {
		return
	}
	account := a.findAccountFromRequest(w, r, authz)
	if account == nil {
		return
	}

	// decode existing policies
	var dbPolicies []keppel.SecurityScanPolicy
	err := json.Unmarshal([]byte(account.SecurityScanPoliciesJSON), &dbPolicies)
	if respondwith.ErrorText(w, err) {
		return
	}

	// decode request body
	var req struct {
		Policies []keppel.SecurityScanPolicy `json:"policies"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, "request body is not valid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// apply computed values and validate each input policy on its own
	currentUserName := authz.UserIdentity.UserName()
	var errs errext.ErrorSet
	for idx, policy := range req.Policies {
		path := fmt.Sprintf("policies[%d]", idx)
		errs.Append(policy.Validate(path))

		switch policy.ManagingUserName {
		case "$REQUESTER":
			req.Policies[idx].ManagingUserName = currentUserName
		case "", currentUserName:
			// acceptable
		default:
			if !slices.Contains(dbPolicies, policy) {
				errs.Addf("cannot apply this new or updated policy that is managed by a different user: %s", policy)
			}
		}
	}

	// check that updated or deleted policies are either unmanaged or managed by
	// the requester
	for _, dbPolicy := range dbPolicies {
		if slices.Contains(req.Policies, dbPolicy) {
			continue
		}
		managingUserName := dbPolicy.ManagingUserName
		if managingUserName != "" && managingUserName != currentUserName {
			errs.Addf("cannot update or delete this existing policy that is managed by a different user: %s", dbPolicy)
		}
	}

	// report validation errors
	if !errs.IsEmpty() {
		http.Error(w, errs.Join("\n"), http.StatusUnprocessableEntity)
		return
	}

	// update policies in DB
	jsonBuf, err := json.Marshal(req.Policies)
	if respondwith.ErrorText(w, err) {
		return
	}
	_, err = a.db.Exec(`UPDATE accounts SET security_scan_policies_json = $1 WHERE name = $2`,
		string(jsonBuf), account.Name)
	if respondwith.ErrorText(w, err) {
		return
	}

	// generate audit events
	submitAudit := func(action cadf.Action, target audittools.TargetRenderer) {
		if userInfo := authz.UserIdentity.UserInfo(); userInfo != nil {
			a.auditor.Record(audittools.EventParameters{
				Time:       time.Now(),
				Request:    r,
				User:       userInfo,
				ReasonCode: http.StatusOK,
				Action:     action,
				Target:     target,
			})
		}
	}
	for _, policy := range req.Policies {
		if !slices.Contains(dbPolicies, policy) {
			submitAudit("create/security-scan-policy", AuditSecurityScanPolicy{
				Account: *account,
				Policy:  policy,
			})
		}
	}
	for _, policy := range dbPolicies {
		if !slices.Contains(req.Policies, policy) {
			submitAudit("delete/security-scan-policy", AuditSecurityScanPolicy{
				Account: *account,
				Policy:  policy,
			})
		}
	}

	respondwith.JSON(w, http.StatusOK, map[string]any{"policies": req.Policies})
}
