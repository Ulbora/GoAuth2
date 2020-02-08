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

//AddRoleSuper AddRoleSuper
func (h *OauthRestHandler) AddRoleSuper(w http.ResponseWriter, r *http.Request) {
	//url of this endpoint
	var addRlsURL = "/ulbora/rs/clientRoleSuper/add"

	var rlscl oc.Claim
	rlscl.Role = "superAdmin"
	rlscl.URL = addRlsURL
	rlscl.Scope = "write"
	h.Log.Debug("client: ", h.Client)
	auth := h.Client.Authorize(r, &rlscl)
	if auth {
		h.SetContentType(w)
		rlsContOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", rlsContOk)
		if !rlsContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cus m.ClientRole
			rlssuc, rlserr := h.ProcessBody(r, &cus)
			h.Log.Debug("rlssuc: ", rlssuc)
			h.Log.Debug("cu: ", cus)
			h.Log.Debug("rlserr: ", rlserr)
			if !rlssuc && rlserr != nil {
				http.Error(w, rlserr.Error(), http.StatusBadRequest)
			} else {
				arlsSuc, arlsID := h.Manager.AddClientRole(&cus)
				h.Log.Debug("arlsSuc: ", arlsSuc)
				h.Log.Debug("arlsID: ", arlsID)
				var rtn ResponseID
				if arlsSuc && arlsID != 0 {
					rtn.Success = arlsSuc
					rtn.ID = arlsID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var arlsrtn ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(arlsrtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//AddRole AddRole
func (h *OauthRestHandler) AddRole(w http.ResponseWriter, r *http.Request) {
	var cu m.ClientRole
	rlbsuc, rlberr := h.ProcessBody(r, &cu)
	h.Log.Debug("rlbsuc: ", rlbsuc)
	h.Log.Debug("cu: ", cu)
	h.Log.Debug("rlberr: ", rlberr)
	if rlbsuc && rlberr == nil {

		//url of this endpoint
		var addrl = "/ulbora/rs/clientRole/add"

		//make sure the user in not trying to add a prohibited url that has "ulbora" in the url
		//looks through a list of assets for the url and determines the role needed based on the asset part of the url
		arlacsuc, role := h.AssetControl.GetControlledAsset(addrl, "superAdmin")

		var rlcl oc.Claim
		if arlacsuc {
			rlcl.Role = role
		} else {
			rlcl.Role = "admin"
		}
		rlcl.URL = addrl
		rlcl.Scope = "write"
		h.Log.Debug("client: ", h.Client)

		//check that jwt token user role has permission to use the url of this endpoint
		auth := h.Client.Authorize(r, &rlcl)

		if auth {
			// w.Header().Set("Content-Type", "application/json")
			h.SetContentType(w)
			arlContOk := h.CheckContent(r)
			h.Log.Debug("conOk: ", arlContOk)
			if !arlContOk {
				http.Error(w, "json required", http.StatusUnsupportedMediaType)
			} else {
				arlSuc, arlID := h.Manager.AddClientRole(&cu)
				h.Log.Debug("arlSuc: ", arlSuc)
				h.Log.Debug("arlID: ", arlID)
				var rtn ResponseID
				if arlSuc && arlID != 0 {
					rtn.Success = arlSuc
					rtn.ID = arlID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		} else {
			var arlrtn ResponseID
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(arlrtn)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Error(w, rlberr.Error(), http.StatusBadRequest)
	}
}

//GetRoleList GetRoleList
func (h *OauthRestHandler) GetRoleList(w http.ResponseWriter, r *http.Request) {
	var getRllURL = "/ulbora/rs/clientRole/list"

	var grllcl oc.Claim
	grllcl.Role = "admin"
	grllcl.URL = getRllURL
	grllcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &grllcl)
	if auth {
		//var id string
		h.SetContentType(w)
		rllvars := mux.Vars(r)
		h.Log.Debug("vars: ", len(rllvars))
		if rllvars != nil && len(rllvars) != 0 {
			var rllClientIDStr = rllvars["clientId"]
			h.Log.Debug("vars: ", rllvars)
			rllclientID, idErr := strconv.ParseInt(rllClientIDStr, 10, 64)
			if rllclientID != 0 && idErr == nil {
				h.Log.Debug("clientID: ", rllclientID)
				getRll := h.Manager.GetClientRoleList(rllclientID)
				h.Log.Debug("getAul: ", getRll)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getRll)
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

//DeleteRole DeleteRole
func (h *OauthRestHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	var rldURL = "/ulbora/rs/clientRole/delete"

	var rldcl oc.Claim
	rldcl.Role = "admin"
	rldcl.URL = rldURL
	rldcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &rldcl)
	if auth {
		//var id string
		h.SetContentType(w)
		rldvars := mux.Vars(r)
		h.Log.Debug("vars: ", len(rldvars))
		if rldvars != nil && len(rldvars) != 0 {
			var rldidStr = rldvars["id"]
			h.Log.Debug("vars delete: ", rldvars)
			rldid, idErr := strconv.ParseInt(rldidStr, 10, 64)
			h.Log.Debug("id delete: ", rldid)
			if rldid != 0 && idErr == nil {
				h.Log.Debug("id: ", rldid)
				rlsucd := h.Manager.DeleteClientRole(rldid)
				var rlrtn Response
				if rlsucd {
					rlrtn.Success = rlsucd
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rlrtn)
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
