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
	"strconv"
)

//AddClientRole AddClientRole
func (d *MySQLOauthDB) AddClientRole(r *odb.ClientRole) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, r.Role, r.ClientID)
	suc, id := d.DB.Insert(insertRole, a...)
	return suc, id
}

//GetClientRoleList GetClientRoleList
func (d *MySQLOauthDB) GetClientRoleList(clientID int64) *[]odb.ClientRole {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.ClientRole
	var a []interface{}
	a = append(a, clientID)
	rows := d.DB.GetList(getRoleList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientRoleRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteClientRole DeleteClientRole
func (d *MySQLOauthDB) DeleteClientRole(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteRole, a...)
}

func parseClientRoleRow(foundRow *[]string) *odb.ClientRole {
	var rtn odb.ClientRole
	id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		clientID, err := strconv.ParseInt((*foundRow)[2], 10, 64)
		if err == nil {
			rtn.ID = id
			rtn.ClientID = clientID
			rtn.Role = (*foundRow)[1]
		}
	}
	return &rtn
}
