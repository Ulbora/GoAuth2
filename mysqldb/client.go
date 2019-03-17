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
import (
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//AddClient AddClient
func (d *MySQLDB) AddClient(client *odb.Client, uris *[]odb.ClientRedirectURI) (bool, int64) {
	var suc = false
	var id int64
	return suc, id
}

//UpdateClient UpdateClient
func (d *MySQLDB) UpdateClient(client *odb.Client) bool {
	var suc = false

	return suc
}

//GetClient GetClient
func (d *MySQLDB) GetClient(clientID int64) *odb.Client {
	var rtn odb.Client

	return &rtn
}

//GetClients GetClients
func (d *MySQLDB) GetClients() *[]odb.Client {
	var rtn []odb.Client

	return &rtn
}

//SearchClients SearchClients
func (d *MySQLDB) SearchClients(name string) *[]odb.Client {
	var rtn []odb.Client

	return &rtn
}

//DeleteClient DeleteClient
func (d *MySQLDB) DeleteClient(clientID int64) bool {
	var suc = false

	return suc
}
