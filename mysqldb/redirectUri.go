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

//AddClientRedirectURI AddClientRedirectURI
func (d *MySQLOauthDB) AddClientRedirectURI(ru *odb.ClientRedirectURI) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ru.URI, ru.ClientID)
	suc, id := d.DB.Insert(insertRedirectURI, a...)
	return suc, id
}

//GetClientRedirectURIList GetClientRedirectURIList
func (d *MySQLOauthDB) GetClientRedirectURIList(clientID int64) *[]odb.ClientRedirectURI {
	var rtn []odb.ClientRedirectURI
	var a []interface{}
	a = append(a, clientID)
	rows := d.DB.GetList(getRedirectURIList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientURIRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetClientRedirectURI GetClientRedirectURI
func (d *MySQLOauthDB) GetClientRedirectURI(clientID int64, uri string) *odb.ClientRedirectURI {
	var a []interface{}
	a = append(a, clientID, uri)
	row := d.DB.Get(getRedirectURI, a...)
	rtn := parseClientURIRow(&row.Row)
	return rtn
}

//DeleteClientRedirectURI DeleteClientRedirectURI
func (d *MySQLOauthDB) DeleteClientRedirectURI(id int64) bool {
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteRedirectURI, a...)
}

func parseClientURIRow(foundRow *[]string) *odb.ClientRedirectURI {
	var rtn odb.ClientRedirectURI
	id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		clientID, err := strconv.ParseInt((*foundRow)[2], 10, 64)
		if err == nil {
			rtn.ID = id
			rtn.ClientID = clientID
			rtn.URI = (*foundRow)[1]
		}
	}
	return &rtn
}
