//Package handlers ...
package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"

	m "github.com/Ulbora/GoAuth2/managers"
	gs "github.com/Ulbora/go-sessions"
	"github.com/gorilla/sessions"
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

const (
	accessDeniedErrorURL = "/oauthError?error=access_denied"
	authorizeAppURL      = "/authorizeApp"
	invalidGrantErrorURL = "/oauthError?error=invalid_grant"
	loginURL             = "/login"
	loginFailedURL       = "/login?error=Login Failed"

	authorizeHTML  = "authorizeApp.html"
	indexHTML      = "index.html"
	loginHTML      = "login.html"
	oauthErrorHTML = "oauthError.html"
)

//OauthWebHandler OauthWebHandler
type OauthWebHandler struct {
	Manager   m.Manager
	Session   gs.GoSession
	Templates *template.Template
	Store     *sessions.CookieStore
	//SessInit  bool
}

//AuthorizeRequestInfo AuthorizeRequestInfo
type AuthorizeRequestInfo struct {
	ResponseType string
	ClientID     int64
	RedirectURI  string
	Scope        string
	State        string
}

//GetNewWebHandler GetNewWebHandler
func (h *OauthWebHandler) GetNewWebHandler() WebHandler {
	var wh WebHandler
	wh = h
	return wh
}

func (h *OauthWebHandler) getSession(r *http.Request) (*sessions.Session, bool) {
	//fmt.Println("getSession--------------------------------------------------")
	var suc bool
	var srtn *sessions.Session
	if h.Store == nil {
		h.Session.Name = "goauth2"
		h.Session.MaxAge = 3600
		h.Store = h.Session.InitSessionStore()
		//errors without this
		gob.Register(&AuthorizeRequestInfo{})
	}
	if r != nil {
		// fmt.Println("secure in getSession", h.Session.Secure)
		// fmt.Println("name in getSession", h.Session.Name)
		// fmt.Println("MaxAge in getSession", h.Session.MaxAge)
		// fmt.Println("SessionKey in getSession", h.Session.SessionKey)

		//h.Session.HTTPOnly = true

		//h.Session.InitSessionStore()
		s, err := h.Store.Get(r, h.Session.Name)
		//s, err := store.Get(r, "temp-name")
		//s, err := store.Get(r, "goauth2")

		loggedInAuth := s.Values["loggedIn"]
		userAuth := s.Values["user"]
		fmt.Println("loggedIn: ", loggedInAuth)
		fmt.Println("user: ", userAuth)

		larii := s.Values["authReqInfo"]
		fmt.Println("arii-----login", larii)

		fmt.Println("session error in getSession: ", err)
		if err == nil {
			suc = true
			srtn = s
		}
	}
	//fmt.Println("exit getSession--------------------------------------------------")
	return srtn, suc
}

//SetContentType SetContentType
func (h *OauthWebHandler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//SetSecurityHeader SetSecurityHeader
func (h *OauthWebHandler) SetSecurityHeader(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
}
