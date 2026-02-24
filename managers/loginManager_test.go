package managers

import (
	"bytes"
	"fmt"
	"io"

	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	au "github.com/Ulbora/auth_interface"
	dau "github.com/Ulbora/default_auth"
)

func TestOauthManagerLogin_UserLogin(t *testing.T) {
	// var gp px.MockGoProxy
	// gp.MockDoSuccess1 = true
	// gp.MockRespCode = 200
	var proxy px.GoProxy
	var da dau.MockDefaultAuth
	da.MockValid = true
	da.AuthServerURL = authenticationServiceLocal
	da.Proxy = proxy.GetNewProxy()
	//ai := da.GetNew()

	var res http.Response
	res.Body = io.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	//gp.MockResp = &res
	var man OauthManager
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

// func TestOauthManagerLogin_UserLoginEnvBadUrl(t *testing.T) {
// 	os.Setenv("AUTHENTICATION_SERVICE", "://localhost:3001/rs/user/login")
// 	// var gp px.MockGoProxy
// 	// gp.MockDoSuccess1 = true
// 	// gp.MockRespCode = 200
// 	var proxy px.GoProxy
// 	var da dau.MockDefaultAuth
// 	da.MockValid = false
// 	da.AuthServerURL = authenticationServiceLocal
// 	da.Proxy = proxy.GetNewProxy()
// 	var res http.Response
// 	res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
// 	//gp.MockResp = &res
// 	var man OauthManager
// 	man.AuthService = da.GetNew()
// 	//man.Proxy = gp.GetNewProxy()
// 	var m Manager
// 	m = &man
// 	var l au.Login
// 	l.Username = "ken"
// 	l.Password = "ken"
// 	l.ClientID = 1
// 	suc := m.UserLogin(&l)
// 	if suc {
// 		t.Fail()
// 	}
// }
