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

//AddClientRoleURI AddClientRoleURI
func (d *MySQLOauthDB) AddClientRoleURI(r *odb.ClientRoleURI) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, r.ClientRoleID, r.ClientAllowedURIID)
	suc, _ := d.DB.Insert(insertRoleURI, a...)
	return suc
}

//GetClientRoleAllowedURIList GetClientRoleAllowedURIList
func (d *MySQLOauthDB) GetClientRoleAllowedURIList(roleID int64) *[]odb.ClientRoleURI {
	var rtn []odb.ClientRoleURI
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, roleID)
	rows := d.DB.GetList(getRoleURIList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseRowURIRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

// GetClientRoleAllowedURIListByClientID GetClientRoleAllowedURIListByClientID
func (d *MySQLOauthDB) GetClientRoleAllowedURIListByClientID(clientID int64) *[]odb.RoleURI {
	var rtn []odb.RoleURI
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, clientID)
	rows := d.DB.GetList(roleURIJoin, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientRowURIRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteClientRoleURI DeleteClientRoleURI
func (d *MySQLOauthDB) DeleteClientRoleURI(r *odb.ClientRoleURI) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, r.ClientRoleID, r.ClientAllowedURIID)
	return d.DB.Delete(deleteRoleURI, a...)
}

func parseRowURIRow(foundRow *[]string) *odb.ClientRoleURI {
	var rtn odb.ClientRoleURI
	roleID, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		URIID, err := strconv.ParseInt((*foundRow)[1], 10, 64)
		if err == nil {
			rtn.ClientRoleID = roleID
			rtn.ClientAllowedURIID = URIID

		}
	}
	return &rtn
}

func parseClientRowURIRow(foundRow *[]string) *odb.RoleURI {
	var rtn odb.RoleURI
	roleID, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		URIID, err := strconv.ParseInt((*foundRow)[2], 10, 64)
		if err == nil {
			clientID, err := strconv.ParseInt((*foundRow)[4], 10, 64)
			if err == nil {
				rtn.ClientRoleID = roleID
				rtn.ClientAllowedURIID = URIID
				rtn.ClientID = clientID
				rtn.Role = (*foundRow)[1]
				rtn.ClientAllowedURI = (*foundRow)[3]
			}
		}
	}
	return &rtn
}
