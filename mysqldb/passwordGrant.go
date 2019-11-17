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
	"fmt"
	"strconv"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//AddPasswordGrant AddPasswordGrant
func (d *MySQLOauthDB) AddPasswordGrant(pwg *odb.PasswordGrant, at *odb.AccessToken, rt *odb.RefreshToken) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	tx := d.DB.BeginTransaction()
	var cont bool
	if rt != nil && rt.Token != "" {
		rtsuc, rtID := d.AddRefreshToken(tx, rt)
		fmt.Println("refTk res: ", rtsuc)
		fmt.Println("refTk id: ", rtID)
		if rtsuc {
			at.RefreshTokenID = rtID
			cont = true
		}
	} else {
		cont = true
	}
	if cont {
		atsuc, atID := d.AddAccessToken(tx, at)
		fmt.Println("atTk res: ", atsuc)
		fmt.Println("atTk id: ", atID)
		if atsuc {
			pwg.AccessTokenID = atID
			var a []interface{}
			a = append(a, pwg.ClientID, pwg.UserID, pwg.AccessTokenID)
			suc, id = tx.Insert(insertPasswordGrant, a...)
			fmt.Println("pwg res: ", suc)
			fmt.Println("pwg id: ", id)
			if suc {
				tx.Commit()
			} else {
				id = 0
				fmt.Println("pw grant rolling back: ", suc)
				tx.Rollback()
			}
		} else {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
	}
	return suc, id
}

//GetPasswordGrant GetPasswordGrant
func (d *MySQLOauthDB) GetPasswordGrant(clientID int64, userID string) *[]odb.PasswordGrant {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.PasswordGrant
	var a []interface{}
	a = append(a, clientID, userID)
	rows := d.DB.GetList(getPasswordGrant, a...)
	fmt.Println("rows in getbyscope: ", rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		fmt.Println("foundRows in getbyscope: ", foundRows)
		for r := range foundRows {
			foundRow := foundRows[r]
			if len(foundRow) > 0 {
				fmt.Println("foundRow in getbyscope: ", foundRow)
				pgID, err := strconv.ParseInt((foundRow)[0], 10, 64)
				if err == nil {
					cid, err := strconv.ParseInt((foundRow)[1], 10, 64)
					if err == nil {
						tid, err := strconv.ParseInt((foundRow)[3], 10, 64)
						if err == nil {
							var rtnc odb.PasswordGrant
							rtnc.ID = pgID
							rtnc.ClientID = cid
							rtnc.UserID = (foundRow)[2]
							rtnc.AccessTokenID = tid
							fmt.Println("rtnc in getbyscope: ", rtnc)
							rtn = append(rtn, rtnc)
						}
					}
				}
			}
		}
	}
	fmt.Println("pw grant: ", rtn)
	return &rtn
}

//DeletePasswordGrant DeletePasswordGrant
func (d *MySQLOauthDB) DeletePasswordGrant(clientID int64, userID string) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	pwgList := d.GetPasswordGrant(clientID, userID)
	fmt.Println("pwgList: ", pwgList)
	for _, pw := range *pwgList {
		if pw.ID > 0 {
			at := d.GetAccessToken(pw.AccessTokenID)
			fmt.Println("at: ", at)
			var rtid int64
			if at.RefreshTokenID > 0 {
				rt := d.GetRefreshToken(at.RefreshTokenID)
				rtid = rt.ID
			}
			tx := d.DB.BeginTransaction()
			var a []interface{}
			a = append(a, pw.ID)
			pwgdel := tx.Delete(deletePasswordGrantByID, a...)
			fmt.Println("delete pwg: ", pwgdel)
			if pwgdel {
				atdel := d.DeleteAccessToken(tx, pw.AccessTokenID)
				fmt.Println("delete AccessToken: ", atdel)
				if atdel {
					var cont = true
					if rtid > 0 {
						cont = d.DeleteRefreshToken(tx, rtid)
						fmt.Println("delete RefreshToken: ", cont)
					}
					if cont {
						suc = true
						tx.Commit()
					} else {
						tx.Rollback()
					}
				} else {
					tx.Rollback()
				}
			} else {
				tx.Rollback()
			}
		}
	}
	return suc
}
