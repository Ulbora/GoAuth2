package managers

import "fmt"

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
	URI         string `json:"uri"`
	Scope       string `json:"scope"`
}

//ValidateAccessToken ValidateAccessToken
func (m *OauthManager) ValidateAccessToken(at *ValidateAccessTokenReq) bool {
	var rtn bool
	if at.AccessToken != "" && at.UserID != "" && at.ClientID != 0 {
		var userID string
		if at.Hashed {
			userID = unHashUser(at.UserID)
			fmt.Println("unhashed user: ", userID)
		} else {
			userID = at.UserID
		}
		tkkey := m.Db.GetAccessTokenKey()
		fmt.Println("tkkey", tkkey)
		atsuc, pwatpl := m.ValidateJwt(at.AccessToken, tkkey)
		fmt.Println("atsuc", atsuc)
		fmt.Println("pwatpl", pwatpl)
		if atsuc && userID == unHashUser(pwatpl.UserID) && pwatpl.TokenType == accessTokenType &&
			pwatpl.ClientID == at.ClientID && pwatpl.Issuer == tokenIssuer {
			fmt.Println("inside if")
			var roleFound bool
			var scopeFound bool
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
			if at.Scope != "" {
				for _, s := range pwatpl.ScopeList {
					if s == at.Scope {
						scopeFound = true
						break
					}
				}
			} else {
				for _, s := range pwatpl.ScopeList {
					if s == "write" {
						scopeFound = true
						break
					}
				}
			}
			if (pwatpl.Grant == codeGrantType || pwatpl.Grant == implicitGrantType) && roleFound && scopeFound {
				rtn = true
			} else if (pwatpl.Grant == clientGrantType || pwatpl.Grant == passwordGrantType) && roleFound {
				rtn = true
			}
		}
	}
	return rtn
}
