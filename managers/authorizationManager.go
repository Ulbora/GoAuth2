//Package managers ...
package managers

import (
	"time"

	"github.com/Ulbora/GoAuth2/oauth2database"
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

//AuthCode AuthCode
type AuthCode struct {
	ClientID    int64
	UserID      string
	Scope       string
	RedirectURI string
	CallbackURI string
}

//AuthCodeClient AuthCodeClient
type AuthCodeClient struct {
	Valid      bool
	ClientName string
	WebSite    string
}

//AuthorizeAuthCode AuthorizeAuthCode
func (m *OauthManager) AuthorizeAuthCode(ac *AuthCode) (success bool, authCode int64, authCodeString string) {
	client := m.Db.GetClient(ac.ClientID)
	if client.Enabled {
		rtu := m.Db.GetClientRedirectURI(ac.ClientID, ac.RedirectURI)
		if rtu.ID > 0 {
			gton := m.grantTypeTurnedOn(ac.ClientID, codeGrantType)
			m.Log.Debug("grant turned on: ", gton)
			if gton {
				acode := m.Db.GetAuthorizationCode(ac.ClientID, ac.UserID)
				m.Log.Debug("acode: ", acode)
				var scopeStrList []string
				if len(*acode) > 0 && (*acode)[0].AuthorizationCode != 0 {
					scopeList := m.Db.GetAuthorizationCodeScopeList((*acode)[0].AuthorizationCode)
					m.Log.Debug("scopeList: ", scopeList)
					var scopeFound bool
					for _, s := range *scopeList {
						if s.Scope == ac.Scope {
							scopeFound = true
							break
						}
					}
					m.Log.Debug("scopeFound: ", scopeFound)
					for _, s := range *scopeList {
						scopeStrList = append(scopeStrList, s.Scope)
					}
					if scopeFound {
						success, authCode, authCodeString = m.processAuthCodeInsert(ac, &scopeStrList, true)
					} else {
						scopeStrList = append(scopeStrList, ac.Scope)
						success, authCode, authCodeString = m.processAuthCodeInsert(ac, &scopeStrList, true)
					}
				} else {
					scopeStrList = append(scopeStrList, ac.Scope)
					success, authCode, authCodeString = m.processAuthCodeInsert(ac, &scopeStrList, false)
				}
			}
		}
	}

	return success, authCode, authCodeString
}

//CheckAuthCodeApplicationAuthorization CheckAuthCodeApplicationAuthorization
func (m *OauthManager) CheckAuthCodeApplicationAuthorization(ac *AuthCode) (authorized bool) {
	if ac.ClientID != 0 && ac.UserID != "" && ac.Scope != "" {
		facs := m.Db.GetAuthorizationCodeByScope(ac.ClientID, ac.UserID, ac.Scope)
		if len(*facs) > 0 && (*facs)[0].AuthorizationCode != 0 {
			authorized = true
		}
	}
	return authorized
}

//ValidateAuthCodeClientAndCallback ValidateAuthCodeClientAndCallback
func (m *OauthManager) ValidateAuthCodeClientAndCallback(ac *AuthCode) *AuthCodeClient {
	var rtn AuthCodeClient
	if ac.ClientID != 0 && ac.RedirectURI != "" {
		cru := m.Db.GetClientRedirectURI(ac.ClientID, ac.RedirectURI)
		m.Log.Debug("cru: ", cru)
		if cru.ID > 0 && cru.URI == ac.RedirectURI && cru.ClientID == ac.ClientID {
			c := m.Db.GetClient(ac.ClientID)
			if c.Enabled {
				rtn.Valid = true
				rtn.ClientName = c.Name
				rtn.WebSite = c.WebSite
			}
		}
	}
	return &rtn
}

func (m *OauthManager) processAuthCodeInsert(ac *AuthCode, scopeStrList *[]string, existingAuthCode bool) (success bool, authCode int64, authCodeString string) {
	var acdel bool
	if existingAuthCode {
		acdel = m.Db.DeleteAuthorizationCode(ac.ClientID, ac.UserID)
		m.Log.Debug("acdel: ", acdel)
	} else {
		acdel = true
	}
	if acdel {
		refToken := m.GenerateRefreshToken(ac.ClientID, hashUser(ac.UserID), codeGrantType)
		m.Log.Info("refToken:", refToken)
		if refToken != "" {
			roleURIList := m.Db.GetClientRoleAllowedURIListByClientID(ac.ClientID)
			m.Log.Debug("roleURIList", roleURIList)
			var pl Payload
			pl.TokenType = accessTokenType
			pl.UserID = hashUser(ac.UserID)
			pl.ClientID = ac.ClientID
			pl.Subject = codeGrantType
			pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
			pl.Grant = codeGrantType
			pl.RoleURIs = *m.populateRoleURLList(roleURIList)
			pl.ScopeList = *scopeStrList
			accessToken := m.GenerateAccessToken(&pl)
			m.Log.Info("accessToken: ", accessToken)
			if accessToken != "" {
				var code odb.AuthorizationCode
				code.ClientID = ac.ClientID
				code.UserID = ac.UserID
				code.RandonAuthCode = generateRandonAuthCode()
				now := time.Now()
				code.Expires = now.Add(time.Minute * authCodeLifeInMinutes)

				var aToken odb.AccessToken
				aToken.Token = accessToken
				aToken.Expires = now.Add(time.Minute * codeAccessTokenLifeInMinutes)

				var rToken odb.RefreshToken
				rToken.Token = refToken
				acSuc, acID := m.Db.AddAuthorizationCode(&code, &aToken, &rToken, scopeStrList)
				m.Log.Debug("acSuc: ", acSuc)
				m.Log.Debug("acID: ", acID)
				if acSuc {
					newRanCode := generateAuthCodeString(acID, code.RandonAuthCode)
					var uac odb.AuthorizationCode
					uac.AuthorizationCode = acID
					uac.RandonAuthCode = newRanCode
					usuc := m.Db.UpdateAuthorizationCode(&uac)
					m.Log.Debug("update success: ", usuc)
					if usuc {
						success = usuc
						authCode = acID
						authCodeString = newRanCode
					}
				}
			}
		}
	}
	return success, authCode, authCodeString
}

func (m *OauthManager) populateRoleURLList(rl *[]oauth2database.RoleURI) *[]RoleURI {
	var rtn []RoleURI
	for _, r := range *rl {
		var ru RoleURI
		ru.ClientRoleID = r.ClientRoleID
		ru.Role = r.Role
		ru.ClientAllowedURIID = r.ClientAllowedURIID
		ru.ClientAllowedURI = r.ClientAllowedURI
		ru.ClientID = r.ClientID
		rtn = append(rtn, ru)
	}
	return &rtn
}
