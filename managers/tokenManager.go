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
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

const (
	invalidClientError  = "invalid_client"
	invalidGrantError   = "invalid_grant"
	accessDeniedError   = "access_denied"
	invalidRequestError = "invalid_request"
)

//GetAuthCodeToken GetAuthCodeToken
func (m *OauthManager) GetAuthCodeToken(act *AuthCodeTokenReq) (bool, *Token, string) {
	var rtn Token
	var suc bool
	var tokenErr string
	client := m.Db.GetClient(act.ClientID)
	m.Log.Debug("client: ", client)
	if client != nil && client.Secret == act.Secret && client.Enabled {
		rtu := m.Db.GetClientRedirectURI(act.ClientID, act.RedirectURI)
		m.Log.Debug("rtu: ", rtu)
		if rtu.ID > 0 {
			acode := m.Db.GetAuthorizationCodeByCode(act.Code)
			if acode.ClientID == act.ClientID {
				acRev := m.Db.GetAuthCodeRevolk(acode.AuthorizationCode)
				m.Log.Debug("acRev: ", acRev)
				if acRev == nil || acRev.ID == 0 {
					if acode.AlreadyUsed {
						m.Log.Debug("AlreadyUsed: ", acode.AlreadyUsed)
						var rvk odb.AuthCodeRevolk
						rvk.AuthorizationCode = acode.AuthorizationCode
						rvsuc, rvid := m.Db.AddAuthCodeRevolk(nil, &rvk)
						tokenErr = invalidClientError
						m.Log.Debug("rvsuc: ", rvsuc)
						m.Log.Debug("rvid: ", rvid)
					} else {
						acode.AlreadyUsed = true
						usuc := m.Db.UpdateAuthorizationCode(acode)
						m.Log.Debug("usuc: ", usuc)
						if usuc {
							tkn := m.Db.GetAccessToken(acode.AccessTokenID)
							if tkn.ID > 0 {
								m.Log.Info("tkn: ", tkn)
								rtn.AccessToken = tkn.Token
								rtn.TokenType = tokenTypeBearer
								rtn.ExpiresIn = codeAccessTokenLifeInMinutes * 60
								if tkn.RefreshTokenID != 0 {
									rtkn := m.Db.GetRefreshToken(tkn.RefreshTokenID)
									m.Log.Info("rtkn: ", rtkn)
									if rtkn.ID > 0 {
										rtn.RefreshToken = rtkn.Token
										suc = true
									}
								} else {
									suc = true
								}
							} else {
								tokenErr = invalidGrantError
							}
						} else {
							tokenErr = invalidGrantError
						}
					}
				} else {
					tokenErr = invalidClientError
				}
			} else {
				tokenErr = invalidClientError
			}
		} else {
			tokenErr = invalidGrantError
		}
	} else {
		tokenErr = invalidClientError
	}
	return suc, &rtn, tokenErr
}

//GetCredentialsToken GetCredentialsToken
func (m *OauthManager) GetCredentialsToken(ct *CredentialsTokenReq) (bool, *Token, string) {
	var rtn Token
	var suc bool
	var tokenErr string
	client := m.Db.GetClient(ct.ClientID)
	m.Log.Debug("client: ", client)
	if client != nil && client.Secret == ct.Secret && client.Enabled {

		gton := m.grantTypeTurnedOn(ct.ClientID, clientGrantType)
		m.Log.Debug("gton: ", gton)
		if gton {
			delSuc := m.Db.DeleteCredentialsGrant(ct.ClientID)
			m.Log.Debug("delSuc: ", delSuc)
			if delSuc {
				roleURIList := m.Db.GetClientRoleAllowedURIListByClientID(ct.ClientID)
				m.Log.Debug("roleURIList", roleURIList)
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
				m.Log.Info("accessToken: ", accessToken)
				if accessToken != "" {
					now := time.Now()
					var aToken odb.AccessToken
					aToken.Token = accessToken
					aToken.Expires = now.Add(time.Minute * codeAccessTokenLifeInMinutes)

					var cgrant odb.CredentialsGrant
					cgrant.ClientID = ct.ClientID

					cgSuc, _ := m.Db.AddCredentialsGrant(&cgrant, &aToken)
					m.Log.Debug("cgSuc: ", cgSuc)
					if cgSuc {
						rtn.AccessToken = accessToken
						rtn.TokenType = tokenTypeBearer
						rtn.ExpiresIn = credentialsGrantAccessTokenLifeInMinutes * 60
						suc = true
					}
				}
			}
		}
	} else {
		tokenErr = invalidClientError
	}
	if !suc && tokenErr == "" {
		tokenErr = accessDeniedError
	}
	return suc, &rtn, tokenErr
}

