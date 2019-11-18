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
	// Password string
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

//GetPasswordToken GetPasswordToken
func (m *OauthManager) GetPasswordToken(pt *PasswordTokenReq) (bool, *Token) {
	var rtn Token
	var suc bool
	client := m.Db.GetClient(pt.ClientID)
	fmt.Println("pw client: ", client)
	if client != nil && client.Enabled {

		gton := m.grantTypeTurnedOn(pt.ClientID, passwordGrantType)
		fmt.Println("pw gton: ", gton)
		if gton {
			delSuc := m.Db.DeletePasswordGrant(pt.ClientID, pt.Username)
			fmt.Println("delSuc: ", delSuc)
			if delSuc {
				roleURIList := m.Db.GetClientRoleAllowedURIListByClientID(pt.ClientID)
				fmt.Println("roleURIList", roleURIList)
				var pl Payload
				pl.TokenType = accessTokenType
				pl.UserID = hashUser(pt.Username)
				pl.ClientID = pt.ClientID
				pl.Subject = passwordGrantType
				pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
				pl.Grant = passwordGrantType
				pl.RoleURIs = *m.populateRoleURLList(roleURIList)
				//pl.ScopeList = *scopeStrList
				accessToken := m.GenerateAccessToken(&pl)
				fmt.Println("accessToken: ", accessToken)
				if accessToken != "" {
					refToken := m.GenerateRefreshToken(pt.ClientID, hashUser(pt.Username), passwordGrantType)
					fmt.Println("refToken: ", refToken)

					now := time.Now()
					var aToken odb.AccessToken
					aToken.Token = accessToken
					aToken.Expires = now.Add(time.Minute * passwordGrantAccessTokenLifeInMinutes)

					var rToken odb.RefreshToken
					rToken.Token = refToken

					var pgrant odb.PasswordGrant
					pgrant.ClientID = pt.ClientID
					pgrant.UserID = pt.Username

					cgSuc, _ := m.Db.AddPasswordGrant(&pgrant, &aToken, &rToken)
					fmt.Println("cgSuc: ", cgSuc)
					if cgSuc {
						rtn.AccessToken = accessToken
						rtn.TokenType = tokenTypeBearer
						rtn.ExpiresIn = passwordGrantAccessTokenLifeInMinutes * 60
						rtn.RefreshToken = refToken
						suc = true
					}
				}
			}
		}
	}
	return suc, &rtn
}

//GetAuthCodeAccesssTokenWithRefreshToken GetAuthCodeAccesssTokenWithRefreshToken
func (m *OauthManager) GetAuthCodeAccesssTokenWithRefreshToken(rt *RefreshTokenReq) (bool, *Token) {
	var rtn Token
	var suc bool
	if rt.ClientID != 0 && rt.Secret != "" {
		client := m.Db.GetClient(rt.ClientID)
		fmt.Println("client in get with ref: ", client)
		if client.Enabled && client.Secret == rt.Secret {
			fmt.Println("client enabled and secrets match")
			rtk := m.Db.GetRefreshTokenKey()
			if rtk != "" {
				fmt.Println("refresh Token Key", rtk)
				rtsuc, rtpl := m.ValidateJwt(rt.RefreshToken, rtk)
				fmt.Println("rtsuc", rtsuc)
				fmt.Println("rtpl", rtpl)
				if rtsuc && rtpl.ClientID == rt.ClientID && rtpl.Subject == codeGrantType {
					fmt.Println("rtpl in success", rtpl)
					fmt.Println("unhashed user", unHashUser(rtpl.UserID))
					acode := m.Db.GetAuthorizationCode(rt.ClientID, unHashUser(rtpl.UserID))
					fmt.Println("acode", acode)
					fmt.Println("acode user", (*acode)[0].UserID)
					fmt.Println("acode AccessTokenID", (*acode)[0].AccessTokenID)
					if len(*acode) > 0 && (*acode)[0].UserID == unHashUser(rtpl.UserID) {
						fmt.Println("acode user suc")
						atkn := m.Db.GetAccessToken((*acode)[0].AccessTokenID)
						//fmt.Println("atkn", atkn)
						if atkn.ID > 0 {
							fmt.Println("atkn", atkn)
							tkkey := m.Db.GetAccessTokenKey()
							fmt.Println("tkkey", tkkey)
							atsuc, atpl := m.ValidateJwt(atkn.Token, tkkey)
							fmt.Println("atsuc", atsuc)
							fmt.Println("atpl", atpl)
							if atpl.UserID == rtpl.UserID && atpl.ClientID == rt.ClientID {
								fmt.Println("atpl in success", atpl)
								var pl Payload
								pl.TokenType = accessTokenType
								pl.UserID = atpl.UserID
								pl.ClientID = rt.ClientID
								pl.Subject = codeGrantType
								pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
								pl.Grant = codeGrantType
								pl.RoleURIs = atpl.RoleURIs
								pl.ScopeList = atpl.ScopeList
								newAccessToken := m.GenerateAccessToken(&pl)
								fmt.Println("newAccessToken", newAccessToken)
								now := time.Now()
								(*acode)[0].Expires = now.Add(time.Minute * authCodeLifeInMinutes)
								atkn.Token = newAccessToken
								atkn.Expires = now.Add(time.Minute * codeAccessTokenLifeInMinutes)
								suc = m.Db.UpdateAuthorizationCodeAndToken(&(*acode)[0], atkn)
								rtn.AccessToken = newAccessToken
								rtn.TokenType = tokenTypeBearer
								rtn.ExpiresIn = codeAccessTokenLifeInMinutes * 60
								rtn.RefreshToken = rt.RefreshToken
							}
						}
					}
				}
			}
		}
	}
	return suc, &rtn
}

