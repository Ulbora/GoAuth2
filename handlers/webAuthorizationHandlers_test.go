// Package handlers ...
package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/Ulbora/GoAuth2/managers"
	lg "github.com/Ulbora/Level_Logger"
)

func TestOauthWebHandler_Authorize(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
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
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "http://tester.com/test?code=rr666&state=123" {
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
	var l lg.Logger
	wh.Log = &l
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
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "/authorizeApp" {
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
	var l lg.Logger
	wh.Log = &l
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
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "/oauthError?error=access_denied" {
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
	var l lg.Logger
	wh.Log = &l
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
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "http://tester.com/test?token=gjfldflkl&token_type=bearer&state=123" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeImplicitCompressed(t *testing.T) {
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
	var l lg.Logger
	wh.Log = &l
	wh.Manager = &om
	wh.TokenCompressed = true
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
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "http://tester.com/test?token=eNpKz0rLSUnLyc4BBAAA//8SXAOx&token_type=bearer&state=123" {
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
	var l lg.Logger
	wh.Log = &l
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
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "/authorizeApp" {
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
	var l lg.Logger
	wh.Log = &l
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
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "/oauthError?error=access_denied" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeInvalidGrant(t *testing.T) {
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
	var l lg.Logger
	wh.Log = &l
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=someGrant&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "/oauthError?error=invalid_grant" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeBadSession(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
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
	h.Authorize(w, nil)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	//loc := w.HeaderMap["Location"]
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeNotLoggedIn(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Save(r, w)
	h.Authorize(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "/login" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeApp(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.AuthCodeClient
	acc.Valid = true
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockAuthCodeClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "code"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.AuthorizeApp(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeAppNotValid(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.AuthCodeClient
	acc.Valid = false
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockAuthCodeClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "code"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.AuthorizeApp(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeAppNoAuthInfo(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.AuthCodeClient
	acc.Valid = true
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockAuthCodeClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "code"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	//s.Values["authReqInfo"] = ari

	s.Save(r, w)
	h.AuthorizeApp(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeAppBadSession(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.AuthCodeClient
	acc.Valid = true
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockAuthCodeClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "code"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=code&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = ari

	s.Save(r, w)
	h.AuthorizeApp(w, nil)
	fmt.Println("code: ", w.Code)
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeAppToken(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.ImplicitClient
	acc.Valid = true
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockImplicitClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.AuthorizeApp(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeAppTokenNotAuth(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.ImplicitClient
	acc.Valid = false
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockImplicitClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.AuthorizeApp(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeAppTokenNotResponsType(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.ImplicitClient
	acc.Valid = true
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockImplicitClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "sometype"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&client_id=125&redirect_uri=http://tester.com/test&scope=web&state=123", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.AuthorizeApp(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserCode(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	var acc m.ImplicitClient
	acc.Valid = true
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockImplicitClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "code"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "http://test.com/test?code=rr666&state=12eee" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserCodeFailedAuth(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = false
	var acc m.ImplicitClient
	acc.Valid = true
	acc.ClientName = "test client"
	acc.WebSite = "www.test.com"
	om.MockImplicitClient = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "code"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))

	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserToken(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockImplicitAuthorizeSuccess = true
	var acc m.ImplicitReturn
	acc.ID = 55
	acc.Token = "lllkldskldfk"

	om.MockImplicitReturn = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "http://test.com/test?token=lllkldskldfk&token_type=bearer&state=12eee" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserTokenCompressed(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockImplicitAuthorizeSuccess = true
	var acc m.ImplicitReturn
	acc.ID = 55
	acc.Token = "lllkldskldfk"

	om.MockImplicitReturn = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	wh.TokenCompressed = true
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "http://test.com/test?token=eNrKycnJzkkpzs5JScsGBAAA//8gswT/&token_type=bearer&state=12eee" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserTokenFailedAuth(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockImplicitAuthorizeSuccess = false
	var acc m.ImplicitReturn
	acc.ID = 55
	acc.Token = "lllkldskldfk"

	om.MockImplicitReturn = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))

	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserBadResponseType(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockImplicitAuthorizeSuccess = true
	var acc m.ImplicitReturn
	acc.ID = 55
	acc.Token = "lllkldskldfk"

	om.MockImplicitReturn = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "someResponse"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = &ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	loc := w.Header().Get("Location")
	if w.Code != 302 || loc != "http://test.com/test?error=access_denied&state=12eee" {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserTokenBadSession(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockImplicitAuthorizeSuccess = true
	var acc m.ImplicitReturn
	acc.ID = 55
	acc.Token = "lllkldskldfk"

	om.MockImplicitReturn = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, nil)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))
	if w.Code != 500 {
		t.Fail()
	}
}

func TestOauthWebHandler_AuthorizeByUserTokenNoInfo(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockImplicitAuthorizeSuccess = true
	var acc m.ImplicitReturn
	acc.ID = 55
	acc.Token = "lllkldskldfk"

	om.MockImplicitReturn = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?response_type=token&authorize=true", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	//s.Values["authReqInfo"] = ari

	s.Save(r, w)
	h.ApplicationAuthorizationByUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))

	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandler_Error(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockImplicitAuthorizeSuccess = true
	var acc m.ImplicitReturn
	acc.ID = 55
	acc.Token = "lllkldskldfk"

	om.MockImplicitReturn = acc
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test?error=someError", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["user"] = "tester"
	s.Values["authReqInfo"] = ari

	s.Save(r, w)
	h.OauthError(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.Header().Get("Location"))

	if w.Code != 200 {
		t.Fail()
	}
}
