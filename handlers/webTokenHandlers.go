//Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

//TokenError TokenError
type TokenError struct {
	Error string `json:"error"`
}

//Token Token
func (h *OauthWebHandler) Token(w http.ResponseWriter, r *http.Request) {
	var atsuc bool
	var token *m.Token
	var tokenErr string
	var caseDef bool
	fmt.Println("token")
	grantType := r.FormValue("grant_type")
	fmt.Println("grant type: ", grantType)

	clientIDStr := r.FormValue("client_id")
	fmt.Println("clientIDStr: ", clientIDStr)

	clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
	fmt.Println("clientID: ", clientID)
	fmt.Println("err: ", err)
	if err == nil {
		switch grantType {
		case authorizationCodeGrantType:
			fmt.Println("grant type: ", grantType)
			secret := r.FormValue("client_secret")
			fmt.Println("secret: ", secret)
			code := r.FormValue("code")
			fmt.Println("code: ", code)
			redirectURI := r.FormValue("redirect_uri")
			fmt.Println("redirectURI: ", redirectURI)
			var actk m.AuthCodeTokenReq
			actk.ClientID = clientID
			actk.Secret = secret
			actk.Code = code
			actk.RedirectURI = redirectURI

			atsuc, token, tokenErr = h.Manager.GetAuthCodeToken(&actk)
			fmt.Println("atsuc: ", atsuc)

		// case passwordGrantType:
		// 	fmt.Println("grant type: ", grantType)
		// case credentialGrantType:
		// 	fmt.Println("grant type: ", grantType)
		// case refreshTokenGrantType:
		// 	fmt.Println("grant type: ", grantType)
		default:
			caseDef = true
		}

		if atsuc {
			h.SetContentType(w)
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(token)
			fmt.Fprint(w, string(resJSON))
		} else if caseDef {
			http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
		} else {
			var te TokenError
			te.Error = tokenErr
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(te)
			fmt.Fprint(w, string(resJSON))
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
	}
}
