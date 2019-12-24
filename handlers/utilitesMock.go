//Package handlers ...
package handlers

import (
	m "github.com/Ulbora/GoAuth2/managers"
	oa "github.com/Ulbora/GoAuth2/oauthclient"
	rc "github.com/Ulbora/GoAuth2/rolecontrol"
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

//UseMockWeb UseMockWeb
func UseMockWeb() *OauthWebHandler {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockImplicitAuthorized = true
	var acc m.AuthCodeClient
	//acc.Valid = true
	acc.ClientName = "Test Client"
	acc.WebSite = "www.TestClient.com"
	om.MockAuthCodeClient = acc

	var ic m.ImplicitClient
	ic.Valid = true
	ic.ClientName = "Test Client"
	ic.WebSite = "www.TestClient.com"
	om.MockImplicitClient = ic
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockUserLoginSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	om.MockImplicitAuthorizeSuccess = true
	var ir m.ImplicitReturn
	ir.ID = 3
	ir.Token = "12345"
	om.MockImplicitReturn = ir

	var tkn m.Token
	tkn.AccessToken = "65165165"
	tkn.TokenType = "Bearer"
	tkn.RefreshToken = "16161"
	tkn.ExpiresIn = 50000

	om.MockAuthCodeTokenSuccess = true
	om.MockCredentialsTokenSuccess = true
	om.MockAuthCodeRefreshTokenSuccess = true
	om.MockPasswordTokenSuccess = true
	om.MockToken = tkn
	om.MockTokenError = "token failed"

	var wh OauthWebHandler
	wh.Manager = &om
	return &wh
}

//UseMockRest UseMockRest
func UseMockRest() *OauthRestHandler {
	var om m.MockManager
	om.MockInsertSuccess1 = true
	om.MockInsertID1 = 34
	om.MockUpdateSuccess1 = true
	om.MockDeleteSuccess1 = true

	var mc m.Client
	mc.ClientID = 510
	mc.Secret = "12345"
	mc.Name = "test client"
	mc.WebSite = "www.testclient.com"
	mc.Email = "tester@testclient.com"
	mc.Enabled = true

	var cuo m.ClientRedirectURI
	cuo.ID = 4
	cuo.URI = "/test"
	cuo.ClientID = 10
	mc.RedirectURIs = &[]m.ClientRedirectURI{cuo}
	om.MockClient = mc

	om.MockClientList = []m.Client{mc}

	var gt m.ClientGrantType
	gt.ID = 2
	gt.GrantType = "code"
	gt.ClientID = 22

	om.MockClientGrantTypeList = []m.ClientGrantType{gt}

	var au m.ClientAllowedURI
	au.ID = 5
	au.URI = "/testurl"
	au.ClientID = 4

	om.MockClientAllowedURI = au
	om.MockClientAllowedURIList = []m.ClientAllowedURI{au}

	var ru m.ClientRedirectURI
	ru.ID = 4
	ru.URI = "/testuri"
	ru.ClientID = 554

	om.MockClientRedirectURIList = []m.ClientRedirectURI{ru}

	var crl m.ClientRole
	crl.ID = 44
	crl.Role = "hotshot"
	crl.ClientID = 25

	om.MockClientRoleList = []m.ClientRole{crl}

	var rul m.ClientRoleURI
	rul.ClientRoleID = 1
	rul.ClientAllowedURIID = 2

	var rul2 m.ClientRoleURI
	rul2.ClientRoleID = 11
	rul2.ClientAllowedURIID = 21

	om.MockClientRoleURIList = []m.ClientRoleURI{rul, rul2}

	om.MockValidateAccessTokenSuccess = true

	var rh OauthRestHandler
	rh.Manager = &om

	var clt oa.MockOauthClient
	clt.MockValid = true
	rh.Client = &clt

	var ac rc.MockOauthAssets
	ac.MockSuccess = true
	rh.AssetControl = &ac
	return &rh
}
