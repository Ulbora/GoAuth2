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

var dbIgi db.Database
var odbIgi odb.Oauth2DB
var cidIgi int64
var spIDIgi int64
var spID2Igi int64

func TestMySQLOauthDBIgi_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIgi = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbIgi

	odbIgi = &moadb

	dbIgi.Connect()
}

func TestMySQLOauthDBIgi_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbIgi.AddClient(&c, nil)
	fmt.Println("client add res: ", res)
	fmt.Println("client id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidIgi = id
	}
}

func TestMySQLOauthDBIgi_AddImplicitGrantNoScope(t *testing.T) {

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ig odb.ImplicitGrant
	ig.ClientID = cidIgi
	ig.UserID = "1234"
	res, igid := odbIgi.AddImplicitGrant(&ig, &at, nil)
	if !res || igid <= 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBIgi_AddImplicitGrant(t *testing.T) {

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ig odb.ImplicitGrant
	ig.ClientID = cidIgi
	ig.UserID = "1234"
	var scope = []string{"test1", "test2"}
	res, igid := odbIgi.AddImplicitGrant(&ig, &at, &scope)
	if !res || igid <= 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBIgi_GetImplicitGrant(t *testing.T) {
	res := odbIgi.GetImplicitGrant(cidIgi, "1234")
	if len(*res) < 2 {
		t.Fail()
	}
}

func TestMySQLOauthDBIgi_GetImplicitGrantByScope(t *testing.T) {
	res := odbIgi.GetImplicitGrantByScope(cidIgi, "1234", "test1")
	if len(*res) != 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBIgi_DeleteImplicitGrant(t *testing.T) {
	res := odbIgi.DeleteImplicitGrant(cidIgi, "1234")
	if !res {
		t.Fail()
	}

}

func TestMySQLOauthDBIgi_DeleteClient(t *testing.T) {
	suc := odbIgi.DeleteClient(cidIgi)
	if !suc {
		t.Fail()
	}
}
