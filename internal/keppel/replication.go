/******************************************************************************
*
*  Copyright 2024 SAP SE
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

package keppel

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sapcc/keppel/internal/models"
)

// ReplicationPolicy represents a replication policy in the API.
type ReplicationPolicy struct {
	Strategy string `json:"strategy"`
	// only for `on_first_use`
	UpstreamPeerHostName string `json:"upstream_peer_hostname"`
	// only for `from_external_on_first_use`
	ExternalPeer ReplicationExternalPeerSpec `json:"external_peer"`
}

// ReplicationExternalPeerSpec appears in type ReplicationPolicy.
type ReplicationExternalPeerSpec struct {
	URL      string `json:"url"`
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
func (r ReplicationPolicy) MarshalJSON() ([]byte, error) {
	switch r.Strategy {
	case "on_first_use":
		data := struct {
			Strategy             string `json:"strategy"`
			UpstreamPeerHostName string `json:"upstream"`
		}{r.Strategy, r.UpstreamPeerHostName}
		return json.Marshal(data)
	case "from_external_on_first_use":
		data := struct {
			Strategy     string                      `json:"strategy"`
			ExternalPeer ReplicationExternalPeerSpec `json:"upstream"`
		}{r.Strategy, r.ExternalPeer}
		return json.Marshal(data)
	default:
		return nil, fmt.Errorf("do not know how to serialize ReplicationPolicy with strategy %q", r.Strategy)
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *ReplicationPolicy) UnmarshalJSON(buf []byte) error {
	var s struct {
		Strategy string          `json:"strategy"`
		Upstream json.RawMessage `json:"upstream"`
	}
	err := json.Unmarshal(buf, &s)
	if err != nil {
		return err
	}
	r.Strategy = s.Strategy

	switch r.Strategy {
	case "on_first_use":
		return json.Unmarshal(s.Upstream, &r.UpstreamPeerHostName)
	case "from_external_on_first_use":
		return json.Unmarshal(s.Upstream, &r.ExternalPeer)
	default:
		return fmt.Errorf("do not know how to deserialize ReplicationPolicy with strategy %q", r.Strategy)
	}
}

// ApplyToAccount validates this policy and stores it in the given account model.
func (r ReplicationPolicy) ApplyToAccount(db *DB, dbAccount *models.Account) *RegistryV2Error {
	switch r.Strategy {
	case "on_first_use":
		peerCount, err := db.SelectInt(`SELECT COUNT(*) FROM peers WHERE hostname = $1`, r.UpstreamPeerHostName)
		if err != nil {
			return AsRegistryV2Error(err).WithStatus(http.StatusInternalServerError)
		}

		if peerCount == 0 {
			err := fmt.Errorf(`unknown peer registry: %q`, r.UpstreamPeerHostName)
			return AsRegistryV2Error(err).WithStatus(http.StatusUnprocessableEntity)
		}
		dbAccount.UpstreamPeerHostName = r.UpstreamPeerHostName
	case "from_external_on_first_use":
		if r.ExternalPeer.URL == "" {
			err := errors.New(`missing upstream URL for "from_external_on_first_use" replication`)
			return AsRegistryV2Error(err).WithStatus(http.StatusUnprocessableEntity)
		}
		dbAccount.ExternalPeerURL = r.ExternalPeer.URL
		dbAccount.ExternalPeerUserName = r.ExternalPeer.UserName
		dbAccount.ExternalPeerPassword = r.ExternalPeer.Password
	default:
		err := fmt.Errorf("strategy %s is unsupported", r.Strategy)
		return AsRegistryV2Error(err).WithStatus(http.StatusUnprocessableEntity)
	}

	return nil
}
