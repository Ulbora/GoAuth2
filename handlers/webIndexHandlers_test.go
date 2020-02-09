//Package handlers ...
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

func TestOauthWebHandlerIndex_Index(t *testing.T) {
	var om m.MockManager

	om.MockAuthCodeAuthorized = true
	om.MockAuthCodeAuthorizeSuccess = true
	om.MockAuthCode = 55
	om.MockAuthCodeString = "rr666"

	var wh OauthWebHandler
	var l lg.Logger
	wh.Log = &l
	wh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	wh.Manager = &om
	h := wh.GetNewWebHandler()
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Index(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}
