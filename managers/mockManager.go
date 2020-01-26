package managers

import au "github.com/Ulbora/auth_interface"

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

//MockManager MockManager
type MockManager struct {
	MockInsertSuccess1 bool
	MockInsertID1      int64

	MockUpdateSuccess1 bool
	MockDeleteSuccess1 bool

	MockClient     Client
	MockClientList []Client

	MockClientRedirectURIList []ClientRedirectURI
	MockClientRoleList        []ClientRole

	MockClientAllowedURI     ClientAllowedURI
	MockClientAllowedURIList []ClientAllowedURI

	MockClientRoleURIList []ClientRoleURI

	MockClientGrantTypeList []ClientGrantType

	MockAuthCodeAuthorizeSuccess bool
	MockAuthCode                 int64
	MockAuthCodeString           string

	MockAuthCodeAuthorized bool

	MockAuthCodeClient AuthCodeClient

	MockImplicitAuthorizeSuccess bool
	MockImplicitReturn           ImplicitReturn
	MockImplicitAuthorized       bool
	MockImplicitClient           ImplicitClient

	MockToken                       Token
	MockAuthCodeTokenSuccess        bool
	MockCredentialsTokenSuccess     bool
	MockPasswordTokenSuccess        bool
	MockAuthCodeRefreshTokenSuccess bool
	MockPasswordRefreshTokenSuccess bool

	MockTokenError string

	MockValidateAccessTokenSuccess bool

	MockUserLoginSuccess bool
}

//client

//AddClient AddClient
func (m *MockManager) AddClient(client *Client) (bool, int64) {
	return m.MockInsertSuccess1, m.MockInsertID1
}

//UpdateClient UpdateClient
func (m *MockManager) UpdateClient(client *Client) bool {
	return m.MockUpdateSuccess1
}

//GetClient GetClient
func (m *MockManager) GetClient(id int64) *Client {
	return &m.MockClient
}

//GetClientList GetClientList
func (m *MockManager) GetClientList() *[]Client {
	return &m.MockClientList
}

//GetClientSearchList GetClientSearchList
func (m *MockManager) GetClientSearchList(name string) *[]Client {
	return &m.MockClientList
}

//DeleteClient DeleteClient
func (m *MockManager) DeleteClient(id int64) bool {
	return m.MockDeleteSuccess1
}

//client redirect uri

//AddClientRedirectURI AddClientRedirectURI
func (m *MockManager) AddClientRedirectURI(ru *ClientRedirectURI) (bool, int64) {
	return m.MockInsertSuccess1, m.MockInsertID1
}

//GetClientRedirectURIList GetClientRedirectURIList
func (m *MockManager) GetClientRedirectURIList(clientID int64) *[]ClientRedirectURI {
	return &m.MockClientRedirectURIList
}

//DeleteClientRedirectURI DeleteClientRedirectURI
func (m *MockManager) DeleteClientRedirectURI(id int64) bool {
	return m.MockDeleteSuccess1
}

//client roles

//AddClientRole AddClientRole
func (m *MockManager) AddClientRole(r *ClientRole) (bool, int64) {
	return m.MockInsertSuccess1, m.MockInsertID1
}

//GetClientRoleList GetClientRoleList
func (m *MockManager) GetClientRoleList(clientID int64) *[]ClientRole {
	return &m.MockClientRoleList
}

//DeleteClientRole DeleteClientRole
func (m *MockManager) DeleteClientRole(id int64) bool {
	return m.MockDeleteSuccess1
}

//client allowed uri

//AddClientAllowedURI AddClientAllowedURI
func (m *MockManager) AddClientAllowedURI(au *ClientAllowedURI) (bool, int64) {
	return m.MockInsertSuccess1, m.MockInsertID1
}

//UpdateClientAllowedURI UpdateClientAllowedURI
func (m *MockManager) UpdateClientAllowedURI(au *ClientAllowedURI) bool {
	return m.MockUpdateSuccess1
}

//GetClientAllowedURI GetClientAllowedURI
func (m *MockManager) GetClientAllowedURI(id int64) *ClientAllowedURI {
	return &m.MockClientAllowedURI
}

//GetClientAllowedURIList GetClientAllowedURIList
func (m *MockManager) GetClientAllowedURIList(clientID int64) *[]ClientAllowedURI {
	return &m.MockClientAllowedURIList
}

//DeleteClientAllowedURI DeleteClientAllowedURI
func (m *MockManager) DeleteClientAllowedURI(id int64) bool {
	return m.MockDeleteSuccess1
}

//client role uri

