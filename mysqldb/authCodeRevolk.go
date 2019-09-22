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

//AddAuthCodeRevolk AddAuthCodeRevolk
func (d *MySQLOauthDB) AddAuthCodeRevolk(tx dbtx.Transaction, rv *odb.AuthCodeRevolk) (bool, int64) {
	var suc bool
	var id int64
	if tx == nil && !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, rv.AuthorizationCode)
	if tx == nil {
		suc, id = d.DB.Insert(insertAuthCodeRevolk, a...)
	} else {
		suc, id = tx.Insert(insertAuthCodeRevolk, a...)
	}
	return suc, id
}

//GetAuthCodeRevolk GetAuthCodeRevolk
func (d *MySQLOauthDB) GetAuthCodeRevolk(ac int64) *odb.AuthCodeRevolk {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ac)
	row := d.DB.Get(getAuthCodeRevolk, a...)
	rtn := parseAuthCodeRevolkRow(&row.Row)
	return rtn
}

//DeleteAuthCodeRevolk DeleteAuthCodeRevolk
func (d *MySQLOauthDB) DeleteAuthCodeRevolk(tx dbtx.Transaction, ac int64) bool {
	var suc bool
	if tx == nil && !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ac)
	if tx == nil {
		suc = d.DB.Delete(deleteAuthCodeRevolk, a...)
	} else {
		suc = tx.Delete(deleteAuthCodeRevolk, a...)
	}
	return suc
}

func parseAuthCodeRevolkRow(foundRow *[]string) *odb.AuthCodeRevolk {
	var rtn odb.AuthCodeRevolk
	id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		ac, err := strconv.ParseInt((*foundRow)[1], 10, 64)
		if err == nil {
			rtn.ID = id
			rtn.AuthorizationCode = ac
		}
	}
	return &rtn
}
