package managers

import (
	"testing"

	px "github.com/Ulbora/GoProxy"
)

func TestOauthManagerLogin_UserLogin(t *testing.T) {
	var gp px.GoProxy
	var man OauthManager
	man.Proxy = gp.GetNewProxy()
	var m Manager
	m = &man
	var l Login
	l.Username = "ken"
	l.Password = "ken"
	l.ClientID = 1
	suc := m.UserLogin(&l)
	if !suc {
		t.Fail()
	}
}
