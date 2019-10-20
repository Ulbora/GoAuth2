package managers

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
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//Manager Manager
type Manager interface {
	//client
	AddClient(client *Client) (bool, int64)
	UpdateClient(client *Client) bool
	GetClient(id int64) *Client
	GetClientList() *[]Client
	GetClientSearchList(name string) *[]Client
	DeleteClient(id int64) bool

	//client redirect uri
	AddClientRedirectURI(ru *ClientRedirectURI) (bool, int64)
	GetClientRedirectURIList(clientID int64) *[]ClientRedirectURI
	DeleteClientRedirectURI(id int64) bool

	// //client roles
	AddClientRole(r *ClientRole) (bool, int64)
	GetClientRoleList(clientID int64) *[]ClientRole
	DeleteClientRole(id int64) bool

	// //client allowed uri
	// AddClientAllowedUri(au *odb.ClientAllowedURI) (bool, int64)
	// UpdateClientAllowedUri(au *odb.ClientAllowedURI) bool
	// GetClientAllowedUri(id int64) *odb.ClientAllowedURI
	// GetClientAllowedUriList(clientID int64) *[]odb.ClientAllowedURI
	// DeleteClientAllowedUri(id int64) bool

	// //client role uri
	// AddClientRoleUri(r *odb.ClientRoleURI) (bool, int64)
	// GetClientRoleAllowedUriList(roleID int64) *[]odb.ClientRoleURI
	// DeleteClientRoleUri(r *odb.ClientRoleURI) bool

	// //client grant type
	// AddClientGrantType(gt *odb.ClientGrantType) (bool, int64)
	// GetClientGrantTypeList(clientID int64) *[]odb.ClientGrantType
	// DeleteClientGrantType(id int64) bool

	// //auth code
	// AuthorizeAuthCode(ac *AuthCode) (success bool, authCode int64, authCodeString string)
	// CheckAuthCodeApplicationAuthorization(ac *AuthCode) (authorized bool)
	// ValidateAuthCodeClientAndCallback(ac *AuthCode) *AuthCodeClient

	// //implicit
	// AuthorizeImplicit(imp *Implicit) (bool, *ImplicitReturn)
	// CheckImplicitApplicationAuthorization(imp *Implicit) (authorized bool)
	// ValidateImplicitClientAndCallback(imp *Implicit) *ImplicitClient

	// //token manager
	// GetAuthCodeToken(act *AuthCodeTokenReq) *Token
	// GetCredentialsToken(ct *CredentialsTokenReq) *Token
	// GetRefreshToken(rt *RefreshTokenReq) *Token
	// GetPasswordToken(pt *PasswordTokenReq) *Token

	// //validate Token
	// ValidateAccessToken(at *ValidateAccessTokenReq) bool
}

//OauthManager OauthManager
type OauthManager struct {
	Db odb.Oauth2DB
}
