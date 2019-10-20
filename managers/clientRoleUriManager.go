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

//ClientRoleURI ClientRoleURI
type ClientRoleURI struct {
	ClientRoleID       int64
	ClientAllowedURIID int64
}

//AddClientRoleURI AddClientRoleURI
func (m *OauthManager) AddClientRoleURI(r *ClientRoleURI) bool {
	var ru odb.ClientRoleURI
	ru.ClientRoleID = r.ClientRoleID
	ru.ClientAllowedURIID = r.ClientAllowedURIID
	suc := m.Db.AddClientRoleURI(&ru)
	return suc

}

//GetClientRoleAllowedURIList GetClientRoleAllowedURIList
func (m *OauthManager) GetClientRoleAllowedURIList(roleID int64) *[]ClientRoleURI {
	var rtn []ClientRoleURI
	rul := m.Db.GetClientRoleAllowedURIList(roleID)
	for _, ru := range *rul {
		var r ClientRoleURI
		r.ClientRoleID = ru.ClientRoleID
		r.ClientAllowedURIID = ru.ClientAllowedURIID
		rtn = append(rtn, r)
	}
	return &rtn
}

//DeleteClientRoleURI DeleteClientRoleURI
func (m *OauthManager) DeleteClientRoleURI(r *ClientRoleURI) bool {
	var ru odb.ClientRoleURI
	ru.ClientRoleID = r.ClientRoleID
	ru.ClientAllowedURIID = r.ClientAllowedURIID
	suc := m.Db.DeleteClientRoleURI(&ru)
	return suc
}
