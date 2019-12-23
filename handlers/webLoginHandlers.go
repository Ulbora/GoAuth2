//Package handlers ...
package handlers

import (
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

//var store = sessions.NewCookieStore([]byte("12345"))

//Login  login handler
func (h *OauthWebHandler) Login(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("login --------------------------------------------------")
	// s, _ := h.getSession(r)
	// fmt.Println("s", s)
	// fmt.Println("store s", s.Store())
	// fmt.Println("name in getSession s", s.Name())
	// fmt.Println("id getSession s", s.ID)
	// fmt.Println("Options in getSession s", s.Options)
	// fmt.Println("session name", s.Name())
	// fmt.Println("SessionKey in getSession", h.Session.SessionKey)

	// loggedInAuth := s.Values["loggedIn"]
	// userAuth := s.Values["user"]
	// fmt.Println("loggedIn: ", loggedInAuth)
	// fmt.Println("user: ", userAuth)

	// larii := s.Values["authReqInfo"]

	// fmt.Println("arii-----login", larii)
	// someTest := s.Values["testval"]

	// fmt.Println("someTest-----login", someTest)

	//session, _ := store.Get(r, "temp-name")
	//fmt.Println("store options ", store.Options)

	//tval := session.Values["authReqInfo"]
	//fmt.Println("tval-----login", tval)

	loginErr := r.URL.Query().Get("error")
	var lpg PageParams
	lpg.Error = loginErr
	lpg.Title = loginPageTitle
	fmt.Println("login params: ", lpg)
	h.Templates.ExecuteTemplate(w, loginHTML, &lpg)
}

//LoginUser LoginUser
func (h *OauthWebHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("in login form submit--------------------------------------")
	s, suc := h.getSession(r)
	if suc {
		// fmt.Println("store s", s.Store())
		// fmt.Println("name in getSession s", s.Name())
		// fmt.Println("id getSession s", s.ID)
		// fmt.Println("Options in getSession s", s.Options)
		// fmt.Println("session name", s.Name())
		// fmt.Println("SessionKey in getSession", h.Session.SessionKey)

		larii := s.Values["authReqInfo"]
		fmt.Println("arii", larii)
		if larii != nil {
			lari := larii.(*AuthorizeRequestInfo)
			fmt.Println("ari", lari)
			username := r.FormValue("username")
			password := r.FormValue("password")
			fmt.Println("username", username)
			fmt.Println("password", password)
			var lg m.Login
			lg.ClientID = lari.ClientID
			lg.Username = username
			lg.Password = password
			suc := h.Manager.UserLogin(&lg)
			fmt.Println("login suc", suc)
			if suc {
				if lari.ResponseType == codeRespType || lari.ResponseType == tokenRespType {
					s.Values["loggedIn"] = true
					s.Values["user"] = username
					serr := s.Save(r, w)
					fmt.Println("serr", serr)
					//session, sserr := store.Get(r, "temp-name")
					//fmt.Println("sserr", sserr)
					//session.Store()
					//session.Options.Path = "/oauth/"
					//session.Values["loggedIn"] = true
					//fmt.Println("store", session.Store())
					//session.Save(r, w)

					clintStr := strconv.FormatInt(lari.ClientID, 10)
					http.Redirect(w, r, "/oauth/authorize?response_type="+lari.ResponseType+"&client_id="+clintStr+"&redirect_uri="+lari.RedirectURI+"&scope="+lari.Scope+"&state="+lari.State, http.StatusFound)
				} else {
					http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
				}
			} else {
				http.Redirect(w, r, loginFailedURL, http.StatusFound)
			}
		} else {
			// http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
			http.Redirect(w, r, loginFailedURL, http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
