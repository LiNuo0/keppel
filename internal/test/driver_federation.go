/******************************************************************************
*
*  Copyright 2019 SAP SE
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

package test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sapcc/keppel/internal/keppel"
	"github.com/sapcc/keppel/internal/models"
)

var (
	federationDriversForThisUnitTest []*FederationDriver
)

// FederationDriver (driver ID "unittest") is a keppel.FederationDriver for unit tests.
type FederationDriver struct {
	APIPublicHostName              string
	ClaimFailsBecauseOfUserError   bool
	ClaimFailsBecauseOfServerError bool
	ForfeitFails                   bool
	NextSubleaseTokenSecretToIssue string
	ValidSubleaseTokenSecrets      map[string]string // maps accountName => subleaseTokenSecret
	RecordedAccounts               []AccountRecordedByFederationDriver
}

// AccountRecordedByFederationDriver appears in type FederationDriver.
type AccountRecordedByFederationDriver struct {
	Account    models.Account
	RecordedAt time.Time
}

func init() {
	keppel.FederationDriverRegistry.Add(func() keppel.FederationDriver { return &FederationDriver{} })
}

// PluginTypeID implements the keppel.FederationDriver interface.
func (d *FederationDriver) PluginTypeID() string { return "unittest" }

// Init implements the keppel.FederationDriver interface.
func (d *FederationDriver) Init(ctx context.Context, ad keppel.AuthDriver, cfg keppel.Configuration) error {
	d.APIPublicHostName = cfg.APIPublicHostname
	d.ValidSubleaseTokenSecrets = make(map[string]string)
	federationDriversForThisUnitTest = append(federationDriversForThisUnitTest, d)
	return nil
}

// ClaimAccountName implements the keppel.FederationDriver interface.
func (d *FederationDriver) ClaimAccountName(ctx context.Context, account models.Account, subleaseTokenSecret string) (keppel.ClaimResult, error) {
	// simulated failures for primary accounts
	if d.ClaimFailsBecauseOfUserError {
		return keppel.ClaimFailed, fmt.Errorf("cannot assign name %q to auth tenant %q", account.Name, account.AuthTenantID)
	}
	if d.ClaimFailsBecauseOfServerError {
		return keppel.ClaimErrored, fmt.Errorf("failed to assign name %q to auth tenant %q", account.Name, account.AuthTenantID)
	}

	// for replica accounts, do the regular sublease-token dance
	if account.UpstreamPeerHostName != "" {
		expectedTokenSecret, exists := d.ValidSubleaseTokenSecrets[account.Name]
		if !exists || subleaseTokenSecret != expectedTokenSecret {
			return keppel.ClaimFailed, errors.New("wrong sublease token")
		}
		// each sublease token can only be used once
		delete(d.ValidSubleaseTokenSecrets, account.Name)
	}

	return keppel.ClaimSucceeded, nil
}

// IssueSubleaseTokenSecret implements the keppel.FederationDriver interface.
func (d *FederationDriver) IssueSubleaseTokenSecret(ctx context.Context, account models.Account) (string, error) {
	// issue each sublease token only once
	t := d.NextSubleaseTokenSecretToIssue
	d.NextSubleaseTokenSecretToIssue = ""
	return t, nil
}

// ForfeitAccountName implements the keppel.FederationDriver interface.
func (d *FederationDriver) ForfeitAccountName(ctx context.Context, account models.Account) error {
	if d.ForfeitFails {
		return errors.New("ForfeitAccountName failing as requested")
	}
	return nil
}

// RecordExistingAccount implements the keppel.FederationDriver interface.
func (d *FederationDriver) RecordExistingAccount(ctx context.Context, account models.Account, now time.Time) error {
	account.NextFederationAnnouncementAt = nil // this pointer type is poison for DeepEqual tests

	d.RecordedAccounts = append(d.RecordedAccounts, AccountRecordedByFederationDriver{
		Account:    account,
		RecordedAt: now,
	})
	return nil
}

// FindPrimaryAccount implements the keppel.FederationDriver interface.
func (d *FederationDriver) FindPrimaryAccount(ctx context.Context, accountName string) (string, error) {
	for _, fd := range federationDriversForThisUnitTest {
		for _, a := range fd.RecordedAccounts {
			if a.Account.Name == accountName && a.Account.UpstreamPeerHostName == "" {
				return fd.APIPublicHostName, nil
			}
		}
	}
	return "", keppel.ErrNoSuchPrimaryAccount
}
