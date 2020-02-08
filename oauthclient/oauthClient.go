//Package oauthclient ...
package oauthclient

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	cp "github.com/Ulbora/GoAuth2/compresstoken"
	m "github.com/Ulbora/GoAuth2/managers"
	lg "github.com/Ulbora/Level_Logger"
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

//OauthClient OauthClient
type OauthClient struct {
	Manager         m.Manager
	TokenCompressed bool
	JwtCompress     cp.JwtCompress
	Log             *lg.Logger
}

//Authorize Authorize
func (o *OauthClient) Authorize(r *http.Request, c *Claim) bool {
	var rtn bool
	var hashed bool
	tokenHeader := r.Header.Get("Authorization")
	clientIDStr := r.Header.Get("clientId")
	clientID, _ := strconv.ParseInt(clientIDStr, 10, 64)
	userID := r.Header.Get("userId")
	hashedStr := r.Header.Get("hashed")
	if hashedStr == "true" {
		hashed = true
	}
	//fmt.Println("tokenHeader", tokenHeader)
	o.Log.Debug("clientIDStr", clientIDStr)
	o.Log.Debug("clientID", clientID)
	o.Log.Debug("userID", userID)
	o.Log.Debug("hashed", hashed)
	if tokenHeader != "" {
		tokenArray := strings.Split(tokenHeader, " ")
		//fmt.Println("tokenArray", tokenArray)
		if len(tokenArray) == 2 {
			var token string
			if o.TokenCompressed {
				token = o.JwtCompress.UnCompressJwt(tokenArray[1])
			} else {
				token = tokenArray[1]
			}
			fmt.Println("loglevel: ", o.Log.LogLevel)
			o.Log.Info("token:", token)
			var vr m.ValidateAccessTokenReq
			vr.AccessToken = token
			vr.Hashed = hashed
			vr.UserID = userID
			vr.ClientID = clientID
			vr.Role = c.Role
			vr.URI = c.URL
			vr.Scope = c.Scope
			rtn = o.Manager.ValidateAccessToken(&vr)
			o.Log.Debug("valid: ", rtn)
		}
	}
	return rtn
}

//GetNewClient GetNewClient
func (o *OauthClient) GetNewClient() Client {
	var c Client
	c = o
	return c
}
