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

package models

import "time"

// Peer contains a record from the `peers` table.
type Peer struct {
	HostName string `db:"hostname"`

	// OurPassword is what we use to log in at the peer.
	OurPassword string `db:"our_password"`

	// TheirCurrentPasswordHash and TheirPreviousPasswordHash is what the peer
	// uses to log in with us. Passwords are rotated hourly. We allow access with
	// the current *and* the previous password to avoid a race where we enter the
	// new password in the database and then reject authentication attempts from
	// the peer before we told them about the new password.
	TheirCurrentPasswordHash  string `db:"their_current_password_hash"`
	TheirPreviousPasswordHash string `db:"their_previous_password_hash"`

	// LastPeeredAt is when we last issued a new password for this peer.
	LastPeeredAt *time.Time `db:"last_peered_at"` // see tasks.IssueNewPasswordForPeer
}
