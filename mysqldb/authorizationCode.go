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

//UpdateAuthorizationCode UpdateAuthorizationCode
func (d *MySQLOauthDB) UpdateAuthorizationCode(code *odb.AuthorizationCode) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, code.RandonAuthCode, code.AlreadyUsed, code.AuthorizationCode)
	suc := d.DB.Update(updateAuthCode, a...)
	return suc
}

//UpdateAuthorizationCodeAndToken UpdateAuthorizationCodeAndToken
func (d *MySQLOauthDB) UpdateAuthorizationCodeAndToken(code *odb.AuthorizationCode, at *odb.AccessToken) bool {
	var rtn bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	tx := d.DB.BeginTransaction()
	res := d.UpdateAccessToken(tx, at)
	if res {
		var a []interface{}
		a = append(a, code.Expires, code.AuthorizationCode)
		suc := tx.Update(updateAuthCodeToken, a...)
		if suc {
			rtn = true
			tx.Commit()
		} else {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
	}
	return rtn
}

//GetAuthorizationCode GetAuthorizationCode
func (d *MySQLOauthDB) GetAuthorizationCode(clientID int64, userID string) *[]odb.AuthorizationCode {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.AuthorizationCode
	var a []interface{}
	a = append(a, clientID, userID)
	rows := d.DB.GetList(getByAuthorizationCodeClientUser, a...)
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

//GetAuthorizationCodeByScope GetAuthorizationCodeByScope
func (d *MySQLOauthDB) GetAuthorizationCodeByScope(clientID int64, userID string, scope string) *[]odb.AuthorizationCode {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.AuthorizationCode
	var a []interface{}
	a = append(a, clientID, userID, scope)
	rows := d.DB.GetList(getAuthorizationCodeByClientUserScope, a...)
	fmt.Println("rows in getbyscope: ", rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		fmt.Println("foundRows in getbyscope: ", foundRows)
		for r := range foundRows {
			foundRow := foundRows[r]
			fmt.Println("foundRow in getbyscope: ", foundRow)
			ac, err := strconv.ParseInt((foundRow)[0], 10, 64)
			if err == nil {
				cid, err := strconv.ParseInt((foundRow)[1], 10, 64)
				if err == nil {
					var rtnc odb.AuthorizationCode
					rtnc.AuthorizationCode = ac
					rtnc.ClientID = cid
					rtnc.UserID = userID
					rtnc.Scope = (foundRow)[2]
					//rtnc.Expires = cTime
					//rtnc.AccessTokenID = atid
					rtnc.RandonAuthCode = (foundRow)[3]
					rtnc.AlreadyUsed, _ = strconv.ParseBool((foundRow)[4])
					fmt.Println("rtnc in getbyscope: ", rtnc)
					rtn = append(rtn, rtnc)
				}
			}
			//rowContent := parseAuthCodeRow(&foundRow)

		}
	}
	// rtn := parseAuthCodeRow(&row.Row)
	fmt.Println("authCode in scope: ", rtn)
	return &rtn
}

//GetAuthorizationCodeByCode GetAuthorizationCodeByCode
func (d *MySQLOauthDB) GetAuthorizationCodeByCode(code string) *odb.AuthorizationCode {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, code)
	row := d.DB.Get(getAuthorizationCodeByCode, a...)
	rtn := parseAuthCodeRow(&row.Row)
	fmt.Println("authCode: ", rtn)
	return rtn
}

//DeleteAuthorizationCode DeleteAuthorizationCode
func (d *MySQLOauthDB) DeleteAuthorizationCode(clientID int64, userID string) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	//make this a list call
	acodeList := d.GetAuthorizationCode(clientID, userID)
	fmt.Println("auth code list: ", acodeList)
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
			fmt.Println("delete refresh token: ", rvkDel)
			if rvkDel {
				sdel := d.DeleteAuthCodeScopeList(tx, acode.AuthorizationCode)
				fmt.Println("delete scope: ", sdel)
				// d.DeleteAuthCodeScopeList(tx, acode.AuthorizationCode)
				if sdel {
					var a []interface{}
					a = append(a, clientID, userID)
					acdel := tx.Delete(deleteAuthCode, a...)
					fmt.Println("delete authCode: ", acdel)
					if acdel {
						atdel := d.DeleteAccessToken(tx, acode.AccessTokenID)
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
				} else {
					tx.Rollback()
				}
			}
		}
	}
	return suc
}

func parseAuthCodeRow(foundRow *[]string) *odb.AuthorizationCode {
	fmt.Println("foundRow in parseAuthCodeRow: ", foundRow)
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
			fmt.Println("time error:", err)
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
