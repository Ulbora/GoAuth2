// Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
)

func TestOauthRestHandlerLogger_SetDebugLogLevel(t *testing.T) {
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"debug"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "DEBUG" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetDebugLogLevelBadReq(t *testing.T) {
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"debug"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetInfoLogLevel(t *testing.T) {
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"logLevel":"info"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "INFO" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetAllLogLevel(t *testing.T) {
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"logLevel":"all"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "ALL" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetOffLogLevel(t *testing.T) {
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "OFF" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetOffLogLevelLogKey(t *testing.T) {
	os.Setenv("LOGGING_KEY", "45sdbb2345")
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "OFF" {
		t.Fail()
	}
	os.Unsetenv("LOGGING_KEY")
}

func TestOauthRestHandlerLogger_SetOffLogLevelLogKeyWrongKey(t *testing.T) {
	os.Setenv("LOGGING_KEY", "45sdbb23455")
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
	os.Unsetenv("LOGGING_KEY")
}

func TestOauthRestHandlerLogger_SetBadMediaLogLevel(t *testing.T) {
	var oh OauthRestHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNewRestHandler()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 415 {
		t.Fail()
	}
}
