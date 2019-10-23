package managers

import (
	"fmt"
	"testing"
)

func TestOauthManagerJwt_GenerateJwtToken(t *testing.T) {

	var man OauthManager
	//var m Manager
	//m = &man
	var pl Payload
	pl.TokenType = "test"
	pl.UserID = "tester1"
	pl.ClientID = 234
	pl.Subject = "code"
	pl.Issuer = "GoAuth2"
	pl.ExpiresInMinute = 600 //(600 * time.Minute) => (600 * 60) => 36000 minutes => 10 hours
	pl.Grant = "code"
	pl.SecretKey = "secret"
	pl.RoleURIs = []string{"test.com", "test2.com"}
	pl.ScopeList = []string{"web", "sever"}
	token := man.GenerateJwtToken(&pl)
	fmt.Println(token)
	if token == ""{
		t.Fail()
	}

}
