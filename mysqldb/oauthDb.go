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
	"log"
	"strconv"

	dbi "github.com/Ulbora/dbinterface"
)

//MySQLOauthDB MySQLOauthDB
type MySQLOauthDB struct {
	DB dbi.Database
}

func (d *MySQLOauthDB) testConnection() bool {
	log.Println("in testConnection")
	var rtn = false
	var a []interface{}
	log.Println("d.DB: ", d.DB)
	rowPtr := d.DB.Test(oauthTest, a...)
	log.Println("after testConnection test")
	if len(rowPtr.Row) != 0 {
		foundRow := rowPtr.Row
		int64Val, err := strconv.ParseInt(foundRow[0], 10, 0)
		//log.Print("Records found during test ")
		//log.Println("Records found during test :", int64Val)
		if err != nil {
			log.Print(err)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}
