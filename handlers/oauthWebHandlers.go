//Package handlers ...
package handlers

import (
	"net/http"

	m "github.com/Ulbora/GoAuth2/managers"
	gs "github.com/Ulbora/go-sessions"
	ses "github.com/gorilla/sessions"
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
)

//OauthWebHandler OauthWebHandler
type OauthWebHandler struct {
	Manager m.Manager
	Session gs.GoSession
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

func (h *OauthWebHandler) getSession(r *http.Request) (*ses.Session, bool) {
	var suc bool
	h.Session.InitSessionStore()
	s, err := h.Session.GetSession(r)
	if err == nil {
		suc = true
	}
	return s, suc
}
