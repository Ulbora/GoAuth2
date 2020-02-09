//Package handlers ...
package handlers

import (
	//"html/template"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	m "github.com/Ulbora/GoAuth2/managers"
	lg "github.com/Ulbora/Level_Logger"
)

func TestOauthWebHandlerToken_AuthCodeToken(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "bearer"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=authorization_code&client_id=3456&client_secret=aaaa45&code=123abc&redirect_uri=http://someTest/test.com"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy *m.Token
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 200 || bdy.AccessToken != "125444" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_AuthCodeTokenBadGrant(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "bearer"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=authorization_code2&client_id=3456&client_secret=aaaa45&code=123abc&redirect_uri=http://someTest/test.com"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy *m.Token
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]

	if w.Code != 302 || loc[0] != "/oauthError?error=invalid_grant" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_AuthCodeTokenBadClient(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "bearer"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=authorization_code&client_id=a3456&client_secret=aaaa45&code=123abc&redirect_uri=http://someTest/test.com"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	fmt.Println("location: ", w.HeaderMap["Location"])
	loc := w.HeaderMap["Location"]

	if w.Code != 302 || loc[0] != "/oauthError?error=invalid_grant" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_AuthCodeTokenFailed(t *testing.T) {
	var om m.MockManager
	//om.MockAuthCodeTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "bearer"
	om.MockToken = tk
	om.MockTokenError = "some_error"
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=authorization_code&client_id=3456&client_secret=aaaa45&code=123abc&redirect_uri=http://someTest/test.com"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy TokenError
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 401 || bdy.Error != "some_error" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_PasswordToken(t *testing.T) {
	var om m.MockManager
	om.MockPasswordTokenSuccess = true
	om.MockUserLoginSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "bearer"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=password&client_id=3456&password=aaaa45&username=tester1"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy *m.Token
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 200 || bdy.AccessToken != "125444" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_PasswordTokenFailedLogin(t *testing.T) {
	var om m.MockManager
	om.MockPasswordTokenSuccess = true
	//om.MockUserLoginSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "bearer"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=password&client_id=3456&password=aaaa45&username=tester1"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy TokenError
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 401 || bdy.Error != "unauthorized_client" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_CredentialsToken(t *testing.T) {
	var om m.MockManager
	om.MockCredentialsTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "bearer"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=client_credentials&client_id=3456&client_secret=aaaa45"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy *m.Token
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 200 || bdy.AccessToken != "125444" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_AuthCodeRefreshToken(t *testing.T) {
	var om m.MockManager
	om.MockAuthCodeRefreshTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "refresh"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=refresh_token&client_id=3456&client_secret=aaaa45&refresh_token=123abc"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy *m.Token
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 200 || bdy.AccessToken != "125444" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_PasswordRefreshToken(t *testing.T) {
	var om m.MockManager
	om.MockPasswordRefreshTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "refresh"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=refresh_token&client_id=3456&refresh_token=123abc"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy *m.Token
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 200 || bdy.AccessToken != "125444" {
		t.Fail()
	}
}

func TestOauthWebHandlerToken_PasswordRefreshTokenCompressed(t *testing.T) {
	var om m.MockManager
	om.MockPasswordRefreshTokenSuccess = true
	var tk m.Token
	tk.AccessToken = "125444"
	tk.TokenType = "refresh"
	om.MockToken = tk
	var ari AuthorizeRequestInfo
	ari.ResponseType = "token"
	ari.ClientID = 1234
	ari.RedirectURI = "http://test.com/test"
	ari.Scope = "web"
	ari.State = "12eee"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	//wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	wh.TokenCompressed = true
	h := wh.GetNewWebHandler()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("grant_type=refresh_token&client_id=3456&refresh_token=123abc"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//s, suc := wh.getSession(r)
	//fmt.Println("suc: ", suc)

	//s.Values["authReqInfo"] = ari
	//s.Save(r, w)

	h.Token(w, r)
	fmt.Println("code: ", w.Code)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy *m.Token
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))

	if w.Code != 200 || bdy.AccessToken != "eNoyNDI1MTEBBAAA//8EMgE1" {
		t.Fail()
	}
}
