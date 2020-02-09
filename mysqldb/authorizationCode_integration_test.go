// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbAci db.Database
var odbAci odb.Oauth2DB
var cidAci int64
var spIDAci int64
var spID2Aci int64

func TestMySQLOauthDBACi_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAci = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbAci

	odbAci = &moadb

	dbAci.Connect()
}

func TestMySQLOauthDBACi_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbAci.AddClient(&c, nil)
	fmt.Println("client add res: ", res)
	fmt.Println("client id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidAci = id
	}
}

func TestMySQLOauthDBACi_AddAuthorizationCode(t *testing.T) {

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAci
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445a"

	res, id := odbAci.AddAuthorizationCode(&ac, &at, &rt, nil)
	if !res || id < 1 {
		t.Fail()
	}
}

// func TestMySQLOauthDBACi_DeleteAuthorizationCodeScope1(t *testing.T) {
// 	res := odbAci.DeleteAuthorizationCode(cidAci, "1234")
// 	if !res {
// 		t.Fail()
// 	}
// }

func TestMySQLOauthDBACi_AddAuthorizationCodeScope(t *testing.T) {

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAci
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445b"
	var scope = []string{"test1", "test2"}

	res, id := odbAci.AddAuthorizationCode(&ac, &at, &rt, &scope)
	if !res || id < 1 {
		t.Fail()
	} else {
		spID2Aci = id
	}
}

func TestMySQLOauthDBACi_GetAuthCodeScopeList(t *testing.T) {
	res := odbAci.GetAuthorizationCodeScopeList(spID2Aci)
	fmt.Println("auth code scope in get: ", res)
	if res == nil || (*res)[0].Scope != "test1" {
		t.Fail()
	}
}

func TestMySQLOauthDBACi_AddAuthCodeRevolk(t *testing.T) {
	var rv odb.AuthCodeRevolk
	rv.AuthorizationCode = spID2Aci
	res, id := odbAci.AddAuthCodeRevolk(nil, &rv)
	fmt.Println("revolk id: ", id)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBACi_UpdateAuthCode(t *testing.T) {
	var ac odb.AuthorizationCode
	ac.RandonAuthCode = "13445bb"
	ac.AlreadyUsed = true
	ac.AuthorizationCode = spID2Aci
	res := odbAci.UpdateAuthorizationCode(&ac)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBACi_UpdateAuthCodeToken(t *testing.T) {
	ac := odbAci.GetAuthorizationCodeByCode("13445bb")

	fmt.Println("auth code in update token: ", ac)
	var rt odb.RefreshToken
	rt.Token = "somereftoken2upd"
	rfs, rfid := odbAci.AddRefreshToken(nil, &rt)
	fmt.Println("new refresh token: ", rfs)
	if rfs {
		at := odbAci.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at in update token: ", at)
		at.Token = "someacctokenupd"
		at.Expires = time.Now()
		at.RefreshTokenID = rfid
		tt := time.Now()
		ac.Expires = tt
		res := odbAci.UpdateAuthorizationCodeAndToken(ac, at)
		fmt.Println("auth code update token suc: ", res)
		ac2 := odbAci.GetAuthorizationCodeByCode("13445bb")
		fmt.Println("auth2 code in update token: ", ac2)
		fmt.Println("tt in update token: ", tt.UTC())
		fmt.Println("expires in update token: ", ac2.Expires)
		at2 := odbAci.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at2 in update token: ", at2)
		if !res || at2.Token != "someacctokenupd" {
			t.Fail()
		}
	}

}

func TestMySQLOauthDBACi_GetAuthCodeByCode(t *testing.T) {
	res := odbAci.GetAuthorizationCodeByCode("13445bb")
	fmt.Println("auth code in get: ", res)
	if res == nil || res.RandonAuthCode != "13445bb" || res.AlreadyUsed != true {
		t.Fail()
	}
}

func TestMySQLOauthDBACi_GetAuthCodeByClient(t *testing.T) {
	res := odbAci.GetAuthorizationCode(cidAci, "1234")
	fmt.Println("auth code in get by client: ", res)
	if len(*res) < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBACi_GetAuthCodeByScope(t *testing.T) {
	res := odbAci.GetAuthorizationCodeByScope(cidAci, "1234", "test1")
	fmt.Println("auth code in get by scope: ", res)
	if len(*res) < 1 || (*res)[0].Scope != "test1" {
		t.Fail()
	}
}

func TestMySQLOauthDBACi_DeleteAuthorizationCode(t *testing.T) {
	res := odbAci.DeleteAuthorizationCode(cidAci, "1234")
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBACi_DeleteClient(t *testing.T) {
	suc := odbAci.DeleteClient(cidAci)
	if !suc {
		t.Fail()
	}
}
