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

//AddClientScope AddClientScope
func (d *MySQLOauthDB) AddClientScope(s *odb.ClientScope) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Scope, s.ClientID)
	suc, id := d.DB.Insert(insertScope, a...)
	return suc, id
}

//GetClientScopeList GetClientScopeList
func (d *MySQLOauthDB) GetClientScopeList(clientID int64) *[]odb.ClientScope {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.ClientScope
	var a []interface{}
	a = append(a, clientID)
	rows := d.DB.GetList(getScopeList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientScopeRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteClientScope DeleteClientScope
func (d *MySQLOauthDB) DeleteClientScope(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteScope, a...)
}

func parseClientScopeRow(foundRow *[]string) *odb.ClientScope {
	var rtn odb.ClientScope
	id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		clientID, err := strconv.ParseInt((*foundRow)[2], 10, 64)
		if err == nil {
			rtn.ID = id
			rtn.ClientID = clientID
			rtn.Scope = (*foundRow)[1]
		}
	}
	return &rtn
}
