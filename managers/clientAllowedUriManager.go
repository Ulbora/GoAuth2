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

//ClientAllowedURI ClientAllowedURI
type ClientAllowedURI struct {
	ID       int64  `json:"id"`
	URI      string `json:"uri"`
	ClientID int64  `json:"clientId"`
}

//AddClientAllowedURI AddClientAllowedURI
func (m *OauthManager) AddClientAllowedURI(au *ClientAllowedURI) (bool, int64) {
	var cu odb.ClientAllowedURI
	cu.URI = au.URI
	cu.ClientID = au.ClientID
	suc, id := m.Db.AddClientAllowedURI(&cu)
	return suc, id
}

//UpdateClientAllowedURI UpdateClientAllowedURI
func (m *OauthManager) UpdateClientAllowedURI(au *ClientAllowedURI) bool {
	var cu odb.ClientAllowedURI
	cu.URI = au.URI
	cu.ID = au.ID
	suc := m.Db.UpdateClientAllowedURI(&cu)
	return suc
}

//GetClientAllowedURI GetClientAllowedURI
func (m *OauthManager) GetClientAllowedURI(id int64) *ClientAllowedURI {
	var rtn ClientAllowedURI
	au := m.Db.GetClientAllowedURIByID(id)
	rtn.ID = au.ID
	rtn.URI = au.URI
	rtn.ClientID = au.ClientID
	return &rtn
}

//GetClientAllowedURIList GetClientAllowedURIList
func (m *OauthManager) GetClientAllowedURIList(clientID int64) *[]ClientAllowedURI {
	var rtn = []ClientAllowedURI{}
	aul := m.Db.GetClientAllowedURIList(clientID)
	for _, au := range *aul {
		var u ClientAllowedURI
		u.ID = au.ID
		u.URI = au.URI
		u.ClientID = au.ClientID
		rtn = append(rtn, u)
	}
	return &rtn
}

//DeleteClientAllowedURI DeleteClientAllowedURI
func (m *OauthManager) DeleteClientAllowedURI(id int64) bool {
	suc := m.Db.DeleteClientAllowedURI(id)
	return suc
}
