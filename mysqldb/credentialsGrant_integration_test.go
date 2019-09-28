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

var dbCgi db.Database
var odbCgi odb.Oauth2DB
var cidCgi int64
var spIDCgi int64
var spID2Cgi int64

func TestMySQLOauthDBCgi_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbCgi = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbCgi

	odbCgi = &moadb

	dbCgi.Connect()
}

func TestMySQLOauthDBCgi_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbCgi.AddClient(&c, nil)
	fmt.Println("client add res: ", res)
	fmt.Println("client id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidCgi = id
	}
}

func TestMySQLOauthDBCgi_AddCredentialsGrant(t *testing.T) {

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.CredentialsGrant
	pwg.ClientID = cidCgi
	res, id := odbCgi.AddCredentialsGrant(&pwg, &at)
	if !res || id < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBCgi_DeleteCredentialsGrant(t *testing.T) {
	res := odbCgi.DeleteCredentialsGrant(cidCgi)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBCgi_DeleteClient(t *testing.T) {
	suc := odbCgi.DeleteClient(cidCgi)
	if !suc {
		t.Fail()
	}
}
