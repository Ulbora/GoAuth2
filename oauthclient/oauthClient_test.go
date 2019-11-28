package oauthclient

import (
	"fmt"
	"net/http"
	"testing"

	m "github.com/Ulbora/GoAuth2/managers"
)

func TestOauthClient_Authorize(t *testing.T) {
	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true

	var oc OauthClient
	oc.Manager = &man
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
