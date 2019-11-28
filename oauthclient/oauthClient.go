package oauthclient

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	m "github.com/Ulbora/GoAuth2/managers"
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
	Manager m.Manager
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
	fmt.Println("tokenHeader", tokenHeader)
	fmt.Println("clientIDStr", clientIDStr)
	fmt.Println("clientID", clientID)
	fmt.Println("userID", userID)
	fmt.Println("hashed", hashed)
	if tokenHeader != "" {
		tokenArray := strings.Split(tokenHeader, " ")
		fmt.Println("tokenArray", tokenArray)
		if len(tokenArray) == 2 {
			fmt.Println("tokenArray[1]", tokenArray[1])
			var vr m.ValidateAccessTokenReq
			vr.AccessToken = tokenArray[1]
			vr.Hashed = hashed
			vr.UserID = userID
			vr.ClientID = clientID
			vr.Role = c.Role
			vr.UserID = c.URL
			vr.Scope = c.Scope
			rtn = o.Manager.ValidateAccessToken(&vr)
			fmt.Println("valid: ", rtn)
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
