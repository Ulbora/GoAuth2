package main

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

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	//"encoding/json"
	hd "github.com/Ulbora/GoAuth2/handlers"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"github.com/gorilla/mux"
)

//GO111MODULE=on go mod init github.com/Ulbora/GoAuth2
func main() {

	mock := flag.Bool("mock", false, "use mock backend")
	compressJwt := flag.Bool("compressJwt", false, "compress Jwt access token")
	assets := flag.String("assets", "", "Assets to control")
	flag.Parse()
	var logger lg.Logger
	logger.LogLevel = lg.InfoLevel
	logger.Info("mock: ", *mock)
	logger.Info("assets: ", *assets)
	logger.Info("compressJwt: ", *compressJwt)
	logger.LogLevel = lg.OffLevel
	var dbi db.Database
	var mydb mdb.MyDB
	var goauth2Host string
	var goauth2User string
	var goauth2Password string
	var goauth2Database string
	var authURL string

	if os.Getenv("GOAUTH2_HOST") != "" {
		goauth2Host = os.Getenv("GOAUTH2_HOST")
	} else {
		goauth2Host = "localhost:3306"
	}

	if os.Getenv("GOAUTH2_USER") != "" {
		goauth2User = os.Getenv("GOAUTH2_USER")
	} else {
		goauth2User = "admin"
	}

	if os.Getenv("GOAUTH2_PASSWORD") != "" {
		goauth2Password = os.Getenv("GOAUTH2_PASSWORD")
	} else {
		goauth2Password = "admin"
	}

	if os.Getenv("GOAUTH2_DATABASE") != "" {
		goauth2Database = os.Getenv("GOAUTH2_DATABASE")
	} else {
		goauth2Database = "go_auth2"
	}

	if os.Getenv("AUTHENTICATION_SERVICE") != "" {
		authURL = os.Getenv("AUTHENTICATION_SERVICE")
	} else {
		authURL = "http://localhost:3001/rs/user/login"
	}

	mydb.Host = goauth2Host
	mydb.User = goauth2User
	mydb.Password = goauth2Password
	mydb.Database = goauth2Database
	dbi = &mydb
	dbi.Connect()

	var wh hd.WebHandler
	var rh hd.RestHandler
	var owh *hd.OauthWebHandler

	if *mock {
		owh = hd.UseMockWeb()
		wh = owh.GetNewWebHandler()
		orh := hd.UseMockRest()
		rh = orh.GetNewRestHandler()
	} else {
		owh = hd.UseWebHandler(dbi, *compressJwt, authURL)
		wh = owh.GetNewWebHandler()
		orh := hd.UseRestHandler(dbi, *assets, *compressJwt, authURL, &logger)
		rh = orh.GetNewRestHandler()
	}

	owh.Templates = template.Must(template.ParseFiles("./static/head.html", "./static/index.html",
		"./static/login.html", "./static/authorizeApp.html", "./static/oauthError.html"))

	router := mux.NewRouter()
	port := "3000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	//web routes-------------------------------------
	router.HandleFunc("/", owh.Index).Methods("GET")
	router.HandleFunc("/login", wh.Login).Methods("GET")
	router.HandleFunc("/login", wh.LoginUser).Methods("POST")
	router.HandleFunc("/oauth/authorize", wh.Authorize).Methods("GET")
	router.HandleFunc("/authorizeApp", wh.AuthorizeApp).Methods("GET")
	router.HandleFunc("/applicationAuthorize", wh.ApplicationAuthorizationByUser).Methods("GET")
	router.HandleFunc("/oauth/token", wh.Token).Methods("POST")
	router.HandleFunc("/oauthError", wh.OauthError).Methods("GET")

	//REST routes--------------------------------------
	//Client
	router.HandleFunc("/rs/client/add", rh.AddClient).Methods("POST")
	router.HandleFunc("/rs/client/update", rh.UpdateClient).Methods("PUT")
	router.HandleFunc("/rs/client/get/{id}", rh.GetClient).Methods("GET")
	router.HandleFunc("/rs/client/admin/get", rh.GetClientAdmin).Methods("GET")
	router.HandleFunc("/rs/client/list", rh.GetClientList).Methods("GET")
	router.HandleFunc("/rs/client/search", rh.GetClientSearchList).Methods("POST")
	router.HandleFunc("/rs/client/delete/{id}", rh.DeleteClient).Methods("DELETE")

	//clientGrantType
	router.HandleFunc("/rs/clientGrantType/add", rh.AddGrantType).Methods("POST")
	router.HandleFunc("/rs/clientGrantType/list/{clientId}", rh.GetGrantTypeList).Methods("GET")
	router.HandleFunc("/rs/clientGrantType/delete/{id}", rh.DeleteGrantType).Methods("DELETE")

	//clientAllowedUri
	router.HandleFunc("/rs/clientAllowedUriSuper/add", rh.AddAllowedURISuper).Methods("POST")
	router.HandleFunc("/rs/clientAllowedUri/add", rh.AddAllowedURI).Methods("POST")
	router.HandleFunc("/rs/clientAllowedUriSuper/update", rh.UpdateAllowedURISuper).Methods("PUT")
	router.HandleFunc("/rs/clientAllowedUri/update", rh.UpdateAllowedURI).Methods("PUT")
	router.HandleFunc("/rs/clientAllowedUri/get/{id}", rh.GetAllowedURI).Methods("GET")
	router.HandleFunc("/rs/clientAllowedUri/list/{clientId}", rh.GetAllowedURIList).Methods("GET")
	router.HandleFunc("/rs/clientAllowedUri/delete/{id}", rh.DeleteAllowedURI).Methods("DELETE")

	//clientRedirectUri
	router.HandleFunc("/rs/clientRedirectUri/add", rh.AddRedirectURI).Methods("POST")
	router.HandleFunc("/rs/clientRedirectUri/list/{clientId}", rh.GetRedirectURIList).Methods("GET")
	router.HandleFunc("/rs/clientRedirectUri/delete/{id}", rh.DeleteRedirectURI).Methods("DELETE")

	//clientRole
	router.HandleFunc("/rs/clientRole/add", rh.AddRole).Methods("POST")
	router.HandleFunc("/rs/clientRoleSuper/add", rh.AddRoleSuper).Methods("POST")
	router.HandleFunc("/rs/clientRole/list/{clientId}", rh.GetRoleList).Methods("GET")
	router.HandleFunc("/rs/clientRole/delete/{id}", rh.DeleteRole).Methods("DELETE")

	//clientRoleUri
	router.HandleFunc("/rs/clientRoleUri/add", rh.AddRoleURI).Methods("POST")
	router.HandleFunc("/rs/clientRoleUri/list/{clientRoleId}", rh.GetRoleURIList).Methods("GET")
	//---- added post delete for backwards compatibility
	router.HandleFunc("/rs/clientRoleUri/delete", rh.DeleteRoleURI).Methods("POST")
	router.HandleFunc("/rs/clientRoleUri/delete/{clientRoleId}/{clientAllowedUriId}", rh.DeleteRoleURI).Methods("DELETE")

	//validate token
	router.HandleFunc("/rs/token/validate", rh.ValidateAccessToken).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Starting server Oauth2 Server on " + port)
	http.ListenAndServe(":"+port, router)

}