//GetPasswordToken GetPasswordToken
func (m *OauthManager) GetPasswordToken(pt *PasswordTokenReq) (bool, *Token, string) {
	var rtn Token
	var suc bool
	var tokenErr string
	client := m.Db.GetClient(pt.ClientID)
	m.Log.Debug("pw client: ", client)
	if client != nil && client.Enabled {

		gton := m.grantTypeTurnedOn(pt.ClientID, passwordGrantType)
		m.Log.Debug("pw gton: ", gton)
		if gton {
			delSuc := m.Db.DeletePasswordGrant(pt.ClientID, pt.Username)
			m.Log.Debug("delSuc: ", delSuc)
			if delSuc {
				roleURIList := m.Db.GetClientRoleAllowedURIListByClientID(pt.ClientID)
				m.Log.Debug("roleURIList", roleURIList)
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
				m.Log.Info("accessToken: ", accessToken)
				if accessToken != "" {
					refToken := m.GenerateRefreshToken(pt.ClientID, hashUser(pt.Username), passwordGrantType)
					m.Log.Info("refToken: ", refToken)

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
					m.Log.Debug("cgSuc: ", cgSuc)
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
	} else {
		tokenErr = invalidClientError
	}
	if !suc && tokenErr == "" {
		tokenErr = accessDeniedError
	}
	return suc, &rtn, tokenErr
}

//GetAuthCodeAccesssTokenWithRefreshToken GetAuthCodeAccesssTokenWithRefreshToken
func (m *OauthManager) GetAuthCodeAccesssTokenWithRefreshToken(rt *RefreshTokenReq) (bool, *Token, string) {
	var rtn Token
	var suc bool
	var tokenErr string
	if rt.ClientID != 0 && rt.Secret != "" {
		client := m.Db.GetClient(rt.ClientID)
		m.Log.Debug("client in get with ref: ", client)
		if client.Enabled && client.Secret == rt.Secret {
			m.Log.Debug("client enabled and secrets match")
			rtk := m.Db.GetRefreshTokenKey()
			if rtk != "" {
				m.Log.Info("refresh Token Key", rtk)
				rtsuc, rtpl := m.ValidateJwt(rt.RefreshToken, rtk)
				m.Log.Debug("rtsuc", rtsuc)
				m.Log.Debug("rtpl", rtpl)
				if rtsuc && rtpl.ClientID == rt.ClientID && rtpl.Subject == codeGrantType {
					m.Log.Debug("rtpl in success", rtpl)
					m.Log.Debug("unhashed user", unHashUser(rtpl.UserID))
					acode := m.Db.GetAuthorizationCode(rt.ClientID, unHashUser(rtpl.UserID))
					m.Log.Info("acode", acode)
					m.Log.Debug("acode user", (*acode)[0].UserID)
					m.Log.Info("acode AccessTokenID", (*acode)[0].AccessTokenID)
					if len(*acode) > 0 && (*acode)[0].UserID == unHashUser(rtpl.UserID) {
						m.Log.Debug("acode user suc")
						atkn := m.Db.GetAccessToken((*acode)[0].AccessTokenID)
						//fmt.Println("atkn", atkn)
						if atkn.ID > 0 {
							m.Log.Debug("atkn", atkn)
							tkkey := m.Db.GetAccessTokenKey()
							m.Log.Info("tkkey", tkkey)
							atsuc, atpl := m.ValidateJwt(atkn.Token, tkkey)
							m.Log.Debug("atsuc", atsuc)
							m.Log.Debug("atpl", atpl)
							if atpl.UserID == rtpl.UserID && atpl.ClientID == rt.ClientID {
								m.Log.Debug("atpl in success", atpl)
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
								m.Log.Info("newAccessToken", newAccessToken)
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
	} else {
		tokenErr = invalidRequestError
	}
	if !suc && tokenErr == "" {
		tokenErr = invalidClientError
	}
	return suc, &rtn, tokenErr
}

//GetPasswordAccesssTokenWithRefreshToken GetPasswordAccesssTokenWithRefreshToken
func (m *OauthManager) GetPasswordAccesssTokenWithRefreshToken(rt *RefreshTokenReq) (bool, *Token, string) {
	var rtn Token
	var suc bool
	var tokenErr string
	if rt.ClientID != 0 {
		client := m.Db.GetClient(rt.ClientID)
		m.Log.Debug("client in get with ref: ", client)
		rtk := m.Db.GetRefreshTokenKey()
		if rtk != "" {
			m.Log.Info("refresh Token Key", rtk)
			rtsuc, rtpl := m.ValidateJwt(rt.RefreshToken, rtk)
			m.Log.Debug("rtsuc", rtsuc)
			m.Log.Debug("rtpl", rtpl)
			if rtsuc && rtpl.ClientID == rt.ClientID && rtpl.Subject == passwordGrantType {
				m.Log.Debug("rtpl in success", rtpl)
				m.Log.Debug("unhashed user", unHashUser(rtpl.UserID))
				pgnt := m.Db.GetPasswordGrant(rt.ClientID, unHashUser(rtpl.UserID))
				m.Log.Debug("pgnt", pgnt)
				m.Log.Debug("pgnt user", (*pgnt)[0].UserID)
				m.Log.Debug("pgnt AccessTokenID", (*pgnt)[0].AccessTokenID)
				if len(*pgnt) > 0 && (*pgnt)[0].UserID == unHashUser(rtpl.UserID) {
					m.Log.Debug("pgnt user suc")
					pgatkn := m.Db.GetAccessToken((*pgnt)[0].AccessTokenID)
					m.Log.Debug("pgatkn", pgatkn)
					if pgatkn.ID > 0 {
						m.Log.Debug("pgatkn", pgatkn)
						tkkey := m.Db.GetAccessTokenKey()
						m.Log.Info("tkkey", tkkey)
						atsuc, pwatpl := m.ValidateJwt(pgatkn.Token, tkkey)
						m.Log.Debug("atsuc", atsuc)
						m.Log.Debug("pwatpl", pwatpl)
						if pwatpl.UserID == rtpl.UserID && pwatpl.ClientID == rt.ClientID {
							m.Log.Debug("pwatpl in success", pwatpl)
							var pl Payload
							pl.TokenType = accessTokenType
							pl.UserID = pwatpl.UserID
							pl.ClientID = rt.ClientID
							pl.Subject = passwordGrantType
							pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
							pl.Grant = passwordGrantType
							pl.RoleURIs = pwatpl.RoleURIs
							pl.ScopeList = pwatpl.ScopeList
							newAccessToken := m.GenerateAccessToken(&pl)
							m.Log.Info("newAccessToken", newAccessToken)
							now := time.Now()
							//(*pgnt)[0].Expires = now.Add(time.Minute * authCodeLifeInMinutes)
							pgatkn.Token = newAccessToken
							pgatkn.Expires = now.Add(time.Minute * passwordGrantAccessTokenLifeInMinutes)
							suc = m.Db.UpdateAccessToken(nil, pgatkn)
							rtn.AccessToken = newAccessToken
							rtn.TokenType = tokenTypeBearer
							rtn.ExpiresIn = passwordGrantAccessTokenLifeInMinutes * 60
							rtn.RefreshToken = rt.RefreshToken
						}
					}
				}
			}
		}
	} else {
		tokenErr = invalidRequestError
	}
	if !suc && tokenErr == "" {
		tokenErr = invalidClientError
	}
	return suc, &rtn, tokenErr
}
