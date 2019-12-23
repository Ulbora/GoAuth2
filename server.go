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
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	//"encoding/json"
	hd "github.com/Ulbora/GoAuth2/handlers"
	m "github.com/Ulbora/GoAuth2/managers"
	oa "github.com/Ulbora/GoAuth2/oauthclient"
	rc "github.com/Ulbora/GoAuth2/rolecontrol"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"github.com/gorilla/mux"
)

//GO111MODULE=on go mod init github.com/Ulbora/GoAuth2
func main() {
	var dbi db.Database
	var mydb mdb.MyDBMock
	dbi = &mydb
	dbi.Connect()

	var wh hd.WebHandler
	owh := useMockWeb()
	wh = owh.GetNewWebHandler()

	var rh hd.RestHandler
	orh := useMockRest()
	rh = orh.GetNewRestHandler()

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

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Starting server Oauth2 Server on " + port)
	http.ListenAndServe(":"+port, router)

}

func useMockWeb() *hd.OauthWebHandler {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockImplicitAuthorized = true
	var acc m.AuthCodeClient
	//acc.Valid = true
	acc.ClientName = "Test Client"
	acc.WebSite = "www.TestClient.com"
	om.MockAuthCodeClient = acc

	var ic m.ImplicitClient
	ic.Valid = true
	ic.ClientName = "Test Client"
	ic.WebSite = "www.TestClient.com"
	om.MockImplicitClient = ic
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockUserLoginSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	om.MockImplicitAuthorizeSuccess = true
	var ir m.ImplicitReturn
	ir.ID = 3
	ir.Token = "12345"
	om.MockImplicitReturn = ir

	var tkn m.Token
	tkn.AccessToken = "65165165"
	tkn.TokenType = "Bearer"
	tkn.RefreshToken = "16161"
	tkn.ExpiresIn = 50000

	om.MockAuthCodeTokenSuccess = true
	om.MockCredentialsTokenSuccess = true
	om.MockAuthCodeRefreshTokenSuccess = true
	om.MockPasswordTokenSuccess = true
	om.MockToken = tkn
	om.MockTokenError = "token failed"

	var wh hd.OauthWebHandler
	wh.Manager = &om
	return &wh
}

func useMockRest() *hd.OauthRestHandler {
	var om m.MockManager
	om.MockInsertSuccess1 = true
	om.MockInsertID1 = 34
	om.MockUpdateSuccess1 = true
	om.MockDeleteSuccess1 = true

	var mc m.Client
	mc.ClientID = 510
	mc.Secret = "12345"
	mc.Name = "test client"
	mc.WebSite = "www.testclient.com"
	mc.Email = "tester@testclient.com"
	mc.Enabled = true

	var cuo m.ClientRedirectURI
	cuo.ID = 4
	cuo.URI = "/test"
	cuo.ClientID = 10
	mc.RedirectURIs = &[]m.ClientRedirectURI{cuo}
	om.MockClient = mc

	om.MockClientList = []m.Client{mc}

	var gt m.ClientGrantType
	gt.ID = 2
	gt.GrantType = "code"
	gt.ClientID = 22

	om.MockClientGrantTypeList = []m.ClientGrantType{gt}

	var au m.ClientAllowedURI
	au.ID = 5
	au.URI = "/testurl"
	au.ClientID = 4

	om.MockClientAllowedURI = au
	om.MockClientAllowedURIList = []m.ClientAllowedURI{au}

	var ru m.ClientRedirectURI
	ru.ID = 4
	ru.URI = "/testuri"
	ru.ClientID = 554

	om.MockClientRedirectURIList = []m.ClientRedirectURI{ru}

	var rh hd.OauthRestHandler
	rh.Manager = &om

	var clt oa.MockOauthClient
	clt.MockValid = true
	rh.Client = &clt

	var ac rc.MockOauthAssets
	ac.MockSuccess = true
	rh.AssetControl = &ac
	return &rh
}
