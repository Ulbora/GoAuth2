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

//GetAccessTokenKey GetAccessTokenKey
func (d *MySQLOauthDB) GetAccessTokenKey() string {
	var rtn string
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	row := d.DB.Get(getAccessTokenKey, a...)
	if len(row.Row) > 0 {
		rtn = row.Row[1]
	}
	return rtn
}

//GetRefreshTokenKey GetRefreshTokenKey
func (d *MySQLOauthDB) GetRefreshTokenKey() string {
	var rtn string
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	row := d.DB.Get(getRefreshTokenKey, a...)
	if len(row.Row) > 0 {
		rtn = row.Row[1]
	}
	return rtn
}

//GetSessionKey GetSessionKey
func (d *MySQLOauthDB) GetSessionKey() string {
	var rtn string
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	row := d.DB.Get(getSessionKey, a...)
	if len(row.Row) > 0 {
		rtn = row.Row[1]
	}
	return rtn
}
