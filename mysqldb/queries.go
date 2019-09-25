package mysqldb

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
	oauthTest = "select count(*) from client "

	//Client queries
	insertClient = "insert into client (secret, name, web_site, email, enabled, paid) values(?, ?, ?, ?, ?, ?)"
	updateClient = " UPDATE client SET secret = ?, name = ?, web_site = ?, email = ?, " +
		" enabled = ?, paid = ? WHERE client_id = ? "
	getClientByID      = "SELECT client_id, secret, name, web_site, email, enabled, paid FROM client WHERE client_id = ?"
	getClientsAll      = "SELECT client_id, secret, name, web_site, email, enabled, paid FROM client "
	searchClientByName = "SELECT c.client_id, c.secret, c.name, c.web_site, c.email, c.enabled, c.paid " +
		"FROM client c where c.name like ? "
	deleteClient = "DELETE FROM client WHERE client_id = ? "

	//Redirect URI queries
	insertRedirectURI    = "INSERT INTO client_redirect_uri (uri, client_id) values(?, ?)"
	getRedirectURIList   = "SELECT id, uri, client_id FROM client_redirect_uri WHERE client_id = ? "
	getRedirectURI       = "SELECT id, uri, client_id FROM client_redirect_uri WHERE client_id = ? and uri = ? "
	deleteAllRedirectURI = "DELETE FROM client_redirect_uri WHERE client_id = ? "
	deleteRedirectURI    = "DELETE FROM client_redirect_uri WHERE id = ?"

	//Allowed URI queries
	insertAllowedURI  = "INSERT INTO client_allowed_uri (uri, client_id) values(?, ?) "
	updateAllowedURI  = "UPDATE client_allowed_uri SET uri = ? WHERE id = ? "
	getAllowedURIByID = "SELECT id, uri, client_id from client_allowed_uri WHERE id = ? "
	getAllowedURIList = "SELECT id, uri, client_id from client_allowed_uri WHERE client_id = ? order by uri "
	getAllowedURI     = "SELECT id, uri, client_id from client_allowed_uri WHERE client_id = ? and uri = ? "
	deleteAllowedURI  = "DELETE FROM client_allowed_uri WHERE id = ? "

	//Role
	insertRole  = "INSERT INTO client_role (role, client_id) values(?, ?) "
	getRoleList = "SELECT id, role, client_id FROM client_role WHERE client_id = ? "
	deleteRole  = "DELETE FROM client_role WHERE id = ? "

	//Scope
	insertScope  = "INSERT INTO client_scope (scope, client_id) values(?, ?) "
	getScopeList = "SELECT id, scope, client_id FROM client_scope WHERE client_id = ? "
	deleteScope  = "DELETE FROM client_scope WHERE id = ? "

	//RoleURI
	insertRoleURI  = "INSERT INTO uri_role (client_role_id, client_allowed_uri_id) values(?, ?) "
	getRoleURIList = "SELECT client_role_id, client_allowed_uri_id FROM uri_role WHERE client_role_id = ?"
	roleURIJoin    = "SELECT cr.id as role_id, cr.role, " +
		"cau.id as uri_id, cau.uri, cr.client_id " +
		"FROM client_role cr inner join " +
		"uri_role ur on cr.id = ur.client_role_id " +
		"left join client_allowed_uri cau on cau.id = ur.client_allowed_uri_id " +
		"WHERE cr.client_id = ? " +
		"order by ur.client_role_id "
	deleteRoleURI = "DELETE FROM uri_role WHERE client_role_id = ? and client_allowed_uri_id = ? "

	//Refresh Token
	insertRefreshToken = "INSERT INTO refresh_token (token) values(?)"
	updateRefreshToken = "UPDATE refresh_token SET token = ? WHERE id = ? "
	getRefreshToken    = "SELECT id, token FROM refresh_token WHERE id = ? "
	deleteRefreshToken = "DELETE FROM refresh_token WHERE id = ? "

	//Access Token
	insertAccessToken     = "INSERT INTO access_token  (token, expires, refresh_token_id) values(?, ?, ?) "
	insertAccessTokenNull = "INSERT INTO access_token  (token, expires) values(?, ?) "
	updateAccessToken     = "UPDATE access_token SET token = ?, expires = ?, refresh_token_id = ? WHERE id = ? "
	updateAccessTokenNull = "UPDATE access_token SET token = ?, expires = ? WHERE id = ? "
	getAccessToken        = "SELECT id, token, expires, refresh_token_id FROM access_token WHERE id = ? "
	deleteAccessToken     = "DELETE FROM access_token WHERE id = ? "

	//Auth Code
	insertAuthCode = "INSERT INTO authorization_code  (client_id, user_id, expires, access_token_id, randon_auth_code, already_used) values(?, ?, ?, ?, ?, ?) "
	updateAuthCode = "UPDATE authorization_code SET randon_auth_code = ?, already_used = ? " +
		"WHERE authorization_code = ? "
	updateAuthCodeToken = "UPDATE authorization_code SET expires = ? " +
		"WHERE authorization_code = ? "

	getByAuthorizationCodeClientUser = "SELECT authorization_code, client_id, user_id, expires,  access_token_id, randon_auth_code, already_used " +
		"FROM authorization_code WHERE client_id = ? and user_id = ?"
	getAuthorizationCodeByCode = "SELECT authorization_code, client_id, user_id, expires,  access_token_id, randon_auth_code, already_used " +
		"FROM authorization_code WHERE randon_auth_code = ?"
	getAuthorizationCodeByClientUserScope = "SELECT a.authorization_code, a.client_id, s.scope, a.randon_auth_code, a.already_used " +
		"FROM authorization_code a inner join auth_code_scope s " +
		"on a.authorization_code = s.authorization_code " +
		"WHERE a.client_id = ? and a.user_id = ? and s.scope = ?"
	deleteAuthCode = "DELETE FROM authorization_code WHERE client_id = ? and user_id = ?"

	//Auth Code Scope
	insertAuthCodeScope           = "INSERT INTO auth_code_scope  (scope, authorization_code) values(?, ?) "
	getAuthorizationCodeScopeList = "SELECT id, scope, authorization_code " +
		"FROM auth_code_scope WHERE authorization_code = ?"
	deleteAllAuthCodeScope = "DELETE FROM auth_code_scope WHERE authorization_code = ?"

	//Auth Code Revolk
	insertAuthCodeRevolk = "INSERT INTO auth_code_revoke  (authorization_code) values(?) "
	getAuthCodeRevolk    = "SELECT id, authorization_code FROM auth_code_revoke WHERE authorization_code = ?"
	deleteAuthCodeRevolk = "DELETE FROM auth_code_revoke WHERE authorization_code = ?"

	//Grant Types
	insertClientGrantType  = "INSERT INTO client_grant_type  (grant_type, client_id) values(?, ?) "
	getClientGrantTypeList = "SELECT * FROM client_grant_type WHERE client_id = ?"
	deleteClientGrantType  = "DELETE FROM client_grant_type WHERE id = ?"
)