//AddClientRoleURI AddClientRoleURI
func (m *MockManager) AddClientRoleURI(r *ClientRoleURI) bool {
	return m.MockInsertSuccess1
}

//GetClientRoleAllowedURIList GetClientRoleAllowedURIList
func (m *MockManager) GetClientRoleAllowedURIList(roleID int64) *[]ClientRoleURI {
	return &m.MockClientRoleURIList
}

//DeleteClientRoleURI DeleteClientRoleURI
func (m *MockManager) DeleteClientRoleURI(r *ClientRoleURI) bool {
	return m.MockDeleteSuccess1
}

//client grant type

//AddClientGrantType AddClientGrantType
func (m *MockManager) AddClientGrantType(gt *ClientGrantType) (bool, int64) {
	return m.MockInsertSuccess1, m.MockInsertID1
}

//GetClientGrantTypeList GetClientGrantTypeList
func (m *MockManager) GetClientGrantTypeList(clientID int64) *[]ClientGrantType {
	return &m.MockClientGrantTypeList
}

//DeleteClientGrantType DeleteClientGrantType
func (m *MockManager) DeleteClientGrantType(id int64) bool {
	return m.MockDeleteSuccess1
}

//auth code

//AuthorizeAuthCode AuthorizeAuthCode
func (m *MockManager) AuthorizeAuthCode(ac *AuthCode) (success bool, authCode int64, authCodeString string) {
	return m.MockAuthCodeAuthorizeSuccess, m.MockAuthCode, m.MockAuthCodeString
}

//CheckAuthCodeApplicationAuthorization CheckAuthCodeApplicationAuthorization
func (m *MockManager) CheckAuthCodeApplicationAuthorization(ac *AuthCode) (authorized bool) {
	return m.MockAuthCodeAuthorized
}

//ValidateAuthCodeClientAndCallback ValidateAuthCodeClientAndCallback
func (m *MockManager) ValidateAuthCodeClientAndCallback(ac *AuthCode) *AuthCodeClient {
	return &m.MockAuthCodeClient
}

//implicit

//AuthorizeImplicit AuthorizeImplicit
func (m *MockManager) AuthorizeImplicit(imp *Implicit) (bool, *ImplicitReturn) {
	return m.MockImplicitAuthorizeSuccess, &m.MockImplicitReturn
}

//CheckImplicitApplicationAuthorization CheckImplicitApplicationAuthorization
func (m *MockManager) CheckImplicitApplicationAuthorization(imp *Implicit) (authorized bool) {
	return m.MockImplicitAuthorized
}

//ValidateImplicitClientAndCallback ValidateImplicitClientAndCallback
func (m *MockManager) ValidateImplicitClientAndCallback(imp *Implicit) *ImplicitClient {
	return &m.MockImplicitClient
}

//token manager

//GetAuthCodeToken GetAuthCodeToken
func (m *MockManager) GetAuthCodeToken(act *AuthCodeTokenReq) (bool, *Token, string) {
	return m.MockAuthCodeTokenSuccess, &m.MockToken, m.MockTokenError
}

//GetCredentialsToken GetCredentialsToken
func (m *MockManager) GetCredentialsToken(ct *CredentialsTokenReq) (bool, *Token, string) {
	return m.MockCredentialsTokenSuccess, &m.MockToken, m.MockTokenError
}

//GetPasswordToken GetPasswordToken
func (m *MockManager) GetPasswordToken(pt *PasswordTokenReq) (bool, *Token, string) {
	return m.MockPasswordTokenSuccess, &m.MockToken, m.MockTokenError
}

//GetAuthCodeAccesssTokenWithRefreshToken GetAuthCodeAccesssTokenWithRefreshToken
func (m *MockManager) GetAuthCodeAccesssTokenWithRefreshToken(rt *RefreshTokenReq) (bool, *Token, string) {
	return m.MockAuthCodeRefreshTokenSuccess, &m.MockToken, m.MockTokenError
}

//GetPasswordAccesssTokenWithRefreshToken GetPasswordAccesssTokenWithRefreshToken
func (m *MockManager) GetPasswordAccesssTokenWithRefreshToken(rt *RefreshTokenReq) (bool, *Token, string) {
	return m.MockPasswordRefreshTokenSuccess, &m.MockToken, m.MockTokenError
}

//validate Token

//ValidateAccessToken ValidateAccessToken
func (m *MockManager) ValidateAccessToken(at *ValidateAccessTokenReq) bool {
	return m.MockValidateAccessTokenSuccess
}

//UserLogin UserLogin
func (m *MockManager) UserLogin(login *au.Login) bool {
	return m.MockUserLoginSuccess
}
