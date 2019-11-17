package managers

import (
	"fmt"
	"testing"
	"time"
)

var tkn string

var tknExp string

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
	var ruList []RoleURI
	var ru1 RoleURI
	ru1.ClientRoleID = 1
	ru1.Role = "user"
	ru1.ClientAllowedURIID = 2
	ru1.ClientAllowedURI = "test.com"
	ru1.ClientID = 5
	ruList = append(ruList, ru1)

	var ru2 RoleURI
	ru2.ClientRoleID = 12
	ru2.Role = "user"
	ru2.ClientAllowedURIID = 21
	ru2.ClientAllowedURI = "test2.com"
	ru2.ClientID = 5
	ruList = append(ruList, ru2)

	pl.RoleURIs = ruList

	pl.ScopeList = []string{"web", "sever"}
	token := man.GenerateJwtToken(&pl)
	fmt.Println(token)
	if token == "" {
		t.Fail()
	} else {
		tkn = token
	}
}

func TestOauthManagerJwt_GenerateJwtTokenExp(t *testing.T) {

	var man OauthManager
	var pl Payload
	pl.TokenType = "test"
	pl.UserID = "tester1"
	pl.ClientID = 234
	pl.Subject = "code"
	pl.Issuer = "GoAuth2"
	pl.Audience = "GoAuth2.com"
	pl.ExpiresInMinute = 1 //(600 * time.Minute) => (600 * 60) => 36000 minutes => 10 hours
	pl.Grant = "code"
	pl.SecretKey = "secret"
	var ruList []RoleURI
	var ru1 RoleURI
	ru1.ClientRoleID = 1
	ru1.Role = "user"
	ru1.ClientAllowedURIID = 2
	ru1.ClientAllowedURI = "test.com"
	ru1.ClientID = 5
	ruList = append(ruList, ru1)

	var ru2 RoleURI
	ru2.ClientRoleID = 12
	ru2.Role = "user"
	ru2.ClientAllowedURIID = 21
	ru2.ClientAllowedURI = "test2.com"
	ru2.ClientID = 5
	ruList = append(ruList, ru2)

	pl.RoleURIs = ruList

	pl.ScopeList = []string{"web", "sever"}
	token := man.GenerateJwtToken(&pl)
	fmt.Println("expiring Token:", token)
	if token == "" {
		t.Fail()
	} else {
		tknExp = token
	}
}

func TestOauthManagerJwt_ValidateJwtToken(t *testing.T) {
	var man OauthManager

	suc, pl := man.ValidateJwt(tkn, "secret")
	if !suc || pl.TokenType != "test" || pl.UserID != "tester1" || pl.ClientID != 234 {
		t.Fail()
	} else {
		fmt.Println("pl: ", pl)
		fmt.Println("Audience: ", pl.Audience)
	}

}

func TestOauthManagerJwt_ValidateJwtTokenExpired(t *testing.T) {
	time.Sleep(2 * time.Minute)
	var man OauthManager

	suc, pl := man.ValidateJwt(tknExp, "secret")
	fmt.Println("suc: ", suc)
	if suc || pl.TokenType != "test" || pl.UserID != "tester1" || pl.ClientID != 234 {
		t.Fail()
	} else {
		fmt.Println("pl: ", pl)
		fmt.Println("Audience: ", pl.Audience)
	}

}
