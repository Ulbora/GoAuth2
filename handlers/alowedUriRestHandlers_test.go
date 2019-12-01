//Package handlers ...
package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/Ulbora/GoAuth2/managers"
	oc "github.com/Ulbora/GoAuth2/oauthclient"
	ac "github.com/Ulbora/GoAuth2/rolecontrol"
)

func TestOauthRestHandler_AddAllowedURISuperBadMedia(t *testing.T) {
	var oh OauthRestHandler

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
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURISuperFailAdd(t *testing.T) {
	var oh OauthRestHandler

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
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURINoAssetControl(t *testing.T) {
	var oh OauthRestHandler

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
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddAllowedURIBadMedia(t *testing.T) {
	var oh OauthRestHandler

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

func TestOauthRestHandler_UpdateAllowedURISuperBadMedia(t *testing.T) {
	var oh OauthRestHandler

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
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthRestHandler_UpdateAllowedURISuperFailAdd(t *testing.T) {
	var oh OauthRestHandler

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
