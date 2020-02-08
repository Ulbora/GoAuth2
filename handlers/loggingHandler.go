//Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	lg "github.com/Ulbora/Level_Logger"
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

//LogLevel LogLevel
type LogLevel struct {
	Level string `json:"logLevel"`
}

//LogResponse LogResponse
type LogResponse struct {
	Success  bool   `json:"success"`
	LogLevel string `json:"logLevel"`
}

const (
	defaultLoggingKey = "45sdbb2345"

	debugLevel = "DEBUG"
	infoLevel  = "INFO"
	allLevel   = "ALL"
	offLevel   = "OFF"
)

//SetLogLevel SetLogLevel
func (h *OauthRestHandler) SetLogLevel(w http.ResponseWriter, r *http.Request) {
	var logRes LogResponse
	h.SetContentType(w)
	logContOk := h.CheckContent(r)

	//fmt.Println("conOk: ", logContOk)

	if !logContOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var loggingKey string
		if os.Getenv("Logging_KEY") != "" {
			loggingKey = os.Getenv("Logging_KEY")
		} else {
			loggingKey = defaultLoggingKey
		}
		loggingKeyHdr := r.Header.Get("Logging_KEY")
		if loggingKey == loggingKeyHdr {
			var lv LogLevel
			lgsuc, lgerr := h.ProcessBody(r, &lv)
			//fmt.Println("lgsuc: ", lgsuc)
			//fmt.Println("LogLevel: ", lv)
			//fmt.Println("lgerr: ", lgerr)
			if !lgsuc && lgerr != nil {
				http.Error(w, lgerr.Error(), http.StatusBadRequest)
			} else {
				switch strings.ToUpper(lv.Level) {
				case debugLevel:
					h.Log.LogLevel = lg.DebugLevel
					logRes.Success = true
					logRes.LogLevel = debugLevel
				case infoLevel:
					h.Log.LogLevel = lg.InfoLevel
					logRes.Success = true
					logRes.LogLevel = infoLevel
				case allLevel:
					h.Log.LogLevel = lg.AllLevel
					logRes.Success = true
					logRes.LogLevel = allLevel
				case offLevel:
					h.Log.LogLevel = lg.OffLevel
					logRes.Success = true
					logRes.LogLevel = offLevel
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		resJSON, _ := json.Marshal(logRes)
		fmt.Fprint(w, string(resJSON))
	}
}
