//Package handlers ...
package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	m "github.com/Ulbora/GoAuth2/managers"
)

func TestOauthWebHandlerLogin_Login(t *testing.T) {
	var om m.MockManager

	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Login(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthWebHandlerLogin_LoginUser(t *testing.T) {
	var om m.MockManager
	om.MockUserLoginSuccess = true
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester1&password=somepassword"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["authReqInfo"] = ari
	s.Save(r, w)

	h.LoginUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]

	if w.Code != 302 || loc[0] != "/oauth/authorize?response_type=token&client_id=1234&redirect_uri=http://test.com/test&scope=web&state=12eee" {
		t.Fail()
	}
}

func TestOauthWebHandlerLogin_LoginUserBadGrant(t *testing.T) {
	var om m.MockManager
	om.MockUserLoginSuccess = true
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token2"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester1&password=somepassword"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["authReqInfo"] = ari
	s.Save(r, w)

	h.LoginUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]

	if w.Code != 302 || loc[0] != "/oauthError?error=invalid_grant" {
		t.Fail()
	}
}

func TestOauthWebHandlerLogin_LoginUserFailLogin(t *testing.T) {
	var om m.MockManager
	//om.MockUserLoginSuccess = true
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester1&password=somepassword"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["authReqInfo"] = ari
	s.Save(r, w)

	h.LoginUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]

	if w.Code != 302 || loc[0] != "/login" {
		t.Fail()
	}
}

func TestOauthWebHandlerLogin_LoginUserNoSessionInfo(t *testing.T) {
	var om m.MockManager
	om.MockUserLoginSuccess = true
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester1&password=somepassword"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	s.Save(r, w)

	h.LoginUser(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]

	if w.Code != 302 || loc[0] != "/oauthError?error=invalid_grant" {
		t.Fail()
	}
}

func TestOauthWebHandlerLogin_LoginUserNoSession(t *testing.T) {
	var om m.MockManager
	om.MockUserLoginSuccess = true
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester1&password=somepassword"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := wh.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["authReqInfo"] = ari
	s.Save(r, w)

	h.LoginUser(w, nil)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])

	if w.Code != 500 {
		t.Fail()
	}
}
