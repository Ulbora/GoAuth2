//Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	// m "github.com/Ulbora/GoAuth2/managers"
	// oa "github.com/Ulbora/GoAuth2/oauthclient"
	// rc "github.com/Ulbora/GoAuth2/rolecontrol"
)

type testObj struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

func TestOauthRestHandler_ProcessBody(t *testing.T) {
	var oh OauthRestHandler
	var robj testObj
	robj.Valid = true
	robj.Code = "3"
	// var res http.Response
	// res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	var sURL = "http://localhost/test"
	aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
	var obj testObj
	suc, _ := oh.ProcessBody(r, &obj)
	if !suc || obj.Valid != true || obj.Code != "3" {
		t.Fail()
	}
}

func TestOauthRestHandler_ProcessBodyBadObj(t *testing.T) {
	var oh OauthRestHandler
	var robj testObj
	robj.Valid = true
	robj.Code = "3"
	// var res http.Response
	// res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	var sURL = "http://localhost/test"
	aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
	var obj testObj
	suc, _ := oh.ProcessBody(r, nil)
	if suc || obj.Valid != false || obj.Code != "" {
		t.Fail()
	}
}
