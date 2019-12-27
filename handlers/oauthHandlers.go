//Package handlers ...
package handlers

import (
	"encoding/json"
	"errors"

	"log"
	"net/http"

	cp "github.com/Ulbora/GoAuth2/compresstoken"
	m "github.com/Ulbora/GoAuth2/managers"
	oa "github.com/Ulbora/GoAuth2/oauthclient"
	rc "github.com/Ulbora/GoAuth2/rolecontrol"
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

//OauthRestHandler OauthRestHandler
type OauthRestHandler struct {
	Manager         m.Manager
	Client          oa.Client
	AssetControl    rc.AssetControl
	TokenCompressed bool
	JwtCompress     cp.JwtCompress
}

//GetNewRestHandler GetNewRestHandler
func (h *OauthRestHandler) GetNewRestHandler() RestHandler {
	var rh RestHandler
	rh = h
	return rh
}

//CheckContent CheckContent
func (h *OauthRestHandler) CheckContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		// http.Error(w, "json required", http.StatusUnsupportedMediaType)
		rtn = true
	}
	return rtn
}

//SetContentType SetContentType
func (h *OauthRestHandler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//ProcessBody ProcessBody
func (h *OauthRestHandler) ProcessBody(r *http.Request, obj interface{}) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			log.Println("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}

	return suc, err
}
