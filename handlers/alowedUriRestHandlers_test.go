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

func TestOauthRestHandler_AddAllowedURISuperBadMedia(t *testing.T) {
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
	h.AddAllowedURISuper(w, r)
	if w.Code != 415 {
		t.Fail()
	}

}

func TestOauthRestHandler_AddAllowedURISuperBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	var robj m.ClientAllowedURI
	robj.ID = 3
	robj.URI = "/test"
	robj.ClientID = 1
	//aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURISuper(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURISuper(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURISuper(w, r)
	hd := w.Header()
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURISuperFailAdd(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURISuper(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURISuperNotAuth(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURISuper(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

// add url non super

func TestOauthRestHandler_AddAllowedURI(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURI(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURINoAssetControl(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 5
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	asc.MockAllowedRole = "admin"
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURI(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURIBadMedia(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURI(w, r)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURIAddFail(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURI(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURINotAuth(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddAllowedURI(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURIBadBody(t *testing.T) {
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
	h.AddAllowedURI(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

// update uri super
func TestOauthRestHandler_UpdateAllowedURISuperBadMedia(t *testing.T) {
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
	h.UpdateAllowedURISuper(w, r)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURISuperBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	var robj m.ClientAllowedURI
	robj.ID = 3
	robj.URI = "/test"
	robj.ClientID = 1
	//aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURISuper(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURISuper(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURISuper(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURISuperFailAdd(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURISuper(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURISuperNotAuth(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURISuper(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

// update uri

// add url non super

func TestOauthRestHandler_updateAllowedURI(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockUpdateSuccess1 = true
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURI(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURINoAssetControl(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockUpdateSuccess1 = true
	oh.Manager = &man

	var asc ac.MockOauthAssets
	asc.MockSuccess = false
	asc.MockAllowedRole = "admin"
	oh.AssetControl = &asc

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct
	fmt.Println("oh.Client: ", oh.Client)

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURI(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURIBadMedia(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockUpdateSuccess1 = true
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURI(w, r)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURIAddFail(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockUpdateSuccess1 = false
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURI(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURINotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockUpdateSuccess1 = true
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "url":"/test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.UpdateAllowedURI(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURIBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockUpdateSuccess1 = true
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
	h.UpdateAllowedURI(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetAllowedURI(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	man.MockClientAllowedURI = cuo
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
	h.GetAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.ID == 0 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetAllowedURINotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	man.MockClientAllowedURI = cuo
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
	h.GetAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetAllowedURIBadParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	man.MockClientAllowedURI = cuo
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
	h.GetAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_GetAllowedURINoParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	man.MockClientAllowedURI = cuo
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
	h.GetAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

// get list

func TestOauthRestHandler_GetAllowedURIlist(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var cuol = []m.ClientAllowedURI{cuo}

	var man m.MockManager
	man.MockClientAllowedURIList = cuol
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
	h.GetAllowedURIList(w, r)
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

func TestOauthRestHandler_GetAllowedURIListNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	man.MockClientAllowedURI = cuo
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
	h.GetAllowedURIList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetAllowedURIListBadParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	man.MockClientAllowedURI = cuo
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
	h.GetAllowedURIList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_GetAllowedURIListNoParam(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l
	var cuo m.ClientAllowedURI
	cuo.ID = 4
	cuo.URI = "/test/url"
	cuo.ClientID = 10

	var man m.MockManager
	man.MockClientAllowedURI = cuo
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
	h.GetAllowedURIList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

// delete uri

func TestOauthRestHandler_DeleteAllowedURI(t *testing.T) {
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
	h.DeleteAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Success != true {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteAllowedURIFail(t *testing.T) {
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
	h.DeleteAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 500 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteAllowedURINotAuth(t *testing.T) {
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
	h.DeleteAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteAllowedURIBadParam(t *testing.T) {
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
	h.DeleteAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteAllowedURINoParam(t *testing.T) {
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
	h.DeleteAllowedURI(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}
