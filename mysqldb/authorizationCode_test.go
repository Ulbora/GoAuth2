package mysqldb

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
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
	ac.UserID = 1234
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445"

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
	ac.UserID = 1234
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445"
	var scope = []string{"test1", "test2"}

	res, id := odbAci.AddAuthorizationCode(&ac, &at, &rt, &scope)
	if !res || id < 1 {
		t.Fail()
	} else {
		spID2Aci = id
	}
}

func TestMySQLOauthDBACi_DeleteAuthorizationCodeScope2(t *testing.T) {
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
