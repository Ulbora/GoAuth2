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
	AddRefreshToken(t *RefreshToken) (bool, int64)
	UpdateRefreshToken(t *RefreshToken) bool
	GetRefreshToken(id int64) *RefreshToken
	DeleteRefreshToken(id int64) bool

	AddAccessToken(t *AccessToken) (bool, int64)
	UpdateAccessToken(t *AccessToken) bool
	GetAccessToken(id int64) *AccessToken
	DeleteAccessToken(id int64) bool
}
