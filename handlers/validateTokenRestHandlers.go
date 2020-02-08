//Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

//ValidationResponse ValidationResponse
type ValidationResponse struct {
	Valid bool `json:"valid"`
}

//ValidateAccessToken ValidateAccessToken
func (h *OauthRestHandler) ValidateAccessToken(w http.ResponseWriter, r *http.Request) {
	var vdtr m.ValidateAccessTokenReq
	vdtsuc, gerr := h.ProcessBody(r, &vdtr)
	if h.TokenCompressed {
		vdtr.AccessToken = h.JwtCompress.UnCompressJwt(vdtr.AccessToken)
		fmt.Println("uncompressed token in validate: ", vdtr.AccessToken)
	}
	h.Log.Debug("vdtsuc: ", vdtsuc)
	h.Log.Debug("vdtr: ", vdtr)
	h.Log.Debug("gerr: ", gerr)
	if vdtsuc && gerr == nil {
		h.SetContentType(w)
		vdtContOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", vdtContOk)
		if !vdtContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			vdtSuc := h.Manager.ValidateAccessToken(&vdtr)
			h.Log.Debug("vdtSuc: ", vdtSuc)
			var rtn ValidationResponse
			if vdtSuc {
				rtn.Valid = vdtSuc
			}
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(rtn)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Error(w, gerr.Error(), http.StatusBadRequest)
	}
}
