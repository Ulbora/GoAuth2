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
	dbtx "github.com/Ulbora/dbinterface"
)

//AddAuthCodeScope AddAuthCodeScope
func (d *MySQLOauthDB) AddAuthCodeScope(tx dbtx.Transaction, as *odb.AuthCodeScope) (bool, int64) {
	var suc bool
	var id int64
	if tx != nil {
		var a []interface{}
		a = append(a, as.Scope, as.AuthorizationCode)
		suc, id = tx.Insert(insertAuthCodeScope, a...)
	}
	return suc, id
}

//DeleteAuthCodeScopeList DeleteAuthCodeScopeList
func (d *MySQLOauthDB) DeleteAuthCodeScopeList(tx dbtx.Transaction, ac int64) bool {
	var suc bool
	if tx != nil {
		var a []interface{}
		a = append(a, ac)
		suc = tx.Delete(deleteAllAuthCodeScope, a...)
	}
	return suc
}
