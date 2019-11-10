package managers

import (
	"fmt"
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

//AuthCodeTokenReq AuthCodeTokenReq
type AuthCodeTokenReq struct {
	ClientID    int64
	Secret      string
	Code        string
	RedirectURI string
}

//CredentialsTokenReq CredentialsTokenReq
type CredentialsTokenReq struct {
	ClientID int64
	Secret   string
}

//RefreshTokenReq RefreshTokenReq
type RefreshTokenReq struct {
	ClientID     int64
	Secret       string
	RefreshToken string
}

//PasswordTokenReq PasswordTokenReq
type PasswordTokenReq struct {
	Username string
	Password string
	ClientID int64
}

//Token Token
type Token struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int64
	RefreshToken string
}

//GetAuthCodeToken GetAuthCodeToken
func (m *OauthManager) GetAuthCodeToken(act *AuthCodeTokenReq) (bool, *Token) {
	var rtn Token
	var suc bool
	client := m.Db.GetClient(act.ClientID)
	fmt.Println("client: ", client)
	if client != nil && client.Secret == act.Secret && client.Enabled {
		rtu := m.Db.GetClientRedirectURI(act.ClientID, act.RedirectURI)
		fmt.Println("rtu: ", rtu)
		if rtu.ID > 0 {
			acode := m.Db.GetAuthorizationCodeByCode(act.Code)
			if acode.ClientID == act.ClientID {
				acRev := m.Db.GetAuthCodeRevolk(acode.AuthorizationCode)
				fmt.Println("acRev: ", acRev)
				if acRev == nil || acRev.ID == 0 {
					if acode.AlreadyUsed {
						fmt.Println("AlreadyUsed: ", acode.AlreadyUsed)
						var rvk odb.AuthCodeRevolk
						rvk.AuthorizationCode = acode.AuthorizationCode
						rvsuc, rvid := m.Db.AddAuthCodeRevolk(nil, &rvk)
						fmt.Println("rvsuc: ", rvsuc)
						fmt.Println("rvid: ", rvid)
					} else {
						acode.AlreadyUsed = true
						usuc := m.Db.UpdateAuthorizationCode(acode)
						fmt.Println("usuc: ", usuc)
						if usuc {
							tkn := m.Db.GetAccessToken(acode.AccessTokenID)
							if tkn.ID > 0 {
								fmt.Println("tkn: ", tkn)
								rtn.AccessToken = tkn.Token
								rtn.TokenType = tokenTypeBearer
								rtn.ExpiresIn = codeAccessTokenLifeInMinutes * 60
								if tkn.RefreshTokenID != 0 {
									rtkn := m.Db.GetRefreshToken(tkn.RefreshTokenID)
									fmt.Println("rtkn: ", rtkn)
									if rtkn.ID > 0 {
										rtn.RefreshToken = rtkn.Token
										suc = true
									}
								} else {
									suc = true
								}
							}
						}
					}
				}
			}
		}
	}
	return suc, &rtn
}

//GetCredentialsToken GetCredentialsToken
func (m *OauthManager) GetCredentialsToken(ct *CredentialsTokenReq) (bool, *Token) {
	var rtn Token
	var suc bool
	client := m.Db.GetClient(ct.ClientID)
	fmt.Println("client: ", client)
	if client != nil && client.Secret == ct.Secret && client.Enabled {

		gton := m.grantTypeTurnedOn(ct.ClientID, clientGrantType)
		fmt.Println("gton: ", gton)
		if gton {
			delSuc := m.Db.DeleteCredentialsGrant(ct.ClientID)
			fmt.Println("delSuc: ", delSuc)
			if delSuc {
				roleURIList := m.Db.GetClientRoleAllowedURIListByClientID(ct.ClientID)
				fmt.Println("roleURIList", roleURIList)
				var pl Payload
				pl.TokenType = accessTokenType
				//pl.UserID = hashUser(ac.UserID)
				pl.ClientID = ct.ClientID
				pl.Subject = clientGrantType
				pl.ExpiresInMinute = credentialsGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
				pl.Grant = clientGrantType
				pl.RoleURIs = *m.populateRoleURLList(roleURIList)
				//pl.ScopeList = *scopeStrList
				accessToken := m.GenerateAccessToken(&pl)
				fmt.Println("accessToken: ", accessToken)
				if accessToken != "" {
					now := time.Now()
					var aToken odb.AccessToken
					aToken.Token = accessToken
					aToken.Expires = now.Add(time.Minute * codeAccessTokenLifeInMinutes)

					var cgrant odb.CredentialsGrant
					cgrant.ClientID = ct.ClientID

					cgSuc, _ := m.Db.AddCredentialsGrant(&cgrant, &aToken)
					fmt.Println("cgSuc: ", cgSuc)
					if cgSuc {
						rtn.AccessToken = accessToken
						rtn.TokenType = tokenTypeBearer
						rtn.ExpiresIn = credentialsGrantAccessTokenLifeInMinutes * 60
						suc = true
					}
				}
			}
		}
	}
	return suc, &rtn
}
