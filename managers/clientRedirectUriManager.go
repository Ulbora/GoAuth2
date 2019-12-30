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

//ClientRedirectURI ClientRedirectURI
type ClientRedirectURI struct {
	ID       int64  `json:"id"`
	URI      string `json:"uri"`
	ClientID int64  `json:"clientId"`
}

//AddClientRedirectURI AddClientRedirectURI
func (m *OauthManager) AddClientRedirectURI(ru *ClientRedirectURI) (bool, int64) {
	// var suc bool
	// var id int64
	var cru odb.ClientRedirectURI
	cru.URI = ru.URI
	cru.ClientID = ru.ClientID
	suc, id := m.Db.AddClientRedirectURI(nil, &cru)
	return suc, id
}

//GetClientRedirectURIList GetClientRedirectURIList
func (m *OauthManager) GetClientRedirectURIList(clientID int64) *[]ClientRedirectURI {
	var rtn []ClientRedirectURI
	ul := m.Db.GetClientRedirectURIList(clientID)
	for _, u := range *ul {
		var ui ClientRedirectURI
		ui.ID = u.ID
		ui.URI = u.URI
		ui.ClientID = u.ClientID
		rtn = append(rtn, ui)
	}
	return &rtn
}

//DeleteClientRedirectURI DeleteClientRedirectURI
func (m *OauthManager) DeleteClientRedirectURI(id int64) bool {
	suc := m.Db.DeleteClientRedirectURI(nil, id)
	return suc
}
