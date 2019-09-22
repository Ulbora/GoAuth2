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
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//AddAuthorizationCode AddAuthorizationCode
func (d *MySQLOauthDB) AddAuthorizationCode(code *odb.AuthorizationCode, at *odb.AccessToken, rt *odb.RefreshToken, scopeList *[]string) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	tx := d.DB.BeginTransaction()
	var cont bool
	if rt.Token != "" {
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
		//at.RefreshTokenID = rtID
		atsuc, acID := d.AddAccessToken(tx, at)
		fmt.Println("atTk res: ", atsuc)
		fmt.Println("atTk id: ", acID)
		if atsuc {
			code.AccessTokenID = acID
			var a []interface{}
			a = append(a, code.ClientID, code.UserID, code.Expires, code.AccessTokenID, code.RandonAuthCode, code.AlreadyUsed)
			suc, id = tx.Insert(insertAuthCode, a...)
			fmt.Println("ac res: ", suc)
			fmt.Println("ac id: ", id)
			if suc {
				//add code for adding scopes
				var scSuc = true
				if scopeList != nil {
					for _, s := range *scopeList {
						var acs odb.AuthCodeScope
						acs.AuthorizationCode = id
						acs.Scope = s
						ssuc, sid := d.AddAuthCodeScope(tx, &acs)
						fmt.Println("scope res: ", ssuc)
						fmt.Println("scope id: ", sid)
						if !ssuc {
							scSuc = false
						}
					}
				}
				if scSuc {
					suc = true
					tx.Commit()
				} else {
					tx.Rollback()
				}
			} else {
				tx.Rollback()
			}
		}
	} else {
		tx.Rollback()
	}
	return suc, id
}

//GetAuthorizationCode GetAuthorizationCode
func (d *MySQLOauthDB) GetAuthorizationCode(clientID int64, userID string) *[]odb.AuthorizationCode {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.AuthorizationCode
	var a []interface{}
	a = append(a, clientID, userID)
	rows := d.DB.GetList(authorizationCodeGetByID, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseAuthCodeRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	// rtn := parseAuthCodeRow(&row.Row)
	fmt.Println("authCode: ", rtn)
	return &rtn
}

//DeleteAuthorizationCode DeleteAuthorizationCode
func (d *MySQLOauthDB) DeleteAuthorizationCode(clientID int64, userID string) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	//make this a list call
	acodeList := d.GetAuthorizationCode(clientID, userID)
	for _, acode := range *acodeList {
		if acode.AuthorizationCode > 0 {
			at := d.GetAccessToken(acode.AccessTokenID)
			var rtid int64
			if at.RefreshTokenID > 0 {
				rt := d.GetRefreshToken(at.RefreshTokenID)
				rtid = rt.ID
			}
			tx := d.DB.BeginTransaction()
			// authCodeRevokeProcessor.deleteAuthCodeRevoke do this
			rvkDel := d.DeleteAuthCodeRevolk(tx, acode.AuthorizationCode)
			if rvkDel {
				sdel := d.DeleteAuthCodeScopeList(tx, acode.AuthorizationCode)
				fmt.Println("delete scope: ", sdel)
				// d.DeleteAuthCodeScopeList(tx, acode.AuthorizationCode)
				if sdel {
					var a []interface{}
					a = append(a, clientID, userID)
					acdel := tx.Delete(deleteAuthCode, a...)
					if acdel {
						atdel := d.DeleteAccessToken(tx, acode.AccessTokenID)
						if atdel {
							var cont = true
							if rtid > 0 {
								cont = d.DeleteRefreshToken(tx, rtid)
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
				} else {
					tx.Rollback()
				}
			}
		}
	}
	return suc
}

func parseAuthCodeRow(foundRow *[]string) *odb.AuthorizationCode {
	var rtn odb.AuthorizationCode
	id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
	if err == nil {
		// rtn.AuthorizationCode = id
		cid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
		if err == nil {
			// rtn.ClientID = cid
			//uid, err := strconv.ParseInt((*foundRow)[2], 10, 64)
			//if err == nil {
			// rtn.UserID = uid
			cTime, err := time.Parse(odb.TimeFormat, (*foundRow)[3])
			if err == nil {
				atid, err := strconv.ParseInt((*foundRow)[4], 10, 64)
				if err == nil {
					rtn.AuthorizationCode = id
					rtn.ClientID = cid
					rtn.UserID = (*foundRow)[2]
					rtn.Expires = cTime
					rtn.AccessTokenID = atid
					rtn.RandonAuthCode = (*foundRow)[5]
					rtn.AlreadyUsed, _ = strconv.ParseBool((*foundRow)[6])
				}
			}
			//}
		}
	}
	return &rtn
}
