package oauthclient

import (
	"fmt"
	"net/http"
	"testing"
	//m "github.com/Ulbora/GoAuth2/managers"
)

func TestMockOauthClient_Authorize(t *testing.T) {

	var oc MockOauthClient
	oc.MockValid = true
	c := oc.GetNewClient()
	var cl Claim
	cl.Role = "testRole"
	cl.URL = "testURL"
	cl.Scope = "web"
	r, _ := http.NewRequest("GET", "/testurl", nil)
	r.Header.Set("Authorization", "Bearer jdljdfldjslkjdslkldksldfks")
	r.Header.Set("hashed", "true")
	r.Header.Set("clientId", "22")
	r.Header.Set("userId", "lfo")

	suc := c.Authorize(r, &cl)
	fmt.Println("suc", suc)
	if !suc {
		t.Fail()
	}
}
