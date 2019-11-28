package managers

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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

//px "github.com/Ulbora/GoProxy"

//Login Login
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID int64  `json:"clientId"`
}

//LoginRes LoginRes
type LoginRes struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

//UserLogin UserLogin
func (m *OauthManager) UserLogin(login *Login) bool {
	var rtn bool
	var url string
	if os.Getenv("AUTHENTICATION_SERVICE") != "" {
		url = os.Getenv("AUTHENTICATION_SERVICE")
	} else {
		url = authenticationServiceLocal
	}
	aJSON, _ := json.Marshal(login)	
	req, rErr := http.NewRequest("POST", url, bytes.NewBuffer(aJSON))
	if rErr != nil {
		log.Println("Request error: ", rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		var lres LoginRes
		fmt.Println("m.Proxy: ", m.Proxy)
		suc, code := m.Proxy.Do(req, &lres)
		fmt.Println("suc: ", suc)
		fmt.Println("code: ", code)
		fmt.Println("lres: ", lres)
		fmt.Println("lres code: ", lres.Code)
		fmt.Println("clientid: ", strconv.FormatInt(login.ClientID, 10))
		if suc && code == http.StatusOK && strconv.FormatInt(login.ClientID, 10) == lres.Code {
			rtn = lres.Valid
		}
	}
	return rtn
}
