// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbPgi db.Database
var odbPgi odb.Oauth2DB
var cidPgi int64
var spIDPgi int64
var spID2Pgi int64

func TestMySQLOauthDBPgi_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPgi = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbPgi

	odbPgi = &moadb

	dbPgi.Connect()
}

func TestMySQLOauthDBPgi_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbPgi.AddClient(&c, nil)
	fmt.Println("client add res: ", res)
	fmt.Println("client id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidPgi = id
	}
}

func TestMySQLOauthDBPgi_AddPasswordGrant(t *testing.T) {

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.PasswordGrant
	pwg.ClientID = cidPgi
	pwg.UserID = "1234"
	res, id := odbPgi.AddPasswordGrant(&pwg, &at, &rt)
	if !res || id < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBPgi_DeletePasswordGrant(t *testing.T) {
	res := odbPgi.DeletePasswordGrant(cidPgi, "1234")
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBPgi_DeleteClient(t *testing.T) {
	suc := odbPgi.DeleteClient(cidPgi)
	if !suc {
		t.Fail()
	}
}
