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
)

//AddClientAllowedURI AddClientAllowedURI
func (d *MySQLOauthDB) AddClientAllowedURI(au *odb.ClientAllowedURI) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, au.URI, au.ClientID)
	suc, id := d.DB.Insert(insertAllowedURI, a...)
	return suc, id
}

//UpdateClientAllowedURI UpdateClientAllowedURI
func (d *MySQLOauthDB) UpdateClientAllowedURI(au *odb.ClientAllowedURI) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, au.URI, au.ID)
	suc := d.DB.Update(updateAllowedURI, a...)
	return suc
}

//GetClientAllowedURIByID GetClientAllowedURIByID
func (d *MySQLOauthDB) GetClientAllowedURIByID(id int64) *odb.ClientAllowedURI {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getAllowedURIByID, a...)
	rtn := parseClientAllowedURIRow(&row.Row)
	return rtn
}

//GetClientAllowedURIList GetClientAllowedURIList
func (d *MySQLOauthDB) GetClientAllowedURIList(clientID int64) *[]odb.ClientAllowedURI {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.ClientAllowedURI
	var a []interface{}
	a = append(a, clientID)
	rows := d.DB.GetList(getAllowedURIList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientAllowedURIRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetClientAllowedURI GetClientAllowedURI
func (d *MySQLOauthDB) GetClientAllowedURI(clientID int64, uri string) *odb.ClientAllowedURI {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, clientID, uri)
	row := d.DB.Get(getAllowedURI, a...)
	rtn := parseClientAllowedURIRow(&row.Row)
	return rtn
}

//DeleteClientAllowedURI DeleteClientAllowedURI
func (d *MySQLOauthDB) DeleteClientAllowedURI(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteAllowedURI, a...)
}

func parseClientAllowedURIRow(foundRow *[]string) *odb.ClientAllowedURI {
	var rtn odb.ClientAllowedURI
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		if err == nil {
			clientID, err := strconv.ParseInt((*foundRow)[2], 10, 64)
			if err == nil {
				rtn.ID = id
				rtn.ClientID = clientID
				rtn.URI = (*foundRow)[1]
			}
		}
	}
	return &rtn
}
