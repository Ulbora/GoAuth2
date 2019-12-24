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

//AddRoleURI AddRoleURI
func (h *OauthRestHandler) AddRoleURI(w http.ResponseWriter, r *http.Request) {
	var crui m.ClientRoleURI
	ruibsuc, gerr := h.ProcessBody(r, &crui)
	fmt.Println("ruisuc: ", ruibsuc)
	fmt.Println("crui: ", crui)
	fmt.Println("gerr: ", gerr)
	if ruibsuc && gerr == nil {

		//url of this endpoint
		var addRoleU = "/ulbora/rs/clientRoleUri/add"

		var rlucl oc.Claim
		rlucl.Role = "admin"

		rlucl.URL = addRoleU
		rlucl.Scope = "write"
		fmt.Println("client: ", h.Client)

		//check that jwt token user role has permission to use the url of this endpoint
		auth := h.Client.Authorize(r, &rlucl)

		if auth {
			// w.Header().Set("Content-Type", "application/json")
			h.SetContentType(w)
			rluContOk := h.CheckContent(r)
			fmt.Println("conOk: ", rluContOk)
			if !rluContOk {
				http.Error(w, "json required", http.StatusUnsupportedMediaType)
			} else {
				rluiaSuc := h.Manager.AddClientRoleURI(&crui)
				fmt.Println("gtaSuc: ", rluiaSuc)
				//fmt.Println("gtID: ", gtID)
				var rtn Response
				if rluiaSuc {
					rtn.Success = rluiaSuc
					// rtn.ID = gtID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		} else {
			var rluiartn ResponseID
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(rluiartn)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Error(w, gerr.Error(), http.StatusBadRequest)
	}
}

//GetRoleURIList GetRoleURIList
func (h *OauthRestHandler) GetRoleURIList(w http.ResponseWriter, r *http.Request) {
	var ruialURL = "/ulbora/rs/clientRoleUri/list"

	var ruilcl oc.Claim
	ruilcl.Role = "admin"
	ruilcl.URL = ruialURL
	ruilcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &ruilcl)
	if auth {
		//var id string
		h.SetContentType(w)
		ruilvars := mux.Vars(r)
		fmt.Println("vars: ", len(ruilvars))
		if ruilvars != nil && len(ruilvars) != 0 {
			var ruiRoleIDStr = ruilvars["clientRoleId"]
			fmt.Println("vars: ", ruilvars)
			ruiRoleID, ruiidErr := strconv.ParseInt(ruiRoleIDStr, 10, 64)
			if ruiRoleID != 0 && ruiidErr == nil {
				fmt.Println("roleId: ", ruiRoleID)
				getRuil := h.Manager.GetClientRoleAllowedURIList(ruiRoleID)
				fmt.Println("getgtl: ", getRuil)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getRuil)
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

//DeleteRoleURI DeleteRoleURI
func (h *OauthRestHandler) DeleteRoleURI(w http.ResponseWriter, r *http.Request) {
	var droleURI = "/ulbora/rs/clientRoleUri/delete"

	var ruidcl oc.Claim
	ruidcl.Role = "admin"
	ruidcl.URL = droleURI
	ruidcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &ruidcl)
	if auth {
		h.SetContentType(w)
		ruidvars := mux.Vars(r)
		fmt.Println("vars del ru: ", len(ruidvars))
		var clientRoleID int64
		var clientAllowedURIID int64

		if ruidvars != nil && len(ruidvars) == 2 {
			var cridStr = ruidvars["clientRoleId"]
			fmt.Println("vars delete crid: ", cridStr)
			crid, cridErr := strconv.ParseInt(cridStr, 10, 64)
			if crid != 0 && cridErr == nil {
				clientRoleID = crid
			}
			var cauidStr = ruidvars["clientAllowedUriId"]
			fmt.Println("vars delete cauidStr: ", cauidStr)
			cauid, cauidErr := strconv.ParseInt(cauidStr, 10, 64)
			if cauid != 0 && cauidErr == nil {
				clientAllowedURIID = cauid
			}
		} else {
			rlsContOk := h.CheckContent(r)
			fmt.Println("conOk: ", rlsContOk)
			if !rlsContOk {
				http.Error(w, "json required", http.StatusUnsupportedMediaType)
			} else {
				var crui m.ClientRoleURI
				rlssuc, rlserr := h.ProcessBody(r, &crui)
				fmt.Println("rlssuc: ", rlssuc)
				fmt.Println("crui: ", crui)
				fmt.Println("rlserr: ", rlserr)
				if !rlssuc && rlserr != nil {
					http.Error(w, rlserr.Error(), http.StatusBadRequest)
				} else {
					clientRoleID = crui.ClientRoleID
					clientAllowedURIID = crui.ClientAllowedURIID
				}
			}
		}
		if clientRoleID != 0 && clientAllowedURIID != 0 {
			fmt.Println("clientRoleID: ", clientRoleID)
			var crui m.ClientRoleURI
			crui.ClientRoleID = clientRoleID
			crui.ClientAllowedURIID = clientAllowedURIID
			cruidsuc := h.Manager.DeleteClientRoleURI(&crui)
			var crudrtn Response
			if cruidsuc {
				crudrtn.Success = cruidsuc
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			resJSON, _ := json.Marshal(crudrtn)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
