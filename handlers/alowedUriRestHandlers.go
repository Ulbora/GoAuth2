package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/Ulbora/GoAuth2/managers"
	oc "github.com/Ulbora/GoAuth2/oauthclient"
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

//AddAllowedURISuper AddAllowedURISuper
func (h *OauthRestHandler) AddAllowedURISuper(w http.ResponseWriter, r *http.Request) {
	//var theURL = "/ulbora/rs/clientAllowedUriSuper/add"
	//acsuc, role := h.AssetControl.GetControlledAsset(theURL, "ulbora")
	var cl oc.Claim
	cl.Role = "testRole"
	cl.URL = "testURL"
	cl.Scope = "web"
	fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &cl)
	if auth {
		w.Header().Set("Content-Type", "application/json")
		aaURIContOk := h.CheckContent(w, r)
		fmt.Println("conOk: ", aaURIContOk)
		if !aaURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cu m.ClientAllowedURI
			bsuc, berr := h.ProcessBody(r, &cu)
			fmt.Println("bsuc: ", bsuc)
			fmt.Println("cu: ", cu)
			fmt.Println("berr: ", berr)
			if !bsuc && berr != nil {
				http.Error(w, berr.Error(), http.StatusBadRequest)
			} else {
				auSuc, auID := h.Manager.AddClientAllowedURI(&cu)
				fmt.Println("auSuc: ", auSuc)
				fmt.Println("auID: ", auID)
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
		}
	} else {
		var frtn ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(frtn)
		fmt.Fprint(w, string(resJSON))
	}

}
