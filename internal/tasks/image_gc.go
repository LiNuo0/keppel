/******************************************************************************
*
*  Copyright 2021 SAP SE
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

package tasks

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/opencontainers/go-digest"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sapcc/go-bits/jobloop"
	"github.com/sapcc/go-bits/logg"
	"github.com/sapcc/go-bits/sqlext"

	"github.com/sapcc/keppel/internal/keppel"
	"github.com/sapcc/keppel/internal/models"
	"github.com/sapcc/keppel/internal/processor"
)

var imageGCRepoSelectQuery = sqlext.SimplifyWhitespace(`
	SELECT * FROM repos
		WHERE (next_gc_at IS NULL OR next_gc_at < $1)
	-- repos without any syncs first, then sorted by last sync
	ORDER BY next_gc_at IS NULL DESC, next_gc_at ASC
	-- only one repo at a time
	LIMIT 1
`)

var imageGCResetStatusQuery = sqlext.SimplifyWhitespace(`
	UPDATE manifests SET gc_status_json = '{"relevant_policies":[]}' WHERE repo_id = $1
`)

var imageGCRepoDoneQuery = sqlext.SimplifyWhitespace(`
	UPDATE repos SET next_gc_at = $2 WHERE id = $1
`)

// ManifestGarbageCollectionJob is a job. Each task finds the a where GC has
// not been performed for more than an hour, and performs GC based on the GC
// policies configured on the repo's account.
func (j *Janitor) ManifestGarbageCollectionJob(registerer prometheus.Registerer) jobloop.Job { //nolint: dupl // interface implementation of different things
	return (&jobloop.ProducerConsumerJob[models.Repository]{
		Metadata: jobloop.JobMetadata{
			ReadableName: "manifest garbage collection",
			CounterOpts: prometheus.CounterOpts{
				Name: "keppel_image_garbage_collections",
				Help: "Counter for image garbage collection runs in repos.",
			},
		},
		DiscoverTask: func(_ context.Context, _ prometheus.Labels) (repo models.Repository, err error) {
			err = j.db.SelectOne(&repo, imageGCRepoSelectQuery, j.timeNow())
			return repo, err
		},
		ProcessTask: j.garbageCollectManifestsInRepo,
	}).Setup(registerer)
}

func (j *Janitor) garbageCollectManifestsInRepo(ctx context.Context, repo models.Repository, labels prometheus.Labels) (returnErr error) {
	// load GC policies for this repository
	account, err := keppel.FindAccount(j.db, repo.AccountName)
	if err != nil {
		return fmt.Errorf("cannot find account for repo %s: %w", repo.FullName(), err)
	}
	policies, err := keppel.ParseGCPolicies(*account)
	if err != nil {
		return fmt.Errorf("cannot load GC policies for account %s: %w", account.Name, err)
	}
	var policiesForRepo []keppel.GCPolicy
	for idx, policy := range policies {
		err := policy.Validate()
		if err != nil {
			return fmt.Errorf("GC policy #%d for account %s is invalid: %w", idx+1, account.Name, err)
		}
		if policy.MatchesRepository(repo.Name) {
			policiesForRepo = append(policiesForRepo, policy)
		}
	}

	// execute GC policies
	if len(policiesForRepo) > 0 {
		err = j.executeGCPolicies(ctx, account.Reduced(), repo, policiesForRepo)
		if err != nil {
			return err
		}
	} else {
		// if there are no policies to apply, we can skip a whole bunch of work, but
		// we still need to update the GCStatusJSON field on the repo's manifests to
		// make sure those statuses don't refer to deleted GC policies
		_, err = j.db.Exec(imageGCResetStatusQuery, repo.ID)
		if err != nil {
			return err
		}
	}

	_, err = j.db.Exec(imageGCRepoDoneQuery, repo.ID, j.timeNow().Add(j.addJitter(1*time.Hour)))
	return err
}

type manifestData struct {
	Manifest      models.Manifest
	TagNames      []string
	ParentDigests []string
	GCStatus      keppel.GCStatus
	IsDeleted     bool
}

func (j *Janitor) executeGCPolicies(ctx context.Context, account models.ReducedAccount, repo models.Repository, policies []keppel.GCPolicy) error {
	// load manifests in repo
	var dbManifests []models.Manifest
	_, err := j.db.Select(&dbManifests, `SELECT * FROM manifests WHERE repo_id = $1`, repo.ID)
	if err != nil {
		return err
	}

	// setup a bit of structure to track state in during the policy evaluation
	var manifests []*manifestData
	for _, m := range dbManifests {
		manifests = append(manifests, &manifestData{
			Manifest: m,
			GCStatus: keppel.GCStatus{
				ProtectedByRecentUpload: m.PushedAt.After(j.timeNow().Add(-10 * time.Minute)),
			},
			IsDeleted: false,
		})
	}

	// load tags (for matching policies on match_tag, except_tag and only_untagged)
	query := `SELECT digest, name FROM tags WHERE repo_id = $1`
	err = sqlext.ForeachRow(j.db, query, []any{repo.ID}, func(rows *sql.Rows) error {
		var (
			digest  digest.Digest
			tagName string
		)
		err := rows.Scan(&digest, &tagName)
		if err != nil {
			return err
		}
		for _, m := range manifests {
			if m.Manifest.Digest == digest {
				m.TagNames = append(m.TagNames, tagName)
				break
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// check manifest-manifest relations to fill GCStatus.ProtectedByManifest
	query = `SELECT parent_digest, child_digest FROM manifest_manifest_refs WHERE repo_id = $1`
	err = sqlext.ForeachRow(j.db, query, []any{repo.ID}, func(rows *sql.Rows) error {
		var (
			parentDigest string
			childDigest  digest.Digest
		)
		err := rows.Scan(&parentDigest, &childDigest)
		if err != nil {
			return err
		}
		for _, m := range manifests {
			if m.Manifest.Digest == childDigest {
				m.ParentDigests = append(m.ParentDigests, parentDigest)
				break
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, m := range manifests {
		if len(m.ParentDigests) > 0 {
			sort.Strings(m.ParentDigests) // for deterministic test behavior
			m.GCStatus.ProtectedByParentManifest = m.ParentDigests[0]
		}
	}

	// evaluate policies in order
	proc := j.processor()
	for _, policy := range policies {
		err := j.evaluatePolicy(ctx, proc, manifests, account, repo, policy)
		if err != nil {
			return err
		}
	}

	return j.persistGCStatus(manifests, repo.ID)
}

func (j *Janitor) evaluatePolicy(ctx context.Context, proc *processor.Processor, manifests []*manifestData, account models.ReducedAccount, repo models.Repository, policy keppel.GCPolicy) error {
	// for some time constraint matches, we need to know which manifests are
	// still alive
	var aliveManifests []models.Manifest
	for _, m := range manifests {
		if !m.IsDeleted {
			aliveManifests = append(aliveManifests, m.Manifest)
		}
	}

	// evaluate policy for each manifest
	for _, m := range manifests {
		// skip those manifests that are already deleted, and those which are
		// protected by an earlier policy or one of the baseline checks above
		if m.IsDeleted || m.GCStatus.IsProtected() {
			continue
		}

		// track matching "delete" policies in GCStatus to allow users insight
		// into how policies match
		if policy.Action == "delete" {
			m.GCStatus.RelevantPolicies = append(m.GCStatus.RelevantPolicies, policy)
		}

		// evaluate constraints
		if !policy.MatchesTags(m.TagNames) {
			continue
		}
		if !policy.MatchesTimeConstraint(m.Manifest, aliveManifests, j.timeNow()) {
			continue
		}

		pCopied := policy
		// execute policy action
		switch policy.Action {
		case "protect":
			m.GCStatus.ProtectedByPolicy = &pCopied
		case "delete":
			err := proc.DeleteManifest(ctx, account, repo, m.Manifest.Digest, keppel.AuditContext{
				UserIdentity: janitorUserIdentity{
					TaskName: "policy-driven-gc",
					GCPolicy: &pCopied,
				},
				Request: janitorDummyRequest,
			})
			if err != nil {
				return err
			}
			m.IsDeleted = true
			policyJSON, _ := json.Marshal(policy)
			logg.Info("GC on repo %s: deleted manifest %s because of policy %s", repo.FullName(), m.Manifest.Digest, string(policyJSON))
		default:
			// defense in depth: we already did p.Validate() earlier
			return fmt.Errorf("unexpected GC policy action: %q (why was this not caught by Validate!?)", policy.Action)
		}
	}

	return nil
}

func (j *Janitor) persistGCStatus(manifests []*manifestData, repoID int64) error {
	// finalize and persist GCStatus for all affected manifests
	query := `UPDATE manifests SET gc_status_json = $1 WHERE repo_id = $2 AND digest = $3`
	err := sqlext.WithPreparedStatement(j.db, query, func(stmt *sql.Stmt) error {
		for _, m := range manifests {
			if m.IsDeleted {
				continue
			}
			// to simplify UI, show only EITHER protection status OR relevant deleting
			// policies, not both
			if m.GCStatus.IsProtected() {
				m.GCStatus.RelevantPolicies = nil
			}
			gcStatusJSON, err := json.Marshal(m.GCStatus)
			if err != nil {
				return err
			}
			_, err = stmt.Exec(string(gcStatusJSON), repoID, m.Manifest.Digest)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("while persisting GCStatus: %w", err)
	}
	return nil
}
