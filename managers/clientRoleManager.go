package managers

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

import (
	//"fmt"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//ClientRole ClientRole
type ClientRole struct {
	ID       int64
	Role     string
	ClientID int64
}

//AddClientRole AddClientRole
func (m *OauthManager) AddClientRole(r *ClientRole) (bool, int64) {
	// var suc bool
	// var id int64
	var cr odb.ClientRole
	cr.Role = r.Role
	cr.ClientID = r.ClientID
	suc, id := m.Db.AddClientRole(&cr)
	return suc, id
}

//GetClientRoleList GetClientRoleList
func (m *OauthManager) GetClientRoleList(clientID int64) *[]ClientRole {
	var rtn []ClientRole
	rl := m.Db.GetClientRoleList(clientID)
	for _, r := range *rl {
		var cr ClientRole
		cr.ID = r.ID
		cr.Role = r.Role
		cr.ClientID = r.ClientID
		rtn = append(rtn, cr)
	}
	return &rtn
}

//DeleteClientRole DeleteClientRole
func (m *OauthManager) DeleteClientRole(id int64) bool {
	suc := m.Db.DeleteClientRole(id)
	return suc
}
