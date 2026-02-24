// Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	cp "github.com/Ulbora/GoAuth2/compresstoken"
	m "github.com/Ulbora/GoAuth2/managers"
	lg "github.com/Ulbora/Level_Logger"
)

// validate token

func TestOauthRestHandler_ValidateToken(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true
	oh.Manager = &man

	h := oh.GetNewRestHandler()

	aJSON := io.NopCloser(bytes.NewBufferString(`{"accessToken":"someaccesstoken", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !bdy.Valid {
		t.Fail()
	}
}

func TestOauthRestHandler_ValidateTokenCompressed(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true
	oh.Manager = &man
	oh.TokenCompressed = true

	h := oh.GetNewRestHandler()
	var token = "jdljdfldjslkjdsdfgdfgdffgdfgfdfgdfgdfgdfgdfdfdfdfdfdfdfdfgdgdfgdffgdfgdfdfgfdfgdfdfgddddgdgdgdgdgdgdgddggdgdgdgdggdfgdfgdfgdgflkldksldfks"
	var jc cp.JwtCompress
	tkn := jc.CompressJwt(token)
	fmt.Println("compressed token in test", tkn)
	fmt.Println("uncompressed token in test", jc.UnCompressJwt(tkn))

	aJSON := io.NopCloser(bytes.NewBufferString(`{"accessToken":"eNpUjFEKRDEIAy87JKD58/6w4NJHSwYcRVKkUKhJF4O87NDZ/rwx1+cedATAT/DnV6OVDj1BPb8AAAD//8ZtNs8=", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !bdy.Valid {
		t.Fail()
	}
}

func TestOauthRestHandler_ValidateTokenNotValid(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = false
	oh.Manager = &man

	h := oh.GetNewRestHandler()

	aJSON := io.NopCloser(bytes.NewBufferString(`{"accessToken":"someaccesstoken", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Valid {
		t.Fail()
	}
}

func TestOauthRestHandler_ValidateTokenBadMedia(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true
	oh.Manager = &man

	h := oh.GetNewRestHandler()

	aJSON := io.NopCloser(bytes.NewBufferString(`{"accessToken":"someaccesstoken", "hashed": false, "userId":"someUser", "clientId": 2, "role": "someRole", "uri": "someUri", "scope":"someScope"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ValidateAccessToken(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_ValidateTokenBadBody(t *testing.T) {
	var oh OauthRestHandler
	var l lg.Logger
	oh.Log = &l

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
	body, _ := io.ReadAll(resp.Body)
	var bdy ValidationResponse
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 {
		t.Fail()
	}
}
