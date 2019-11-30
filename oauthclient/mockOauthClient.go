//Package oauthclient ...
package oauthclient

import (
	"net/http"

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

//MockOauthClient MockOauthClient
type MockOauthClient struct {
	Manager   m.Manager
	MockValid bool
}

//Authorize Authorize
func (o *MockOauthClient) Authorize(r *http.Request, c *Claim) bool {
	// var mrtn bool
	// var mhashed bool
	// mtokenHeader := r.Header.Get("Authorization")
	// mclientIDStr := r.Header.Get("clientId")
	// mclientID, _ := strconv.ParseInt(mclientIDStr, 10, 64)
	// muserID := r.Header.Get("userId")
	// mhashedStr := r.Header.Get("hashed")
	// if mhashedStr == "true" {
	// 	mhashed = true
	// }
	// fmt.Println("tokenHeader", mtokenHeader)
	// fmt.Println("clientIDStr", mclientIDStr)
	// fmt.Println("clientID", mclientID)
	// fmt.Println("userID", muserID)
	// fmt.Println("hashed", mhashed)
	// if mtokenHeader != "" {
	// 	mtokenArray := strings.Split(mtokenHeader, " ")
	// 	fmt.Println("tokenArray", mtokenArray)
	// 	if len(mtokenArray) == 2 {
	// 		fmt.Println("tokenArray[1]", mtokenArray[1])
	// 		mrtn = o.MockValid
	// 	}
	// }

	return o.MockValid
}

//GetNewClient GetNewClient
func (o *MockOauthClient) GetNewClient() Client {
	var c Client
	c = o
	return c
}
