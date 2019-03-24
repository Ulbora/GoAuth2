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
	AddClientRedirectURI(ru *ClientRedirectURI) (bool, int64)
	GetClientRedirectURIList(clientID int64) *[]ClientRedirectURI
	GetClientRedirectURI(clientID int64, uri string) *ClientRedirectURI
	DeleteClientRedirectURI(id int64) bool

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
}

//Client Client
type Client struct {
	ClientID int64  `json:"clientId"`
	Secret   string `json:"secret"`
	Name     string `json:"name"`
	WebSite  string `json:"webSite"`
	Email    string `json:"email"`
	Enabled  bool   `json:"enabled"`
	Paid     bool   `json:"paid"`
}

//ClientRedirectURI ClientRedirectURI
type ClientRedirectURI struct {
	ID       int64
	URI      string
	ClientID int64
}

//ClientAllowedURI ClientAllowedURI
type ClientAllowedURI struct {
	ID       int64
	URI      string
	ClientID int64
}

//ClientRole ClientRole
type ClientRole struct {
	ID       int64
	Role     string
	ClientID int64
}

//ClientScope ClientScope
type ClientScope struct {
	ID       int64
	Scope    string
	ClientID int64
}

//ClientRoleURI ClientRoleURI
type ClientRoleURI struct {
	ClientRoleID       int64
	ClientAllowedURIID int64
}

//RoleURI RoleURI
type RoleURI struct {
	ClientRoleID       int64
	Role               string
	ClientAllowedURIID int64
	ClientAllowedURI   string
	ClientID           int64
}
