//Package handlers ...
package handlers

import (
	"fmt"

	b64 "encoding/base64"
	"encoding/json"

	m "github.com/Ulbora/GoAuth2/managers"
	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
	oa "github.com/Ulbora/GoAuth2/oauthclient"
	rc "github.com/Ulbora/GoAuth2/rolecontrol"
	db "github.com/Ulbora/dbinterface"
	// mdb "github.com/Ulbora/dbinterface_mysql"
)

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

//UseWebHandler UseWebHandler
func UseWebHandler(dbi db.Database) *OauthWebHandler {
	var oauthManagerw m.OauthManager
	var oauthMySqldbw msdb.MySQLOauthDB
	oauthMySqldbw.DB = dbi
	var oauthDbw odb.Oauth2DB
	oauthDbw = &oauthMySqldbw
	oauthManagerw.Db = oauthDbw

	var wh OauthWebHandler
	wh.Manager = &oauthManagerw
	return &wh
}

//UseRestHandler UseRestHandler
func UseRestHandler(dbi db.Database, assets string) *OauthRestHandler {
	var oauthManager m.OauthManager
	var oauthMySqldb msdb.MySQLOauthDB
	oauthMySqldb.DB = dbi
	var oauthDb odb.Oauth2DB
	oauthDb = &oauthMySqldb
	oauthManager.Db = oauthDb

	var rh OauthRestHandler
	rh.Manager = &oauthManager

	var clt oa.OauthClient
	clt.Manager = &oauthManager
	rh.Client = &clt

	var curs []rc.ControlledURL
	if assets != "" {
		//parse assets
		sDec, err := b64.StdEncoding.DecodeString(assets)
		if err == nil {
			json.Unmarshal(sDec, &curs)
			fmt.Println("curs: ", curs)
		}
	}
	var ac rc.OauthAssets
	ac.AddControledURLs(&curs)
	rh.AssetControl = &ac
	return &rh
}
