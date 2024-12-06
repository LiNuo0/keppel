/*******************************************************************************
*
* Copyright 2024 SAP SE
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You should have received a copy of the License along with this
* program. If not, you may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

package processor

import (
	"encoding/json"

	"github.com/sapcc/go-api-declarations/cadf"
	"github.com/sapcc/go-bits/must"

	"github.com/sapcc/keppel/internal/models"
)

// AuditAccount is an audittools.Target.
type AuditAccount struct {
	Account models.Account
}

// Render implements the audittools.Target interface.
func (a AuditAccount) Render() cadf.Resource {
	res := cadf.Resource{
		TypeURI:   "docker-registry/account",
		ID:        string(a.Account.Name),
		ProjectID: a.Account.AuthTenantID,
	}

	gcPoliciesJSON := a.Account.GCPoliciesJSON
	if gcPoliciesJSON != "" && gcPoliciesJSON != "[]" {
		attachment := must.Return(cadf.NewJSONAttachment("gc-policies", json.RawMessage(gcPoliciesJSON)))
		res.Attachments = append(res.Attachments, attachment)
	}

	rbacPoliciesJSON := a.Account.RBACPoliciesJSON
	if rbacPoliciesJSON != "" && rbacPoliciesJSON != "[]" {
		attachment := must.Return(cadf.NewJSONAttachment("rbac-policies", json.RawMessage(rbacPoliciesJSON)))
		res.Attachments = append(res.Attachments, attachment)
	}

	return res
}

// AuditQuotas is an audittools.Target.
type AuditQuotas struct {
	QuotasBefore models.Quotas
	QuotasAfter  models.Quotas
}

// Render implements the audittools.Target interface.
func (a AuditQuotas) Render() cadf.Resource {
	return cadf.Resource{
		TypeURI:   "docker-registry/project-quota",
		ID:        a.QuotasAfter.AuthTenantID,
		ProjectID: a.QuotasAfter.AuthTenantID,
		Attachments: []cadf.Attachment{
			must.Return(cadf.NewJSONAttachment("payload-before", a.QuotasBefore)),
			must.Return(cadf.NewJSONAttachment("payload", a.QuotasAfter)),
		},
	}
}
