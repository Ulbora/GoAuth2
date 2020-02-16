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

//ValidateAccessTokenReq ValidateAccessTokenReq
type ValidateAccessTokenReq struct {
	AccessToken string `json:"accessToken"`
	Hashed      bool   `json:"hashed"`
	UserID      string `json:"userId"`
	ClientID    int64  `json:"clientId"`
	Role        string `json:"role"`
	URI         string `json:"url"`
	Scope       string `json:"scope"`
}

//ValidateAccessToken ValidateAccessToken
func (m *OauthManager) ValidateAccessToken(at *ValidateAccessTokenReq) bool {
	var rtn bool
	//fix issue with no user needed with client grant
	//fmt.Println("log level in validate token: ", m.Log.LogLevel)
	m.Log.Debug("at role", at.Role)
	if at.AccessToken != "" && at.ClientID != 0 {
		var userID string
		if at.Hashed && at.UserID != "" {
			userID = unHashUser(at.UserID)
			m.Log.Debug("unhashed user: ", userID)
		} else {
			userID = at.UserID
		}
		tkkey := m.Db.GetAccessTokenKey()
		m.Log.Debug("tkkey", tkkey)
		atsuc, pwatpl := m.ValidateJwt(at.AccessToken, tkkey)
		m.Log.Debug("atsuc", atsuc)
		m.Log.Debug("userPass", userID == unHashUser(pwatpl.UserID))
		m.Log.Debug("user", userID)
		m.Log.Debug("userUnhash", unHashUser(pwatpl.UserID))
		m.Log.Debug("pwatpl.UserID", pwatpl.UserID)
		m.Log.Info("pwatpl", *pwatpl)
		var noUser bool
		if pwatpl.UserID == "" || at.UserID == "" {
			noUser = true
		}
		m.Log.Debug("noUser", noUser)
		if m.isAcceptable(atsuc, noUser, userID, pwatpl, at) {
			// if atsuc && (noUser || userID == unHashUser(pwatpl.UserID)) && pwatpl.TokenType == accessTokenType &&
			// 	pwatpl.ClientID == at.ClientID && pwatpl.Issuer == tokenIssuer {
			m.Log.Debug("inside if")
			var roleFound bool
			var scopeFound bool
			roleFound = m.findRole(pwatpl, at)
			// if at.Role != "" && at.URI != "" {
			// 	for _, r := range pwatpl.RoleURIs {
			// 		if r.Role == at.Role && r.ClientAllowedURI == at.URI {
			// 			roleFound = true
			// 			break
			// 		}
			// 	}
			// } else {
			// 	roleFound = true
			// }
			m.Log.Debug("at.Scope", at.Scope)
			scopeFound = m.findScope(pwatpl, at)
			// if at.Scope != "" {
			// 	var foundWrite bool
			// 	for _, s := range pwatpl.ScopeList {
			// 		if s == "write" {
			// 			foundWrite = true
			// 		}
			// 		if s == at.Scope {
			// 			scopeFound = true
			// 			//break
			// 		}
			// 	}
			// 	if at.Scope == "read" && foundWrite {
			// 		scopeFound = true
			// 	}
			// } else {
			// 	for _, s := range pwatpl.ScopeList {
			// 		if s == "write" {
			// 			scopeFound = true
			// 			break
			// 		}
			// 	}
			// }
			m.Log.Debug("roleFound", roleFound)
			m.Log.Debug("scopeFound", scopeFound)
			if (pwatpl.Grant == codeGrantType || pwatpl.Grant == implicitGrantType) && roleFound && scopeFound {
				rtn = true
			} else if (pwatpl.Grant == clientGrantType || pwatpl.Grant == passwordGrantType) && roleFound {
				rtn = true
			}
		}
	}
	return rtn
}

func (m *OauthManager) findRole(pwatpl *Payload, at *ValidateAccessTokenReq) bool {
	var roleFound bool
	if at.Role != "" && at.URI != "" {
		for _, r := range pwatpl.RoleURIs {
			if r.Role == at.Role && r.ClientAllowedURI == at.URI {
				roleFound = true
				break
			}
		}
	} else {
		roleFound = true
	}
	return roleFound
}

func (m *OauthManager) findScope(pwatpl *Payload, at *ValidateAccessTokenReq) bool {
	var scopeFound bool
	if at.Scope != "" {
		var foundWrite bool
		for _, s := range pwatpl.ScopeList {
			if s == "write" {
				foundWrite = true
			}
			if s == at.Scope {
				scopeFound = true
				//break
			}
		}
		if at.Scope == "read" && foundWrite {
			scopeFound = true
		}
	} else {
		for _, s := range pwatpl.ScopeList {
			if s == "write" {
				scopeFound = true
				break
			}
		}
	}
	return scopeFound
}

func (m *OauthManager) isAcceptable(atsuc bool, noUser bool, userID string, pwatpl *Payload, at *ValidateAccessTokenReq) bool {
	var rtn bool
	if atsuc && (noUser || userID == unHashUser(pwatpl.UserID)) && pwatpl.TokenType == accessTokenType &&
		pwatpl.ClientID == at.ClientID && pwatpl.Issuer == tokenIssuer {
		rtn = true
	}
	return rtn
}
