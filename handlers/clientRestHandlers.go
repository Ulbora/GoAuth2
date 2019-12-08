//Package handlers ...
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

//AddClient AddClient
func (h *OauthRestHandler) AddClient(w http.ResponseWriter, r *http.Request) {
	//url of this endpoint
	var addCltURL = "/ulbora/rs/client/add"

	var acltcl oc.Claim
	acltcl.Role = "superAdmin"
	acltcl.URL = addCltURL
	acltcl.Scope = "write"
	fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &acltcl)
	if auth {
		h.SetContentType(w)
		acltContOk := h.CheckContent(r)
		fmt.Println("conOk: ", acltContOk)
		if !acltContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cus m.Client
			bcsuc, bcerr := h.ProcessBody(r, &cus)
			fmt.Println("bcsuc: ", bcsuc)
			fmt.Println("client in add: ", cus)
			fmt.Println("bcerr: ", bcerr)
			if !bcsuc && bcerr != nil {
				http.Error(w, bcerr.Error(), http.StatusBadRequest)
			} else {
				// if cus.RedirectURIs == nil {
				// 	var rui []m.ClientRedirectURI
				// 	cus.RedirectURIs = &rui
				// }
				// //test only
				// for _, u := range *cus.RedirectURIs {
				// 	fmt.Println("RedirectURIs in client: ", u)
				// }

				acltSuc, acltID := h.Manager.AddClient(&cus)
				fmt.Println("acltSuc: ", acltSuc)
				fmt.Println("acltID: ", acltID)
				var rtn ResponseID
				if acltSuc && acltID != 0 {
					rtn.Success = acltSuc
					rtn.ID = acltID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var facltrtn ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(facltrtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateClient UpdateClient
func (h *OauthRestHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	//url of this endpoint
	var upcltURL = "/ulbora/rs/client/update"

	var upcltscl oc.Claim
	upcltscl.Role = "superAdmin"
	upcltscl.URL = upcltURL
	upcltscl.Scope = "write"
	fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &upcltscl)
	if auth {
		// w.Header().Set("Content-Type", "application/json")
		h.SetContentType(w)
		uPcltContOk := h.CheckContent(r)
		fmt.Println("conOk: ", uPcltContOk)
		if !uPcltContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var upclt m.Client
			ubcltsuc, uberr := h.ProcessBody(r, &upclt)
			fmt.Println("ubcltsuc: ", ubcltsuc)
			fmt.Println("client in update: ", upclt)
			fmt.Println("uberr: ", uberr)
			if !ubcltsuc && uberr != nil {
				http.Error(w, uberr.Error(), http.StatusBadRequest)
			} else {
				uPcltSuc := h.Manager.UpdateClient(&upclt)
				fmt.Println("uPcltSuc: ", uPcltSuc)
				var rtn Response
				if uPcltSuc {
					rtn.Success = uPcltSuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var fucltrtn Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(fucltrtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetClient GetClient
func (h *OauthRestHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	var getCltURL = "/ulbora/rs/client/get"

	var gcltcl oc.Claim
	gcltcl.Role = "superAdmin"
	gcltcl.URL = getCltURL
	gcltcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gcltcl)
	if auth {
		//var id string
		h.SetContentType(w)
		gcltvars := mux.Vars(r)
		fmt.Println("vars: ", len(gcltvars))
		if gcltvars != nil && len(gcltvars) != 0 {
			var cltidStr = gcltvars["id"]
			fmt.Println("vars: ", gcltvars)
			cltid, idErr := strconv.ParseInt(cltidStr, 10, 64)
			if cltid != 0 && idErr == nil {
				fmt.Println("id: ", cltid)
				getClt := h.Manager.GetClient(cltid)
				fmt.Println("getClt: ", getClt)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getClt)
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

//GetClientAdmin GetClientAdmin
func (h *OauthRestHandler) GetClientAdmin(w http.ResponseWriter, r *http.Request) {
	var getCltaURL = "/ulbora/rs/client/admin/get"

	var gcltacl oc.Claim
	gcltacl.Role = "admin"
	gcltacl.URL = getCltaURL
	gcltacl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gcltacl)
	if auth {
		//var id string
		h.SetContentType(w)
		//gcltavars := mux.Vars(r)
		//fmt.Println("vars: ", len(gcltavars))
		cltaidStr := r.Header.Get("clientId")
		if cltaidStr != "" {
			//var cltaidStr = gcltavars["id"]
			//fmt.Println("vars: ", gcltavars)
			cltaid, idErr := strconv.ParseInt(cltaidStr, 10, 64)
			if cltaid != 0 && idErr == nil {
				fmt.Println("clientId in admin get: ", cltaid)
				getClta := h.Manager.GetClient(cltaid)
				fmt.Println("getClt: ", getClta)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getClta)
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

//GetClientList GetClientList
func (h *OauthRestHandler) GetClientList(w http.ResponseWriter, r *http.Request) {
	var getcltlURL = "/ulbora/rs/client/list"

	var gcltlcl oc.Claim
	gcltlcl.Role = "superAdmin"
	gcltlcl.URL = getcltlURL
	gcltlcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gcltlcl)
	if auth {
		h.SetContentType(w)
		getcltl := h.Manager.GetClientList()
		fmt.Println("getcltl: ", getcltl)
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(getcltl)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetClientSearchList GetClientSearchList
func (h *OauthRestHandler) GetClientSearchList(w http.ResponseWriter, r *http.Request) {
	var getCltsURL = "/ulbora/rs/client/search"

	var gcltscl oc.Claim
	gcltscl.Role = "admin"
	gcltscl.URL = getCltsURL
	gcltscl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gcltscl)
	if auth {
		h.SetContentType(w)
		acltsContOk := h.CheckContent(r)
		fmt.Println("conOk: ", acltsContOk)
		if !acltsContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cus m.Client
			bcltssuc, bcltserr := h.ProcessBody(r, &cus)
			fmt.Println("bcltssuc: ", bcltssuc)
			fmt.Println("client in search: ", cus)
			fmt.Println("bcltserr: ", bcltserr)
			if !bcltssuc && bcltserr != nil {
				http.Error(w, bcltserr.Error(), http.StatusBadRequest)
			} else {
				if cus.Name != "" {
					getCltsl := h.Manager.GetClientSearchList(cus.Name)
					fmt.Println("getCltsl: ", getCltsl)
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(getCltsl)
					fmt.Fprint(w, string(resJSON))
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteClient DeleteClient
func (h *OauthRestHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	var cltdURL = "/ulbora/rs/client/delete"

	var cltdcl oc.Claim
	cltdcl.Role = "superAdmin"
	cltdcl.URL = cltdURL
	cltdcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &cltdcl)
	if auth {
		h.SetContentType(w)
		cltdvars := mux.Vars(r)
		fmt.Println("vars: ", len(cltdvars))
		if cltdvars != nil && len(cltdvars) != 0 {
			var cltdidStr = cltdvars["id"]
			fmt.Println("vars delete: ", cltdvars)
			cltdid, idErr := strconv.ParseInt(cltdidStr, 10, 64)
			fmt.Println("id delete clientId: ", cltdid)
			if cltdid != 0 && idErr == nil {
				fmt.Println("id: ", cltdid)
				cltdSuc := h.Manager.DeleteClient(cltdid)
				var cltdrtn Response
				if cltdSuc {
					cltdrtn.Success = cltdSuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(cltdrtn)
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
