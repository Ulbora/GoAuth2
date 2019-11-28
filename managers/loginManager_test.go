package managers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	px "github.com/Ulbora/GoProxy"
)

func TestOauthManagerLogin_UserLogin(t *testing.T) {
	var gp px.MockGoProxy
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	var res http.Response
	res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	gp.MockResp = &res
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

func TestOauthManagerLogin_UserLoginEnvBadUrl(t *testing.T) {
	os.Setenv("AUTHENTICATION_SERVICE", "://localhost:3001/rs/user/login")
	var gp px.MockGoProxy
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	var res http.Response
	res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	gp.MockResp = &res
	var man OauthManager
	man.Proxy = gp.GetNewProxy()
	var m Manager
	m = &man
	var l Login
	l.Username = "ken"
	l.Password = "ken"
	l.ClientID = 1
	suc := m.UserLogin(&l)
	if suc {
		t.Fail()
	}
}
