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
	"strconv"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	dbtx "github.com/Ulbora/dbinterface"
)

//AddImplicitGrantScope AddImplicitGrantScope
func (d *MySQLOauthDB) AddImplicitGrantScope(tx dbtx.Transaction, igs *odb.ImplicitScope) (bool, int64) {
	var suc bool
	var id int64
	if tx != nil {
		var a []interface{}
		a = append(a, igs.Scope, igs.ImplicitGrantID)
		suc, id = tx.Insert(insertImplicitScope, a...)
	}
	return suc, id
}

//GetImplicitGrantScopeList GetImplicitGrantScopeList
func (d *MySQLOauthDB) GetImplicitGrantScopeList(ig int64) *[]odb.ImplicitScope {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.ImplicitScope
	var a []interface{}
	a = append(a, ig)
	rows := d.DB.GetList(getImplicitScopeList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			if len(foundRow) > 0 {
				var igs odb.ImplicitScope
				id, err := strconv.ParseInt((foundRow)[0], 10, 64)
				if err == nil {
					igid, err := strconv.ParseInt((foundRow)[2], 10, 64)
					if err == nil {
						igs.ID = id
						igs.Scope = (foundRow)[1]
						igs.ImplicitGrantID = igid
					}
				}
				rtn = append(rtn, igs)
			}
		}
	}
	return &rtn
}

//DeleteImplicitGrantScopeList DeleteImplicitGrantScopeList
func (d *MySQLOauthDB) DeleteImplicitGrantScopeList(tx dbtx.Transaction, ig int64) bool {
	var suc bool
	if tx != nil {
		var a []interface{}
		a = append(a, ig)
		suc = tx.Delete(deleteImplicitScope, a...)
	}
	return suc
}
