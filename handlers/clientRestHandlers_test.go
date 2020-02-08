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

func TestOauthRestHandler_AddClientBadMedia(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	// r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddClient(w, r)
	if w.Code != 415 {
		t.Fail()
	}

}

func TestOauthRestHandler_AddClientBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	//var robj m.Client
	// robj.ClientID = 3
	// robj.Secret = "testsecret"
	// robj.Name = "testname"
	//aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddClient(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddClient(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"testsecret", "name":"testname", "webSite": "testwebsite", "email": "testemail", "enabled": true, "paid": false, "redirectUrls": [{"id":3, "uri":"/test", "clientId": 2}]}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddClient(w, r)
	hd := w.Header()
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_AddClientFailAdd(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = false
	//man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"testsecret", "name":"testname", "webSite": "testwebsite", "email": "testemail", "enabled": true, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddClient(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddClientNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"testsecret", "name":"testname", "webSite": "testwebsite", "email": "testemail", "enabled": true, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddClient(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

// update uri super
func TestOauthRestHandler_UpdateClientBadMedia(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	// r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateClient(w, r)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateClientBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	// var robj m.ClientAllowedURI
	// robj.ID = 3
	// robj.URI = "/test"
	// robj.ClientID = 1
	//aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateClient(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateClient(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockUpdateSuccess1 = true
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"clientId": 5, "secret":"testsecret", "name":"testname", "webSite": "testwebsite", "email": "testemail", "enabled": true, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateClient(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateClientFailAdd(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = false
	//man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"clientId": 5, "secret":"testsecret", "name":"testname", "webSite": "testwebsite", "email": "testemail", "enabled": true, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateClient(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateClientNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"clientId": 5, "secret":"testsecret", "name":"testname", "webSite": "testwebsite", "email": "testemail", "enabled": true, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateClient(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClient(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var uri m.ClientRedirectURI
	uri.ID = 33
	uri.URI = "/test/test"
	uri.ClientID = 44
	var uris = []m.ClientRedirectURI{uri}
	cuo.RedirectURIs = &uris

	var man m.MockManager
	man.MockClient = cuo
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
	h.GetClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body client in test: ", string(body))

	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.ClientID != 10 || (*bdy.RedirectURIs)[0].ID != 33 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClientNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var man m.MockManager
	man.MockClient = cuo
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
	h.GetClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClientBadParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var man m.MockManager
	man.MockClient = cuo
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
	h.GetClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClientNoParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var man m.MockManager
	man.MockClient = cuo
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
	h.GetClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

//client admin

func TestOauthRestHandler_GetClientAdmin(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var uri m.ClientRedirectURI
	uri.ID = 33
	uri.URI = "/test/test"
	uri.ClientID = 44
	var uris = []m.ClientRedirectURI{uri}
	cuo.RedirectURIs = &uris

	var man m.MockManager
	man.MockClient = cuo
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("clientId", "3")
	// vars := map[string]string{
	// 	"id": "5",
	// }
	// r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetClientAdmin(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body client admin in test: ", string(body))

	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.ClientID != 10 || (*bdy.RedirectURIs)[0].ID != 33 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClientAdminNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var man m.MockManager
	man.MockClient = cuo
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("clientId", "3")
	// vars := map[string]string{
	// 	"id": "5",
	// }
	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetClientAdmin(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClientAdminBadParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var man m.MockManager
	man.MockClient = cuo
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("clientId", "q")
	// vars := map[string]string{
	// 	"id": "q",
	// }
	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetClientAdmin(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClientAdminNoParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var man m.MockManager
	man.MockClient = cuo
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
	h.GetClientAdmin(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

// get list

func TestOauthRestHandler_GetClientlist(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var cuol = []m.Client{cuo}

	var man m.MockManager
	man.MockClientList = cuol
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
	h.GetClientList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || len(bdy) != 1 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetClientListNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false
	var cuol = []m.Client{cuo}

	var man m.MockManager
	man.MockClientList = cuol
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	// vars := map[string]string{
	// 	"id": "5",
	// }
	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetClientList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetSearchClientlist(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var cuol = []m.Client{cuo}

	var man m.MockManager
	man.MockClientList = cuol
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"", "name":"testname", "webSite": "", "email": "", "enabled": false, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.GetClientSearchList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || len(bdy) != 1 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetSearchClientlistBadName(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var cuol = []m.Client{cuo}

	var man m.MockManager
	man.MockClientList = cuol
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"", "name":"", "webSite": "", "email": "", "enabled": false, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.GetClientSearchList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetSearchClientlistBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var cuol = []m.Client{cuo}

	var man m.MockManager
	man.MockClientList = cuol
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"", "name":"testname", "webSite": "", "email": "", "enabled": false, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.GetClientSearchList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	fmt.Println("w.Code in bad body search: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}
func TestOauthRestHandler_GetSearchClientlistBadMedia(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var cuol = []m.Client{cuo}

	var man m.MockManager
	man.MockClientList = cuol
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"", "name":"testname", "webSite": "", "email": "", "enabled": false, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.GetClientSearchList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetSearchClientlistNoAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.Client
	cuo.ClientID = 10
	cuo.Secret = "secrettest"
	cuo.Name = "test name"
	cuo.WebSite = "testwebsite"
	cuo.Email = "testemail"
	cuo.Enabled = true
	cuo.Paid = false

	var cuol = []m.Client{cuo}

	var man m.MockManager
	man.MockClientList = cuol
	oh.Manager = &man

	var asc ac.MockOauthAssets
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = false
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"secret":"", "name":"testname", "webSite": "", "email": "", "enabled": false, "paid": false}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.GetClientSearchList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.Client
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	if w.Code != 401 {
		t.Fail()
	}
}

// delete uri

func TestOauthRestHandler_DeleteClient(t *testing.T) {
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
	h.DeleteClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Success != true {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteClientFail(t *testing.T) {
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
	h.DeleteClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 500 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteClientNotAuth(t *testing.T) {
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
	h.DeleteClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteClientBadParam(t *testing.T) {
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
	h.DeleteClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteClientNoParam(t *testing.T) {
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
	h.DeleteClient(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}
