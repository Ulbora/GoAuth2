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
	deleteAllRedirectURI = "DELETE FROM client_redirect_uri WHERE client_id = ? "
)
