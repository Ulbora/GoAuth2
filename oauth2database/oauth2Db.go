package oauth2database

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

import (
	dbtx "github.com/Ulbora/dbinterface"
)

const (
	//TimeFormat TimeFormat
	TimeFormat = "2006-01-02 15:04:05"
)

//Oauth2DB Oauth2DB
type Oauth2DB interface {
	//client
	AddClient(client *Client, uris *[]ClientRedirectURI) (bool, int64)
	UpdateClient(client *Client) bool
	GetClient(clientID int64) *Client
	GetClients() *[]Client
	SearchClients(name string) *[]Client
	DeleteClient(clientID int64) bool

	//Redirect URI
	AddClientRedirectURI(tx dbtx.Transaction, ru *ClientRedirectURI) (bool, int64)
	GetClientRedirectURIList(clientID int64) *[]ClientRedirectURI
	GetClientRedirectURI(clientID int64, uri string) *ClientRedirectURI
	DeleteClientRedirectURI(tx dbtx.Transaction, id int64) bool
	DeleteClientAllRedirectURI(tx dbtx.Transaction, clientID int64) bool

	//Allowed URI
	AddClientAllowedURI(au *ClientAllowedURI) (bool, int64)
	UpdateClientAllowedURI(au *ClientAllowedURI) bool
	GetClientAllowedURIByID(id int64) *ClientAllowedURI
	GetClientAllowedURIList(clientID int64) *[]ClientAllowedURI
	GetClientAllowedURI(clientID int64, uri string) *ClientAllowedURI
	DeleteClientAllowedURI(id int64) bool

	//Roles
	AddClientRole(r *ClientRole) (bool, int64)
	GetClientRoleList(clientID int64) *[]ClientRole
	DeleteClientRole(id int64) bool

	//Scope
	AddClientScope(s *ClientScope) (bool, int64)
	GetClientScopeList(clientID int64) *[]ClientScope
	DeleteClientScope(id int64) bool

	//Role URI
	AddClientRoleURI(r *ClientRoleURI) bool
	GetClientRoleAllowedURIList(roleID int64) *[]ClientRoleURI
	GetClientRoleAllowedURIListByClientID(clientID int64) *[]RoleURI
	DeleteClientRoleURI(r *ClientRoleURI) bool

	//Refresh Token
	AddRefreshToken(tx dbtx.Transaction, t *RefreshToken) (bool, int64)
	UpdateRefreshToken(t *RefreshToken) bool
	GetRefreshToken(id int64) *RefreshToken
	DeleteRefreshToken(tx dbtx.Transaction, id int64) bool

	//Access Token
	AddAccessToken(tx dbtx.Transaction, t *AccessToken) (bool, int64)
	UpdateAccessToken(tx dbtx.Transaction, t *AccessToken) bool
	GetAccessToken(id int64) *AccessToken
	DeleteAccessToken(tx dbtx.Transaction, id int64) bool

	//start on auth code

	//AuthorizationCode
	AddAuthorizationCode(code *AuthorizationCode, at *AccessToken, rt *RefreshToken, scopeList *[]string) (bool, int64)
	UpdateAuthorizationCode(code *AuthorizationCode) bool
	UpdateAuthorizationCodeAndToken(code *AuthorizationCode, at *AccessToken) bool
	GetAuthorizationCode(clientID int64, userID string) *[]AuthorizationCode
	GetAuthorizationCodeByCode(code string) *AuthorizationCode
	GetAuthorizationCodeByScope(clientID int64, userID string, scope string) *[]AuthorizationCode
	DeleteAuthorizationCode(clientID int64, userID string) bool

	//authcode revolk
	AddAuthCodeRevolk(tx dbtx.Transaction, rv *AuthCodeRevolk) (bool, int64)
	GetAuthCodeRevolk(id int64) *AuthCodeRevolk
	DeleteAuthCodeRevolk(tx dbtx.Transaction, ac int64) bool

	//Auth code scope
	GetAuthorizationCodeScopeList(ac int64) *[]AuthCodeScope

	//grant types
	AddClientGrantType(gt *ClientGrantType) (bool, int64)
	GetClientGrantTypeList(cid int64) *[]ClientGrantType
	DeleteClientGrantType(id int64) bool

	//start here

	//implicit grant
	AddImplicitGrant(ig *ImplicitGrant, at *AccessToken, scopeList *[]string) (bool, int64)
	GetImplicitGrant(clientID int64, userID string) *[]ImplicitGrant
	GetImplicitGrantByScope(clientID int64, userID string, scope string) *[]ImplicitGrant
	DeleteImplicitGrant(clientID int64, userID string) bool

	//implicit grant scope
	GetImplicitGrantScopeList(ig int64) *[]ImplicitScope
}
