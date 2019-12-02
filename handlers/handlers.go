package handlers

import "net/http"

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

//ResponseID ResponseID
type ResponseID struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
}

//Response Response
type Response struct {
	Success bool `json:"success"`
}

//RestHandler RestHandler
type RestHandler interface {
	AddAllowedURISuper(w http.ResponseWriter, r *http.Request)
	AddAllowedURI(w http.ResponseWriter, r *http.Request)
	UpdateAllowedURISuper(w http.ResponseWriter, r *http.Request)
	UpdateAllowedURI(w http.ResponseWriter, r *http.Request)
	GetAllowedURI(w http.ResponseWriter, r *http.Request)
	GetAllowedURIList(w http.ResponseWriter, r *http.Request)
	// start her with getList
}