// [
//    {
//       "url":"/ulbora/rs/clientAllowedUri/add",
//       "assets":[
//          {
//             "controlledAsset":"ulbora",
//             "allowedRole":"superAdmin"
//          }
//       ]
//    },
//    {
//       "url":"/ulbora/rs/clientAllowedUri/update",
//       "assets":[
//          {
//             "controlledAsset":"ulbora",
//             "allowedRole":"superAdmin"
//          }
//       ]
//    },
//    {
//       "url":"/ulbora/rs/clientRole/add",
//       "assets":[
//          {
//             "controlledAsset":"superAdmin",
//             "allowedRole":"superAdmin"
//          }
//       ]
//    }
// ]

// WwogICB7CiAgICAgICJ1cmwiOiIvdWxib3JhL3JzL2NsaWVudEFsbG93ZWRVcmkvYWRkIiwKICAgICAgImFzc2V0cyI6WwogICAgICAgICB7CiAgICAgICAgICAgICJjb250cm9sbGVkQXNzZXQiOiJ1bGJvcmEiLAogICAgICAgICAgICAiYWxsb3dlZFJvbGUiOiJzdXBlckFkbWluIgogICAgICAgICB9CiAgICAgIF0KICAgfSwKICAgewogICAgICAidXJsIjoiL3VsYm9yYS9ycy9jbGllbnRBbGxvd2VkVXJpL3VwZGF0ZSIsCiAgICAgICJhc3NldHMiOlsKICAgICAgICAgewogICAgICAgICAgICAiY29udHJvbGxlZEFzc2V0IjoidWxib3JhIiwKICAgICAgICAgICAgImFsbG93ZWRSb2xlIjoic3VwZXJBZG1pbiIKICAgICAgICAgfQogICAgICBdCiAgIH0sCiAgIHsKICAgICAgInVybCI6Ii91bGJvcmEvcnMvY2xpZW50Um9sZS9hZGQiLAogICAgICAiYXNzZXRzIjpbCiAgICAgICAgIHsKICAgICAgICAgICAgImNvbnRyb2xsZWRBc3NldCI6InN1cGVyQWRtaW4iLAogICAgICAgICAgICAiYWxsb3dlZFJvbGUiOiJzdXBlckFkbWluIgogICAgICAgICB9CiAgICAgIF0KICAgfQpd
