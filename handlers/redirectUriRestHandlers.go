package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "github.com/Ulbora/GoAuth2/managers"
	oc "github.com/Ulbora/GoAuth2/oauthclient"
	"github.com/gorilla/mux"
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

//AddRedirectURI AddRedirectURI
func (h *OauthRestHandler) AddRedirectURI(w http.ResponseWriter, r *http.Request) {
	var cr m.ClientRedirectURI
	rusuc, ruerr := h.ProcessBody(r, &cr)
	fmt.Println("rusuc: ", rusuc)
	fmt.Println("cr: ", cr)
	fmt.Println("ruerr: ", ruerr)
	if rusuc && ruerr == nil {

		//url of this endpoint
		var addRURL = "/ulbora/rs/clientRedirectUri/add"

		var rucl oc.Claim
		rucl.Role = "admin"
		rucl.URL = addRURL
		rucl.Scope = "write"
		fmt.Println("client: ", h.Client)

		//check that jwt token user role has permission to use the url of this endpoint
		auth := h.Client.Authorize(r, &rucl)

		if auth {
			// w.Header().Set("Content-Type", "application/json")
			h.SetContentType(w)
			rdURIContOk := h.CheckContent(r)
			fmt.Println("conOk: ", rdURIContOk)
			if !rdURIContOk {
				http.Error(w, "json required", http.StatusUnsupportedMediaType)
			} else {
				ruSuc, ruID := h.Manager.AddClientRedirectURI(&cr)
				fmt.Println("ruSuc: ", ruSuc)
				fmt.Println("ruID: ", ruID)
				var rtn ResponseID
				if ruSuc && ruID != 0 {
					rtn.Success = ruSuc
					rtn.ID = ruID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		} else {
			var rurtn ResponseID
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(rurtn)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Error(w, ruerr.Error(), http.StatusBadRequest)
	}
}

//GetRedirectURIList GetRedirectURIList
func (h *OauthRestHandler) GetRedirectURIList(w http.ResponseWriter, r *http.Request) {
	var getRURL = "/ulbora/rs/clientRedirectUri/list"

	var rugcl oc.Claim
	rugcl.Role = "admin"
	rugcl.URL = getRURL
	rugcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &rugcl)
	if auth {
		//var id string
		h.SetContentType(w)
		rugvars := mux.Vars(r)
		fmt.Println("vars: ", len(rugvars))
		if rugvars != nil && len(rugvars) != 0 {
			var rugclientIDStr = rugvars["clientId"]
			fmt.Println("vars: ", rugvars)
			clientID, rugidErr := strconv.ParseInt(rugclientIDStr, 10, 64)
			if clientID != 0 && rugidErr == nil {
				fmt.Println("clientID: ", clientID)
				getrul := h.Manager.GetClientRedirectURIList(clientID)
				fmt.Println("getrul: ", getrul)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getrul)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteRedirectURI DeleteRedirectURI
func (h *OauthRestHandler) DeleteRedirectURI(w http.ResponseWriter, r *http.Request) {
	var drURL = "/ulbora/rs/clientRedirectUri/delete"

	var rudcl oc.Claim
	rudcl.Role = "admin"
	rudcl.URL = drURL
	rudcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &rudcl)
	if auth {
		//var id string
		h.SetContentType(w)
		druvars := mux.Vars(r)
		fmt.Println("vars: ", len(druvars))
		if druvars != nil && len(druvars) != 0 {
			var druidStr = druvars["id"]
			fmt.Println("vars delete: ", druidStr)
			id, idErr := strconv.ParseInt(druidStr, 10, 64)
			fmt.Println("id delete: ", id)
			if id != 0 && idErr == nil {
				fmt.Println("id: ", id)
				rudsuc := h.Manager.DeleteClientRedirectURI(id)
				var rtn Response
				if rudsuc {
					rtn.Success = rudsuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
