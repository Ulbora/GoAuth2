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

//"fmt"

//"fmt"

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
			fmt.Println("grant turned on: ", gton)
			if gton {
				acode := m.Db.GetAuthorizationCode(ac.ClientID, ac.UserID)
				fmt.Println("acode: ", acode)
				if len(*acode) > 0 && (*acode)[0].AuthorizationCode != 0 {
					scopeList := m.Db.GetAuthorizationCodeScopeList((*acode)[0].AuthorizationCode)
					fmt.Println("scopeList: ", scopeList)
					var scopeFound bool
					for _, s := range *scopeList {
						if s.Scope == ac.Scope {
							scopeFound = true
							break
						}
					}
					fmt.Println("scopeFound: ", scopeFound)
					if scopeFound {
						acdel := m.Db.DeleteAuthorizationCode(ac.ClientID, ac.UserID)
						fmt.Println("acdel: ", acdel)
						if acdel {

						}
					} else {

					}
				}
			}

		}
	}

	return success, authCode, authCodeString
}
