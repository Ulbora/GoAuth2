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

//AddRefreshToken AddRefreshToken
func (d *MySQLOauthDB) AddRefreshToken(tx dbtx.Transaction, t *odb.RefreshToken) (bool, int64) {
	var suc bool
	var id int64
	if tx == nil && !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, t.Token)
	if tx == nil {
		suc, id = d.DB.Insert(insertRefreshToken, a...)
	} else {
		suc, id = tx.Insert(insertRefreshToken, a...)
	}
	return suc, id
}

//UpdateRefreshToken UpdateRefreshToken
func (d *MySQLOauthDB) UpdateRefreshToken(t *odb.RefreshToken) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, t.Token, t.ID)
	suc := d.DB.Update(updateRefreshToken, a...)
	return suc
}

//GetRefreshToken GetRefreshToken
func (d *MySQLOauthDB) GetRefreshToken(id int64) *odb.RefreshToken {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getRefreshToken, a...)
	rtn := parseRefreshTokenRow(&row.Row)
	return rtn
}

//DeleteRefreshToken DeleteRefreshToken
func (d *MySQLOauthDB) DeleteRefreshToken(tx dbtx.Transaction, id int64) bool {
	var suc bool
	if tx == nil && !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	if tx == nil {
		suc = d.DB.Delete(deleteRefreshToken, a...)
	} else {
		suc = tx.Delete(deleteRefreshToken, a...)
	}
	return suc
}

func parseRefreshTokenRow(foundRow *[]string) *odb.RefreshToken {
	var rtn odb.RefreshToken
	id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		rtn.ID = id
		rtn.Token = (*foundRow)[1]
	}
	return &rtn
}
