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
	//"fmt"

	

	odb "github.com/Ulbora/GoAuth2/oauth2database"

	//"log"
	"strconv"
)

//AddClient AddClient
func (d *MySQLOauthDB) AddClient(client *odb.Client, uris *[]odb.ClientRedirectURI) (bool, int64) {
	var suc = false
	//log.Println("in add client")
	if !d.testConnection() {
		d.DB.Connect()
	}
	var fail bool
	tx := d.DB.BeginTransaction()
	var a []interface{}
	a = append(a, client.Secret, client.Name, client.WebSite, client.Email, client.Enabled, client.Paid)
	succ, cid := tx.Insert(insertClient, a...)
	if succ && cid > 0 {
		//fmt.Println("cid in add client: ", cid)
		//fail = addURIs(tx, id, uris)
		if uris != nil && len(*uris) > 0 {
			for _, u := range *uris {
				//d.Tx = tx
				u.ClientID = cid
				rsus, rid := d.AddClientRedirectURI(tx, &u)
				// fmt.Println("cid in add client: ", cid)
				// fmt.Println("rid in add client: ", rid)
				// fmt.Println("rsus in add client: ", rsus)
				// var au []interface{}
				// au = append(au, u.URI, id)
				// u.ClientID = id
				// rsus, rid := tx.Insert(insertRedirectURI, au...)
				if !rsus || rid <= 0 {
					tx.Rollback()
					fail = true
					break
				}
			}
		}
	} else {
		fail = true
		tx.Rollback()
	}
	if !fail {
		suc = true
		tx.Commit()
	}
	return suc, cid
}

// func addURIs(tx dbtx.Transaction, id int64, uris *[]odb.ClientRedirectURI) bool {
// 	var fail = false
// 	if uris != nil && len(*uris) > 0 {
// 		for _, u := range *uris {
// 			var au []interface{}
// 			au = append(au, u.URI, id)
// 			u.ClientID = id
// 			rsus, rid := tx.Insert(insertRedirectURI, au...)
// 			if !rsus || rid <= 0 {
// 				tx.Rollback()
// 				fail = true
// 				break
// 			}
// 		}
// 	}
// 	return fail
// }

//UpdateClient UpdateClient
func (d *MySQLOauthDB) UpdateClient(client *odb.Client) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, client.Secret, client.Name, client.WebSite, client.Email, client.Enabled, client.Paid, client.ClientID)
	suc := d.DB.Update(updateClient, a...)
	return suc
}

//GetClient GetClient
func (d *MySQLOauthDB) GetClient(clientID int64) *odb.Client {
	var rtn *odb.Client
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, clientID)
	row := d.DB.Get(getClientByID, a...)
	if row != nil && len(row.Row) != 0 {
		foundRow := row.Row
		rtn = parseClientRow(&foundRow)
		// 	int64Val, err := strconv.ParseInt(foundRow[0], 10, 64)
		// 	if err == nil {
		// 		rtn.ClientID = int64Val
		// 		rtn.Secret = foundRow[1]
		// 		rtn.Name = foundRow[2]
		// 		rtn.WebSite = foundRow[3]
		// 		rtn.Email = foundRow[4]
		// 		rtn.Enabled, _ = strconv.ParseBool(foundRow[5])
		// 		rtn.Paid, _ = strconv.ParseBool(foundRow[6])
		//}
	}

	return rtn
}

//GetClients GetClients
func (d *MySQLOauthDB) GetClients() *[]odb.Client {
	var rtn []odb.Client
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	rows := d.DB.GetList(getClientsAll, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}

	}

	return &rtn
}

func parseClientRow(foundRow *[]string) *odb.Client {
	var rtn odb.Client
	if len(*foundRow) > 0 {
		int64Val, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		if err == nil {
			rtn.ClientID = int64Val
			rtn.Secret = (*foundRow)[1]
			rtn.Name = (*foundRow)[2]
			rtn.WebSite = (*foundRow)[3]
			rtn.Email = (*foundRow)[4]
			rtn.Enabled, _ = strconv.ParseBool((*foundRow)[5])
			rtn.Paid, _ = strconv.ParseBool((*foundRow)[6])
		}
	}
	return &rtn
}

//SearchClients SearchClients
func (d *MySQLOauthDB) SearchClients(name string) *[]odb.Client {
	var rtn []odb.Client
	if !d.testConnection() {
		d.DB.Connect()
	}
	var searchName = "%" + name + "%"
	//fmt.Println("searchName: ", searchName)
	var a []interface{}
	a = append(a, searchName)
	rows := d.DB.GetList(searchClientByName, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}

	}
	return &rtn
}

//DeleteClient DeleteClient
func (d *MySQLOauthDB) DeleteClient(clientID int64) bool {
	var suc = false
	if !d.testConnection() {
		d.DB.Connect()
	}
	var fail = false
	tx := d.DB.BeginTransaction()
	// var au []interface{}
	// au = append(au, clientID)
	// usuc := tx.Delete(deleteAllRedirectURI, au...)
	//d.Tx = tx
	usuc := d.DeleteClientAllRedirectURI(tx, clientID)
	if usuc {
		var a []interface{}
		a = append(a, clientID)
		sucu := tx.Delete(deleteClient, a...)
		if !sucu {
			fail = true
			tx.Rollback()
		}
	} else {
		fail = true
		tx.Rollback()
	}
	if !fail {
		suc = true
		tx.Commit()
	}

	return suc
}
