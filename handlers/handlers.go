//Package handlers ...
package handlers

import "net/http"

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

const (
	codeRespType     = "code"
	tokenRespType    = "token"
	implicitRespType = "implicit"
	clientRespType   = "client_credentials"
	passwordRespType = "password"

	authorizationCodeGrantType = "authorization_code"
	passwordGrantType          = "password"
	credentialGrantType        = "client_credentials"
	refreshTokenGrantType      = "refresh_token"

	invalidReqestError   = "Invalid Request"
	invalidRedirectError = "Invalid redirect URI"
	accessDenidError     = "access_denied"

	authAppPageTitle = "Authorize Application"
	loginPageTitle   = "GoAuth2 Login Page"
)

//ResponseID ResponseID
type ResponseID struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
}

//Response Response
type Response struct {
	Success bool `json:"success"`
}

//RestHandler RestHandler
type RestHandler interface {
	AddAllowedURISuper(w http.ResponseWriter, r *http.Request)
	AddAllowedURI(w http.ResponseWriter, r *http.Request)
	UpdateAllowedURISuper(w http.ResponseWriter, r *http.Request)
	UpdateAllowedURI(w http.ResponseWriter, r *http.Request)
	GetAllowedURI(w http.ResponseWriter, r *http.Request)
	GetAllowedURIList(w http.ResponseWriter, r *http.Request)
	DeleteAllowedURI(w http.ResponseWriter, r *http.Request)

	AddGrantType(w http.ResponseWriter, r *http.Request)
	GetGrantTypeList(w http.ResponseWriter, r *http.Request)
	DeleteGrantType(w http.ResponseWriter, r *http.Request)

	AddRedirectURI(w http.ResponseWriter, r *http.Request)
	GetRedirectURIList(w http.ResponseWriter, r *http.Request)
	DeleteRedirectURI(w http.ResponseWriter, r *http.Request)

	AddRoleSuper(w http.ResponseWriter, r *http.Request)
	AddRole(w http.ResponseWriter, r *http.Request)
	GetRoleList(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)

	AddRoleURI(w http.ResponseWriter, r *http.Request)
	GetRoleURIList(w http.ResponseWriter, r *http.Request)
	DeleteRoleURI(w http.ResponseWriter, r *http.Request)

	AddClient(w http.ResponseWriter, r *http.Request)
	UpdateClient(w http.ResponseWriter, r *http.Request)
	GetClient(w http.ResponseWriter, r *http.Request)
	GetClientAdmin(w http.ResponseWriter, r *http.Request)
	GetClientList(w http.ResponseWriter, r *http.Request)
	GetClientSearchList(w http.ResponseWriter, r *http.Request)
	DeleteClient(w http.ResponseWriter, r *http.Request)

	ValidateAccessToken(w http.ResponseWriter, r *http.Request)
}

//WebHandler WebHandler
type WebHandler interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	AuthorizeApp(w http.ResponseWriter, r *http.Request)
	ApplicationAuthorizationByUser(w http.ResponseWriter, r *http.Request)
	OauthError(w http.ResponseWriter, r *http.Request)

	Login(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)

	Token(w http.ResponseWriter, r *http.Request)
	// RefreshToken(w http.ResponseWriter, r *http.Request)
}
