//Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/Ulbora/GoAuth2/managers"
	oc "github.com/Ulbora/GoAuth2/oauthclient"
	ac "github.com/Ulbora/GoAuth2/rolecontrol"
	lg "github.com/Ulbora/Level_Logger"
	"github.com/gorilla/mux"
)

// add url grant type

func TestOauthRestHandlerGrtType_AddGrantType(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = true
	asc.MockAllowedRole = "admin"
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "grantType":"code", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddGrantType(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy ResponseID
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !bdy.Success {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_AddGrantTypeBadMedia(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = true
	asc.MockAllowedRole = "admin"
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "grantType":"code", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddGrantType(w, r)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_AddGrantTypeAddFail(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = false
	//man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = true
	asc.MockAllowedRole = "admin"
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "grantType":"code", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddGrantType(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_AddGrantTypeNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = true
	asc.MockAllowedRole = "admin"
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "grantType":"code", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddGrantType(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_AddGrantTypeBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = true
	asc.MockAllowedRole = "admin"
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddGrantType(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

// // get list

func TestOauthRestHandlerGrtType_GetGrantTypelist(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientGrantType
	cuo.ID = 4
	cuo.GrantType = "code"
	cuo.ClientID = 10

	var cuol = []m.ClientGrantType{cuo}

	var man m.MockManager
	man.MockClientGrantTypeList = cuol
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"clientId": "55",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetGrantTypeList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || len(bdy) != 1 {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_GetGrantTypeListNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientGrantType
	cuo.ID = 4
	cuo.GrantType = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	//man.MockClientAllowedURI = cuo
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetGrantTypeList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_GetGrantTypeListBadParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientGrantType
	cuo.ID = 4
	cuo.GrantType = "code"
	cuo.ClientID = 10

	var man m.MockManager
	//man.MockClientAllowedURI = cuo
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "q",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetGrantTypeList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_GetGrantTypeListNoParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientGrantType
	cuo.ID = 4
	cuo.GrantType = "code"
	cuo.ClientID = 10

	var man m.MockManager
	//man.MockClientAllowedURI = cuo
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	// vars := map[string]string{
	// 	"id": "q",
	// }
	// r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetGrantTypeList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientGrantType
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

// // delete gt

func TestOauthRestHandlerGrtType_DeleteGrantType(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockDeleteSuccess1 = true
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteGrantType(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Success != true {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_DeleteGrantTypeFail(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockDeleteSuccess1 = false
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteGrantType(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 500 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_DeleteGrantTypeNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockDeleteSuccess1 = true
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteGrantType(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_DeleteGrantTypeBadParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockDeleteSuccess1 = true
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "q",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteGrantType(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandlerGrtType_DeleteGrantTypeNoParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockDeleteSuccess1 = true
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	// vars := map[string]string{
	// 	"id": "q",
	// }
	// r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteGrantType(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}
