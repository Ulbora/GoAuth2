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
