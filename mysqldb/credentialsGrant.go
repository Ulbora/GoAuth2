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

	"fmt"
	"strconv"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//AddCredentialsGrant AddCredentialsGrant
func (d *MySQLOauthDB) AddCredentialsGrant(cg *odb.CredentialsGrant, at *odb.AccessToken) (bool, int64) {
	var suc = false
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	tx := d.DB.BeginTransaction()
	atsuc, acID := d.AddAccessToken(tx, at)
	fmt.Println("atTk res: ", atsuc)
	fmt.Println("atTk id: ", acID)
	if atsuc {
		cg.AccessTokenID = acID
		var a []interface{}
		a = append(a, cg.ClientID, cg.AccessTokenID)
		suc, id = tx.Insert(insertCredentialsGrant, a...)
		fmt.Println("ig res: ", suc)
		fmt.Println("ig id: ", id)
		if suc {
			tx.Commit()
		} else {
			suc = false
			id = 0
			fmt.Println("rolling back suc: ", suc)
			tx.Rollback()
		}
	} else {
		tx.Rollback()
	}
	return suc, id
}

//GetCredentialsGrant GetCredentialsGrant
func (d *MySQLOauthDB) GetCredentialsGrant(clientID int64) *[]odb.CredentialsGrant {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.CredentialsGrant
	var a []interface{}
	a = append(a, clientID)
	rows := d.DB.GetList(getCredentialsGrant, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		fmt.Println("foundRows in getbyscope: ", foundRows)
		for r := range foundRows {
			foundRow := foundRows[r]
			if len(foundRow) > 0 {
				fmt.Println("foundRow in getbyscope: ", foundRow)
				cgID, err := strconv.ParseInt((foundRow)[0], 10, 64)
				if err == nil {
					cid, err := strconv.ParseInt((foundRow)[1], 10, 64)
					if err == nil {
						tid, err := strconv.ParseInt((foundRow)[2], 10, 64)
						if err == nil {
							var rtnc odb.CredentialsGrant
							rtnc.ID = cgID
							rtnc.ClientID = cid
							rtnc.AccessTokenID = tid
							fmt.Println("rtnc in getbyscope: ", rtnc)
							rtn = append(rtn, rtnc)
						}
					}
				}
			}
		}
	}
	fmt.Println("CredentialsGrant list: ", rtn)
	return &rtn
}

//DeleteCredentialsGrant DeleteCredentialsGrant
func (d *MySQLOauthDB) DeleteCredentialsGrant(clientID int64) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	cgList := d.GetCredentialsGrant(clientID)
	fmt.Println("cgList: ", cgList)
	if len(*cgList) == 0 {
		suc = true
	} else {
		for _, cg := range *cgList {
			if cg.ID > 0 {
				tx := d.DB.BeginTransaction()
				var a []interface{}
				a = append(a, cg.ID)
				cgdel := tx.Delete(deleteCredentialsGrant, a...)
				fmt.Println("delete cg: ", cgdel)
				if cgdel {
					atdel := d.DeleteAccessToken(tx, cg.AccessTokenID)
					fmt.Println("delete AccessToken: ", atdel)
					if atdel {
						suc = true
						tx.Commit()
					} else {
						tx.Rollback()
					}
				} else {
					tx.Rollback()
				}
			}
		}
	}
	return suc
}
