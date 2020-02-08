package managers

import (
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

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

//Implicit Implicit
type Implicit struct {
	ClientID    int64
	UserID      string
	Scope       string
	RedirectURI string
	CallbackURI string
}

//ImplicitReturn ImplicitReturn
type ImplicitReturn struct {
	ID    int64
	Token string
}

//ImplicitClient ImplicitClient
type ImplicitClient struct {
	Valid      bool
	ClientName string
	WebSite    string
}

//AuthorizeImplicit AuthorizeImplicit
func (m *OauthManager) AuthorizeImplicit(imp *Implicit) (bool, *ImplicitReturn) {
	var suc bool
	var rtn *ImplicitReturn
	client := m.Db.GetClient(imp.ClientID)
	if client.Enabled {
		rtu := m.Db.GetClientRedirectURI(imp.ClientID, imp.RedirectURI)
		if rtu.ID > 0 {
			//here
			//check that grant type is on
			gton := m.grantTypeTurnedOn(imp.ClientID, implicitGrantType)
			m.Log.Debug("grant turned on: ", gton)
			if gton {
				impgt := m.Db.GetImplicitGrant(imp.ClientID, imp.UserID)
				m.Log.Debug("impgt: ", impgt)
				var scopeStrList []string
				if len(*impgt) > 0 && (*impgt)[0].ID != 0 {
					scopeList := m.Db.GetImplicitGrantScopeList((*impgt)[0].ID)
					m.Log.Debug("scopeList: ", scopeList)
					var scopeFound bool
					for _, s := range *scopeList {
						if s.Scope == imp.Scope {
							scopeFound = true
							break
						}
					}
					m.Log.Debug("scopeFound: ", scopeFound)
					for _, s := range *scopeList {
						scopeStrList = append(scopeStrList, s.Scope)
					}
					if scopeFound {
						suc, rtn = m.processImplicitInsert(imp, &scopeStrList, true)
					} else {
						scopeStrList = append(scopeStrList, imp.Scope)
						suc, rtn = m.processImplicitInsert(imp, &scopeStrList, true)
					}
				} else {
					scopeStrList = append(scopeStrList, imp.Scope)
					suc, rtn = m.processImplicitInsert(imp, &scopeStrList, false)
				}
			}
		}
	}
	return suc, rtn
}

//CheckImplicitApplicationAuthorization CheckImplicitApplicationAuthorization
func (m *OauthManager) CheckImplicitApplicationAuthorization(imp *Implicit) (authorized bool) {
	if imp.ClientID != 0 && imp.UserID != "" && imp.Scope != "" {
		facs := m.Db.GetImplicitGrantByScope(imp.ClientID, imp.UserID, imp.Scope)
		if len(*facs) > 0 && (*facs)[0].ID != 0 {
			authorized = true
		}
	}
	return authorized
}

//ValidateImplicitClientAndCallback ValidateImplicitClientAndCallback
func (m *OauthManager) ValidateImplicitClientAndCallback(imp *Implicit) *ImplicitClient {
	var rtn ImplicitClient
	if imp.ClientID != 0 && imp.RedirectURI != "" {
		cru := m.Db.GetClientRedirectURI(imp.ClientID, imp.RedirectURI)
		m.Log.Debug("cru: ", cru)
		if cru.ID > 0 && cru.URI == imp.RedirectURI && cru.ClientID == imp.ClientID {
			c := m.Db.GetClient(imp.ClientID)
			if c.Enabled {
				rtn.Valid = true
				rtn.ClientName = c.Name
				rtn.WebSite = c.WebSite
			}
		}
	}
	return &rtn
}

func (m *OauthManager) processImplicitInsert(imp *Implicit, scopeStrList *[]string, existingAuthCode bool) (success bool, rtn *ImplicitReturn) {
	var igdel bool
	if existingAuthCode {
		igdel = m.Db.DeleteImplicitGrant(imp.ClientID, imp.UserID)
		m.Log.Debug("igdel: ", igdel)
	} else {
		igdel = true
	}
	if igdel {
		roleURIList := m.Db.GetClientRoleAllowedURIListByClientID(imp.ClientID)
		m.Log.Debug("roleURIList", roleURIList)
		var pl Payload
		pl.TokenType = accessTokenType
		pl.UserID = hashUser(imp.UserID)
		pl.ClientID = imp.ClientID
		pl.Subject = implicitGrantType
		pl.ExpiresInMinute = implicitAccessTokenLifeInMinutes //(600 * time.Minute) => (600 * 60) => 36000 seconds => 10 hours
		pl.Grant = implicitGrantType
		pl.RoleURIs = *m.populateRoleURLList(roleURIList)
		pl.ScopeList = *scopeStrList
		accessToken := m.GenerateAccessToken(&pl)
		m.Log.Info("accessToken: ", accessToken)
		if accessToken != "" {
			now := time.Now()
			var igt odb.ImplicitGrant
			igt.ClientID = imp.ClientID
			igt.UserID = imp.UserID

			var aToken odb.AccessToken
			aToken.Token = accessToken
			aToken.Expires = now.Add(time.Minute * implicitAccessTokenLifeInMinutes)

			igSuc, igID := m.Db.AddImplicitGrant(&igt, &aToken, scopeStrList)
			m.Log.Debug("igSuc: ", igSuc)
			m.Log.Debug("igID: ", igID)
			if igSuc {
				success = igSuc
				rtn = new(ImplicitReturn)
				rtn.ID = igID
				rtn.Token = accessToken
			}
		}
	}
	return success, rtn
}
