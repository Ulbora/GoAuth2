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

// AddImplicitGrant AddImplicitGrant
func (d *MySQLOauthDB) AddImplicitGrant(ig *odb.ImplicitGrant, at *odb.AccessToken, scopeList *[]string) (bool, int64) {
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
		ig.AccessTokenID = acID
		var a []interface{}
		a = append(a, ig.ClientID, ig.UserID, ig.AccessTokenID)
		suc, id = tx.Insert(insertImplicitGrant, a...)
		fmt.Println("ig res: ", suc)
		fmt.Println("ig id: ", id)
		if suc {
			var scSuc = true
			if scopeList != nil {
				for _, s := range *scopeList {
					var igs odb.ImplicitScope
					igs.ImplicitGrantID = id
					igs.Scope = s
					ssuc, sid := d.AddImplicitGrantScope(tx, &igs)
					fmt.Println("scope res: ", ssuc)
					fmt.Println("scope id: ", sid)
					if !ssuc {
						fmt.Println("scope failed rolling back res: ", ssuc)
						scSuc = false
					}
				}
			}
			if scSuc {
				suc = true
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
	} else {
		tx.Rollback()
	}
	return suc, id
}

//GetImplicitGrant GetImplicitGrant
func (d *MySQLOauthDB) GetImplicitGrant(clientID int64, userID string) *[]odb.ImplicitGrant {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.ImplicitGrant
	var a []interface{}
	a = append(a, clientID, userID)
	rows := d.DB.GetList(getImplicitGrant, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseImplicitGrantRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	// rtn := parseAuthCodeRow(&row.Row)
	fmt.Println("ImplicitGrant list: ", rtn)
	return &rtn
}

//GetImplicitGrantByScope GetImplicitGrantByScope
func (d *MySQLOauthDB) GetImplicitGrantByScope(clientID int64, userID string, scope string) *[]odb.ImplicitGrant {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []odb.ImplicitGrant
	var a []interface{}
	a = append(a, clientID, userID, scope)
	rows := d.DB.GetList(getImplicitGrantByScope, a...)
	fmt.Println("rows in getbyscope: ", rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		fmt.Println("foundRows in getbyscope: ", foundRows)
		for r := range foundRows {
			foundRow := foundRows[r]
			fmt.Println("foundRow in getbyscope: ", foundRow)
			id, err := strconv.ParseInt((foundRow)[0], 10, 64)
			if err == nil {
				cid, err := strconv.ParseInt((foundRow)[1], 10, 64)
				if err == nil {
					aid, err := strconv.ParseInt((foundRow)[3], 10, 64)
					if err == nil {
						var rtnc odb.ImplicitGrant
						rtnc.ID = id
						rtnc.ClientID = cid
						rtnc.UserID = userID
						rtnc.Scope = (foundRow)[2]
						rtnc.AccessTokenID = aid
						fmt.Println("rtnc in getbyscope: ", rtnc)
						rtn = append(rtn, rtnc)
					}

				}
			}
		}
	}
	fmt.Println("ImplicitGrant in scope: ", rtn)
	return &rtn
}

//DeleteImplicitGrant DeleteImplicitGrant
func (d *MySQLOauthDB) DeleteImplicitGrant(clientID int64, userID string) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	igList := d.GetImplicitGrant(clientID, userID)
	fmt.Println("ImplicitGrant list in delete: ", igList)
	for _, ig := range *igList {
		if ig.ID > 0 {
			//at := d.GetAccessToken(ig.AccessTokenID)
			tx := d.DB.BeginTransaction()
			sdel := d.DeleteImplicitGrantScopeList(tx, ig.ID)
			fmt.Println("delete scope: ", sdel)
			if sdel {
				var a []interface{}
				a = append(a, ig.ID)
				igdel := tx.Delete(deleteImplicitGrantByID, a...)
				fmt.Println("delete ImplicitGrant: ", igdel)
				if igdel {
					atdel := d.DeleteAccessToken(tx, ig.AccessTokenID)
					fmt.Println("delete AccessToken: ", atdel)
					//if atdel {
					if atdel {
						suc = true
						tx.Commit()
					} else {
						tx.Rollback()
					}
					//}
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

func parseImplicitGrantRow(foundRow *[]string) *odb.ImplicitGrant {
	fmt.Println("foundRow in parseImplicitGrantRow: ", foundRow)
	var rtn odb.ImplicitGrant
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		if err == nil {
			// rtn.ImplicitGrant = id
			cid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			if err == nil {
				atid, err := strconv.ParseInt((*foundRow)[3], 10, 64)
				if err == nil {
					rtn.ID = id
					rtn.ClientID = cid
					rtn.UserID = (*foundRow)[2]
					rtn.AccessTokenID = atid
				}
			}
		}
	}
	return &rtn
}
