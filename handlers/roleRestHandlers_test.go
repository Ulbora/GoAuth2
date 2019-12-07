//Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	m "github.com/Ulbora/GoAuth2/managers"
	ac "github.com/Ulbora/GoAuth2/rolecontrol"
	"github.com/gorilla/mux"

	oc "github.com/Ulbora/GoAuth2/oauthclient"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOauthRestHandler_AddRoleSuperBadMedia(t *testing.T) {
	var oh OauthRestHandler

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	// r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRoleSuper(w, r)
	if w.Code != 415 {
		t.Fail()
	}

}

func TestOauthRestHandler_AddRoleSuperBadBody(t *testing.T) {
	var oh OauthRestHandler

	var oct oc.MockOauthClient
	oct.MockValid = true
	oh.Client = &oct

	h := oh.GetNewRestHandler()
	var robj m.ClientRole
	robj.ID = 3
	robj.Role = "test"
	robj.ClientID = 1
	//aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRoleSuper(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleSuper(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRoleSuper(w, r)
	hd := w.Header()
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleSuperFailAdd(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRoleSuper(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleSuperNotAuth(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRoleSuper(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

// add url non super

func TestOauthRestHandler_AddRole(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRole(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleNoAssetControl(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRole(w, r)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleBadMedia(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRole(w, r)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleAddFail(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRole(w, r)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleNotAuth(t *testing.T) {
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

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":3, "role":"test", "clientId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.AddRole(w, r)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_AddRoleBadBody(t *testing.T) {
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
	h.AddRole(w, r)
	if w.Code != 400 {
		t.Fail()
	}
}

// get list

func TestOauthRestHandler_GetRolelist(t *testing.T) {
	var oh OauthRestHandler
	var cuo m.ClientRole
	cuo.ID = 4
	cuo.Role = "test"
	cuo.ClientID = 10

	var cuol = []m.ClientRole{cuo}

	var man m.MockManager
	man.MockClientRoleList = cuol
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
	h.GetRoleList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.ClientRole
	json.Unmarshal(body, &bdy)
	fmt.Println("body roles: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || len(bdy) != 1 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetRoleListNotAuth(t *testing.T) {
	var oh OauthRestHandler
	var cuo m.ClientRole
	cuo.ID = 4
	cuo.Role = "test"
	cuo.ClientID = 10

	var man m.MockManager
	//man.MockClientRoleList = cuo
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
	h.GetRoleList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientRole
	json.Unmarshal(body, &bdy)
	fmt.Println("body roles: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_GetRoleListBadParam(t *testing.T) {
	var oh OauthRestHandler
	var cuo m.ClientRole
	cuo.ID = 4
	cuo.Role = "test"
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
	h.GetRoleList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientRole
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_GetRoleListNoParam(t *testing.T) {
	var oh OauthRestHandler
	var cuo m.ClientRole
	cuo.ID = 4
	cuo.Role = "test"
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
	h.GetRoleList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientRole
	json.Unmarshal(body, &bdy)
	fmt.Println("body roles: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

// delete uri

func TestOauthRestHandler_DeleteRole(t *testing.T) {
	var oh OauthRestHandler

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
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Success != true {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteRoleFail(t *testing.T) {
	var oh OauthRestHandler

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
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 500 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteRoleNotAuth(t *testing.T) {
	var oh OauthRestHandler

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
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteRoleBadParam(t *testing.T) {
	var oh OauthRestHandler

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
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestOauthRestHandler_DeleteRoleNoParam(t *testing.T) {
	var oh OauthRestHandler

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
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.ClientAllowedURI
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	if w.Code != 400 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}
