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
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	dbtx "github.com/Ulbora/dbinterface"
)

//AddAccessToken AddAccessToken
func (d *MySQLOauthDB) AddAccessToken(tx dbtx.Transaction, t *odb.AccessToken) (bool, int64) {
	if tx == nil && !d.testConnection() {
		d.DB.Connect()
	}
	var suc bool
	var id int64
	var a []interface{}
	if tx == nil && t.RefreshTokenID == 0 {
		a = append(a, t.Token, t.Expires)
		suc, id = d.DB.Insert(insertAccessTokenNull, a...)
	} else if tx == nil {
		a = append(a, t.Token, t.Expires, t.RefreshTokenID)
		suc, id = d.DB.Insert(insertAccessToken, a...)
	} else if t.RefreshTokenID == 0 {
		a = append(a, t.Token, t.Expires)
		suc, id = tx.Insert(insertAccessTokenNull, a...)
	} else {
		a = append(a, t.Token, t.Expires, t.RefreshTokenID)
		suc, id = tx.Insert(insertAccessToken, a...)
	}

	return suc, id
}

//UpdateAccessToken UpdateAccessToken
func (d *MySQLOauthDB) UpdateAccessToken(tx dbtx.Transaction, t *odb.AccessToken) bool {
	if tx == nil && !d.testConnection() {
		d.DB.Connect()
	}
	var suc bool
	var a []interface{}
	if tx == nil && t.RefreshTokenID == 0 {
		a = append(a, t.Token, t.Expires, t.ID)
		suc = d.DB.Update(updateAccessTokenNull, a...)
	} else if tx == nil {
		a = append(a, t.Token, t.Expires, t.RefreshTokenID, t.ID)
		suc = d.DB.Update(updateAccessToken, a...)
	} else if t.RefreshTokenID == 0 {
		a = append(a, t.Token, t.Expires, t.ID)
		suc = tx.Update(updateAccessTokenNull, a...)
	} else {
		a = append(a, t.Token, t.Expires, t.RefreshTokenID, t.ID)
		suc = tx.Update(updateAccessToken, a...)
	}
	return suc
}

//GetAccessToken GetAccessToken
func (d *MySQLOauthDB) GetAccessToken(id int64) *odb.AccessToken {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getAccessToken, a...)
	rtn := parseAccessTokenRow(&row.Row)
	return rtn
}

//DeleteAccessToken DeleteAccessToken
func (d *MySQLOauthDB) DeleteAccessToken(tx dbtx.Transaction, id int64) bool {
	var suc bool
	if tx == nil && !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	if tx == nil {
		suc = d.DB.Delete(deleteAccessToken, a...)
	} else {
		suc = tx.Delete(deleteAccessToken, a...)
	}
	return suc
}

func parseAccessTokenRow(foundRow *[]string) *odb.AccessToken {
	var rtn odb.AccessToken
	id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		cTime, err := time.Parse(odb.TimeFormat, (*foundRow)[2])
		if err == nil {
			rtn.ID = id
			rtn.Token = (*foundRow)[1]
			rtn.Expires = cTime
			if (*foundRow)[3] != "" {
				refTokID, _ := strconv.ParseInt((*foundRow)[3], 10, 64)
				rtn.RefreshTokenID = refTokID
			}
		}
	}
	return &rtn
}
