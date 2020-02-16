// +build integration move to top

package managers

import (
	"testing"

	"fmt"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	au "github.com/Ulbora/auth_interface"
	dau "github.com/Ulbora/default_auth"
)

func TestOauthManagerLoginInt_UserLogin(t *testing.T) {
	var gp px.GoProxy
	var da dau.DefaultAuth
	var authURL = "http://localhost:3001/rs/user/login"
	da.AuthServerURL = authURL
	da.Proxy = gp.GetNewProxy()
	var man OauthManager
	var ll lg.Logger
	man.Log = &ll
	//moadb.Log = &l
	man.AuthService = da.GetNew()
	//man.Proxy = gp.GetNewProxy()
	var m Manager
	m = &man
	var l au.Login
	l.Username = "admin"
	l.Password = "admin"
	l.ClientID = 10
	suc := m.UserLogin(&l)
	fmt.Println("suc: ", suc)
	if !suc {
		t.Fail()
	}
}
