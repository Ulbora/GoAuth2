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
)

// validate token

func TestOauthRestHandler_ValidateToken(t *testing.T) {
	var oh OauthRestHandler

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true
	oh.Manager = &man

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"accessToken":"someaccesstoken", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !bdy.Valid {
		t.Fail()
	}
}

func TestOauthRestHandler_ValidateTokenNotValid(t *testing.T) {
	var oh OauthRestHandler

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = false
	oh.Manager = &man

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"accessToken":"someaccesstoken", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Valid {
		t.Fail()
	}
}

func TestOauthRestHandler_ValidateTokenBadMedia(t *testing.T) {
	var oh OauthRestHandler

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true
	oh.Manager = &man

	h := oh.GetNewRestHandler()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"accessToken":"someaccesstoken", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_ValidateTokenBadBody(t *testing.T) {
	var oh OauthRestHandler

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true
	oh.Manager = &man

	h := oh.GetNewRestHandler()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"accessToken":"someaccesstoken", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 {
		t.Fail()
	}
}
