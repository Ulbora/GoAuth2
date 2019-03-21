// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbb_uri db.Database
var odbb_uri odb.Oauth2DB
var rdid int64
var cid_uri int64

func TestMySQLOauthDB_Connect(t *testing.T) {

	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbb_uri = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbb_uri

	odbb_uri = &moadb

	dbb_uri.Connect()

}

func TestMySQLOauthDB_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbb_uri.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cid_uri = id
	}
}
func TestMySQLOauthDB_AddClientRedirectURI(t *testing.T) {
	var ur odb.ClientRedirectURI
	ur.ClientID = cid_uri
	ur.URI = "someuri"
	res, id := odbb_uri.AddClientRedirectURI(&ur)
	if !res || id <= 0 {
		t.Fail()
	}
}

func TestMySQLOauthDB_DeleteClient(t *testing.T) {
	suc := odbb_uri.DeleteClient(cid_uri)
	if !suc {
		t.Fail()
	}
}
