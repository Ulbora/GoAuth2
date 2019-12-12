//Package handlers ...
package handlers

import (
	"fmt"
	m "github.com/Ulbora/GoAuth2/managers"
	"net/http"
	"strconv"
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

//Authorize Authorize
func (h *OauthWebHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	//h.Session.InitSessionStore()
	s, suc := h.getSession(r)
	if suc {
		loggedInAuth := s.Values["loggedIn"]
		userAuth := s.Values["user"]
		fmt.Println("loggedIn: ", loggedInAuth)
		fmt.Println("user: ", userAuth)

		respTypeAuth := r.URL.Query().Get("response_type")
		fmt.Println("respType: ", respTypeAuth)

		clientIDStrAuth := r.URL.Query().Get("client_id")
		fmt.Println("clientIDStr: ", clientIDStrAuth)

		clientIDAuth, idErr := strconv.ParseInt(clientIDStrAuth, 10, 64)
		fmt.Println("clientIDAuth: ", clientIDAuth)
		fmt.Println("idErr: ", idErr)

		redirectURLAuth := r.URL.Query().Get("redirect_uri")
		fmt.Println("redirURLAuth: ", redirectURLAuth)

		scopeAuth := r.URL.Query().Get("scope")
		fmt.Println("scopeAuth: ", scopeAuth)

		stateAuth := r.URL.Query().Get("state")
		fmt.Println("stateAuth: ", stateAuth)

		var ari AuthorizeRequestInfo
		ari.ResponseType = respTypeAuth
		ari.ClientID = clientIDAuth
		ari.RedirectURI = redirectURLAuth
		ari.Scope = scopeAuth
		ari.State = scopeAuth

		if loggedInAuth == true && userAuth != "" {
			fmt.Println("loggedIn: ", loggedInAuth)
			fmt.Println("user: ", userAuth)
			if respTypeAuth == codeRespType {
				var au m.AuthCode
				au.ClientID = clientIDAuth
				au.UserID = userAuth.(string)
				au.Scope = scopeAuth
				au.RedirectURI = redirectURLAuth
				authed := h.Manager.CheckAuthCodeApplicationAuthorization(&au)
				fmt.Println("authed: ", authed)
				if authed {
					ausuc, acode, acodeStr := h.Manager.AuthorizeAuthCode(&au)
					fmt.Println("ausuc: ", ausuc)
					fmt.Println("acode: ", acode)
					fmt.Println("acodeStr: ", acodeStr)
					if ausuc && acode != 0 && acodeStr != "" {
						http.Redirect(w, r, redirectURLAuth+"?code="+acodeStr+"&state="+stateAuth, http.StatusFound)
					} else {
						http.Redirect(w, r, accessDeniedErrorURL, http.StatusFound)
					}
				} else {
					s.Values["authReqInfo"] = ari
					s.Save(r, w)
					http.Redirect(w, r, authorizeAppURL, http.StatusFound)
				}
			} else if respTypeAuth == tokenRespType {
				var aut m.Implicit
				aut.ClientID = clientIDAuth
				aut.UserID = userAuth.(string)
				aut.Scope = scopeAuth
				aut.RedirectURI = redirectURLAuth
				iauthed := h.Manager.CheckImplicitApplicationAuthorization(&aut)
				fmt.Println("iauthed: ", iauthed)
				if iauthed {
					isuc, im := h.Manager.AuthorizeImplicit(&aut)
					if isuc && im.Token != "" {
						http.Redirect(w, r, redirectURLAuth+"?token="+im.Token+"&token_type=bearer&state="+stateAuth, http.StatusFound)
					} else {
						http.Redirect(w, r, accessDeniedErrorURL, http.StatusFound)
					}
				} else {
					s.Values["authReqInfo"] = ari
					s.Save(r, w)
					http.Redirect(w, r, authorizeAppURL, http.StatusFound)
				}
			} else {
				http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
			}
		} else {
			s.Values["authReqInfo"] = ari
			s.Save(r, w)
			http.Redirect(w, r, loginURL, http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
