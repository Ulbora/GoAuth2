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

var dbAcRvi db.Database
var odbAcRvi odb.Oauth2DB
var cidAcRvi int64
var acIDAcRvi int64
var spID2AcRvi int64

func TestMySQLOauthDBAcRvi_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAcRvi = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbAcRvi

	odbAcRvi = &moadb

	dbAcRvi.Connect()
}

func TestMySQLOauthDBAcRvi_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbAcRvi.AddClient(&c, nil)
	fmt.Println("client add res: ", res)
	fmt.Println("client id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidAcRvi = id
	}
}

func TestMySQLOauthDBAcRvi_AddAuthorizationCode(t *testing.T) {

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAcRvi
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445"

	res, id := odbAcRvi.AddAuthorizationCode(&ac, &at, &rt, nil)

	if !res || id < 1 {
		t.Fail()
	} else {
		acIDAcRvi = id
	}
}

func TestMySQLOauthDBAcRvi_AddAuthCodeRevolk(t *testing.T) {
	var rv odb.AuthCodeRevolk
	rv.AuthorizationCode = acIDAcRvi
	res, id := odbAcRvi.AddAuthCodeRevolk(nil, &rv)
	fmt.Println("revolk id: ", id)
	if !res {
		t.Fail()
	}
}
func TestMySQLOauthDBAcRvi_GetAuthCodeRevolk(t *testing.T) {
	rv := odbAcRvi.GetAuthCodeRevolk(acIDAcRvi)
	fmt.Println("revolk : ", rv)
	if rv == nil {
		t.Fail()
	}
}
func TestMySQLOauthDBAcRvi_DeleteAuthCodeRevolk(t *testing.T) {
	res := odbAcRvi.DeleteAuthCodeRevolk(nil, acIDAcRvi)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcRvi_DeleteAuthorizationCode(t *testing.T) {
	res := odbAcRvi.DeleteAuthorizationCode(cidAcRvi, "1234")
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcRvi_DeleteClient(t *testing.T) {
	suc := odbAcRvi.DeleteClient(cidAcRvi)
	if !suc {
		t.Fail()
	}
}
