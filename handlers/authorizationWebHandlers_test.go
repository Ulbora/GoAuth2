//Package handlers ...
package handlers

import (
	"fmt"
	m "github.com/Ulbora/GoAuth2/managers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOauthWebHandler_Authorize(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]
	if w.Code != 302 || loc[0] != "http://tester.com/test?code=rr666&state=123" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeNotAuth(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = false
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]
	if w.Code != 302 || loc[0] != "/authorizeApp" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeAuthCodeFail(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = false
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]
	if w.Code != 302 || loc[0] != "/oauthError?error=access_denied" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeImplicit(t *testing.T) {
	var om m.MockManager
	om.MockImplicitAuthorized = true
	om.MockImplicitAuthorizeSuccess = true
	var impRtn m.ImplicitReturn
	impRtn.ID = 5
	impRtn.Token = "gjfldflkl"
	om.MockImplicitReturn = impRtn
	//om.MockAuthCode = 55
	//om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]
	if w.Code != 302 || loc[0] != "http://tester.com/test?token=gjfldflkl&token_type=bearer&state=123" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeImplicitNotAuth(t *testing.T) {
	var om m.MockManager
	om.MockImplicitAuthorized = false
	om.MockImplicitAuthorizeSuccess = true
	var impRtn m.ImplicitReturn
	impRtn.ID = 5
	impRtn.Token = "gjfldflkl"
	om.MockImplicitReturn = impRtn
	//om.MockAuthCode = 55
	//om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]
	if w.Code != 302 || loc[0] != "/authorizeApp" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeImplicitFailed(t *testing.T) {
	var om m.MockManager
	om.MockImplicitAuthorized = true
	om.MockImplicitAuthorizeSuccess = false
	var impRtn m.ImplicitReturn
	impRtn.ID = 5
	impRtn.Token = "gjfldflkl"
	om.MockImplicitReturn = impRtn
	//om.MockAuthCode = 55
	//om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]
	if w.Code != 302 || loc[0] != "/oauthError?error=access_denied" {
		t.Fail()
	}
}
