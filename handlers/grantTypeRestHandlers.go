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

//AddGrantType AddGrantType
func (h *OauthRestHandler) AddGrantType(w http.ResponseWriter, r *http.Request) {
	var cg m.ClientGrantType
	gsuc, gerr := h.ProcessBody(r, &cg)
	fmt.Println("gsuc: ", gsuc)
	fmt.Println("cg: ", cg)
	fmt.Println("gerr: ", gerr)
	if gsuc && gerr == nil {

		//url of this endpoint
		var addAuURL = "/ulbora/rs/clientGrantType/add"

		var gtcl oc.Claim
		//if acsuc {
		//aucl.Role = role
		//} else {
		gtcl.Role = "admin"
		//}
		gtcl.URL = addAuURL
		gtcl.Scope = "write"
		fmt.Println("client: ", h.Client)

		//check that jwt token user role has permission to use the url of this endpoint
		auth := h.Client.Authorize(r, &gtcl)

		if auth {
			// w.Header().Set("Content-Type", "application/json")
			h.SetContentType(w)
			aaURIContOk := h.CheckContent(r)
			fmt.Println("conOk: ", aaURIContOk)
			if !aaURIContOk {
				http.Error(w, "json required", http.StatusUnsupportedMediaType)
			} else {
				auSuc, auID := h.Manager.AddClientGrantType(&cg)
				fmt.Println("gtSuc: ", auSuc)
				fmt.Println("gtID: ", auID)
				var rtn ResponseID
				if auSuc && auID != 0 {
					rtn.Success = auSuc
					rtn.ID = auID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		} else {
			var frtn ResponseID
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(frtn)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Error(w, gerr.Error(), http.StatusBadRequest)
	}
}

//GetGrantTypeList GetGrantTypeList
func (h *OauthRestHandler) GetGrantTypeList(w http.ResponseWriter, r *http.Request) {
	var getAulURL = "/ulbora/rs/clientGrantType/list"

	var gtlcl oc.Claim
	gtlcl.Role = "admin"
	gtlcl.URL = getAulURL
	gtlcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gtlcl)
	if auth {
		//var id string
		h.SetContentType(w)
		gtlvars := mux.Vars(r)
		fmt.Println("vars: ", len(gtlvars))
		if gtlvars != nil && len(gtlvars) != 0 {
			var gtclientIDStr = gtlvars["clientId"]
			fmt.Println("vars: ", gtlvars)
			clientID, gtlidErr := strconv.ParseInt(gtclientIDStr, 10, 64)
			if clientID != 0 && gtlidErr == nil {
				fmt.Println("clientID: ", clientID)
				getgtl := h.Manager.GetClientGrantTypeList(clientID)
				fmt.Println("getgtl: ", getgtl)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getgtl)
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

//DeleteGrantType DeleteGrantType
func (h *OauthRestHandler) DeleteGrantType(w http.ResponseWriter, r *http.Request) {
	var getAudURL = "/ulbora/rs/clientGrantType/delete"

	var gtdcl oc.Claim
	gtdcl.Role = "admin"
	gtdcl.URL = getAudURL
	gtdcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gtdcl)
	if auth {
		//var id string
		h.SetContentType(w)
		gtdvars := mux.Vars(r)
		fmt.Println("vars: ", len(gtdvars))
		if gtdvars != nil && len(gtdvars) != 0 {
			var gtlidStr = gtdvars["id"]
			fmt.Println("vars delete: ", gtdvars)
			id, idErr := strconv.ParseInt(gtlidStr, 10, 64)
			fmt.Println("id delete: ", id)
			if id != 0 && idErr == nil {
				fmt.Println("id: ", id)
				gtdsuc := h.Manager.DeleteClientGrantType(id)
				var rtn Response
				if gtdsuc {
					rtn.Success = gtdsuc
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
