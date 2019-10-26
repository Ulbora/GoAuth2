package managers

import (
	"fmt"
	"testing"
)

var tkn string

func TestOauthManagerJwt_GenerateJwtToken(t *testing.T) {

	var man OauthManager
	var pl Payload
	pl.TokenType = "test"
	pl.UserID = "tester1"
	pl.ClientID = 234
	pl.Subject = "code"
	pl.Issuer = "GoAuth2"
	pl.Audience = "GoAuth2.com"
	pl.ExpiresInMinute = 600 //(600 * time.Minute) => (600 * 60) => 36000 minutes => 10 hours
	pl.Grant = "code"
	pl.SecretKey = "secret"
	pl.RoleURIs = []string{"test.com", "test2.com"}
	pl.ScopeList = []string{"web", "sever"}
	token := man.GenerateJwtToken(&pl)
	fmt.Println(token)
	if token == "" {
		t.Fail()
	} else {
		tkn = token
	}
}
func TestOauthManagerJwt_ValidateJwtToken(t *testing.T) {
	var man OauthManager

	suc, pl := man.Validate(tkn, "secret")
	if !suc || pl.TokenType != "test" || pl.UserID != "tester1" || pl.ClientID != 234 {
		t.Fail()
	} else {
		fmt.Println("pl: ", pl)
		fmt.Println("Audience: ", pl.Audience)
	}

}
