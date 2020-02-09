//Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "github.com/Ulbora/GoAuth2/managers"
	au "github.com/Ulbora/auth_interface"
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
	h.Log.Debug("token")
	grantType := r.FormValue("grant_type")
	h.Log.Debug("grant type: ", grantType)

	clientIDStr := r.FormValue("client_id")
	h.Log.Debug("clientIDStr: ", clientIDStr)

	clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
	h.Log.Debug("clientID: ", clientID)
	h.Log.Debug("err: ", err)
	if err == nil {
		switch grantType {
		case authorizationCodeGrantType:
			h.Log.Debug("grant type: ", grantType)
			secret := r.FormValue("client_secret")
			h.Log.Debug("secret: ", secret)
			code := r.FormValue("code")
			h.Log.Debug("code: ", code)
			redirectURI := r.FormValue("redirect_uri")
			h.Log.Debug("redirectURI: ", redirectURI)
			var actk m.AuthCodeTokenReq
			actk.ClientID = clientID
			actk.Secret = secret
			actk.Code = code
			actk.RedirectURI = redirectURI

			atsuc, token, tokenErr = h.Manager.GetAuthCodeToken(&actk)
			h.Log.Debug("atsuc: ", atsuc)

		case passwordGrantType:
			h.Log.Debug("grant type: ", grantType)
			username := r.FormValue("username")
			h.Log.Debug("username: ", username)
			password := r.FormValue("password")
			h.Log.Debug("password: ", password)
			var lg au.Login
			lg.ClientID = clientID
			lg.Username = username
			lg.Password = password
			suc := h.Manager.UserLogin(&lg)
			h.Log.Debug("login suc", suc)
			if suc {
				var pt m.PasswordTokenReq
				pt.Username = username
				pt.ClientID = clientID
				atsuc, token, tokenErr = h.Manager.GetPasswordToken(&pt)
				h.Log.Debug("atsuc: ", atsuc)
			} else {
				tokenErr = unauthorizedClientError
			}
		case credentialGrantType:
			h.Log.Debug("grant type: ", grantType)
			secret := r.FormValue("client_secret")
			h.Log.Debug("secret: ", secret)
			var ct m.CredentialsTokenReq
			ct.ClientID = clientID
			ct.Secret = secret
			atsuc, token, tokenErr = h.Manager.GetCredentialsToken(&ct)
			h.Log.Debug("atsuc: ", atsuc)
		case refreshTokenGrantType:
			h.Log.Debug("grant type: ", grantType)
			secret := r.FormValue("client_secret")
			h.Log.Debug("secret: ", secret)
			refToken := r.FormValue("refresh_token")
			h.Log.Debug("refToken: ", refToken)
			if secret != "" {
				var acrt m.RefreshTokenReq
				acrt.ClientID = clientID
				acrt.Secret = secret
				acrt.RefreshToken = refToken
				atsuc, token, tokenErr = h.Manager.GetAuthCodeAccesssTokenWithRefreshToken(&acrt)
				h.Log.Debug("atsuc: ", atsuc)
			} else {
				var pwrt m.RefreshTokenReq
				pwrt.ClientID = clientID
				pwrt.RefreshToken = refToken
				atsuc, token, tokenErr = h.Manager.GetPasswordAccesssTokenWithRefreshToken(&pwrt)
			}
		default:
			caseDef = true
		}
		if atsuc {
			h.SetContentType(w)
			h.SetSecurityHeader(w)
			w.WriteHeader(http.StatusOK)
			if h.TokenCompressed {
				token.AccessToken = h.JwtCompress.CompressJwt(token.AccessToken)
			}
			resJSON, _ := json.Marshal(token)
			fmt.Fprint(w, string(resJSON))
		} else if caseDef {
			http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
		} else {
			var te TokenError
			te.Error = tokenErr
			h.SetContentType(w)
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(te)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
	}
}
