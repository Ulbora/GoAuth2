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

//ClientGrantType ClientGrantType
type ClientGrantType struct {
	ID        int64
	GrantType string
	ClientID  int64
}

//AddClientGrantType AddClientGrantType
func (m *OauthManager) AddClientGrantType(gt *ClientGrantType) (bool, int64) {
	var cgt odb.ClientGrantType
	cgt.GrantType = gt.GrantType
	cgt.ClientID = gt.ClientID
	suc, id := m.Db.AddClientGrantType(&cgt)
	return suc, id
}

func (m *OauthManager) grantTypeTurnedOn(clientID int64, grantType string) bool {
	var rtn bool
	gtlist := m.Db.GetClientGrantTypeList(clientID)
	for _, gt := range *gtlist {
		if gt.GrantType == grantType {
			rtn = true
		}
	}
	return rtn
}

//GetClientGrantTypeList GetClientGrantTypeList
func (m *OauthManager) GetClientGrantTypeList(clientID int64) *[]ClientGrantType {
	var rtn []ClientGrantType
	gtl := m.Db.GetClientGrantTypeList(clientID)
	for _, gt := range *gtl {
		var g ClientGrantType
		g.ID = gt.ID
		g.GrantType = gt.GrantType
		g.ClientID = gt.ClientID
		rtn = append(rtn, g)
	}
	return &rtn
}

//DeleteClientGrantType DeleteClientGrantType
func (m *OauthManager) DeleteClientGrantType(id int64) bool {
	suc := m.Db.DeleteClientGrantType(id)
	return suc
}