//GetPasswordAccesssTokenWithRefreshToken GetPasswordAccesssTokenWithRefreshToken
func (m *OauthManager) GetPasswordAccesssTokenWithRefreshToken(rt *RefreshTokenReq) (bool, *Token) {
	var rtn Token
	var suc bool
	if rt.ClientID != 0 {
		client := m.Db.GetClient(rt.ClientID)
		fmt.Println("client in get with ref: ", client)
		rtk := m.Db.GetRefreshTokenKey()
		if rtk != "" {
			fmt.Println("refresh Token Key", rtk)
			rtsuc, rtpl := m.ValidateJwt(rt.RefreshToken, rtk)
			fmt.Println("rtsuc", rtsuc)
			fmt.Println("rtpl", rtpl)
			if rtsuc && rtpl.ClientID == rt.ClientID && rtpl.Subject == passwordGrantType {
				fmt.Println("rtpl in success", rtpl)
				fmt.Println("unhashed user", unHashUser(rtpl.UserID))
				pgnt := m.Db.GetPasswordGrant(rt.ClientID, unHashUser(rtpl.UserID))
				fmt.Println("pgnt", pgnt)
				fmt.Println("pgnt user", (*pgnt)[0].UserID)
				fmt.Println("pgnt AccessTokenID", (*pgnt)[0].AccessTokenID)
				if len(*pgnt) > 0 && (*pgnt)[0].UserID == unHashUser(rtpl.UserID) {
					fmt.Println("pgnt user suc")
					atkn := m.Db.GetAccessToken((*pgnt)[0].AccessTokenID)
					fmt.Println("atkn", atkn)
					if atkn.ID > 0 {
						fmt.Println("atkn", atkn)
						tkkey := m.Db.GetAccessTokenKey()
						fmt.Println("tkkey", tkkey)
						atsuc, atpl := m.ValidateJwt(atkn.Token, tkkey)
						fmt.Println("atsuc", atsuc)
						fmt.Println("atpl", atpl)
						if atpl.UserID == rtpl.UserID && atpl.ClientID == rt.ClientID {
							fmt.Println("atpl in success", atpl)
							var pl Payload
							pl.TokenType = accessTokenType
							pl.UserID = atpl.UserID
							pl.ClientID = rt.ClientID
							pl.Subject = passwordGrantType
							pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
							pl.Grant = passwordGrantType
							pl.RoleURIs = atpl.RoleURIs
							pl.ScopeList = atpl.ScopeList
							newAccessToken := m.GenerateAccessToken(&pl)
							fmt.Println("newAccessToken", newAccessToken)
							now := time.Now()
							//(*pgnt)[0].Expires = now.Add(time.Minute * authCodeLifeInMinutes)
							atkn.Token = newAccessToken
							atkn.Expires = now.Add(time.Minute * passwordGrantAccessTokenLifeInMinutes)
							suc = m.Db.UpdateAccessToken(nil, atkn)
							rtn.AccessToken = newAccessToken
							rtn.TokenType = tokenTypeBearer
							rtn.ExpiresIn = passwordGrantAccessTokenLifeInMinutes * 60
							rtn.RefreshToken = rt.RefreshToken
						}
					}
				}
			}
		}
	}
	return suc, &rtn
}
