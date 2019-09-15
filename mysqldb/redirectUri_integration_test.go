// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbbUri db.Database
var odbbUri odb.Oauth2DB
var rdidi int64
var cidUri int64

func TestMySQLOauthDB_Connect(t *testing.T) {

	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbbUri = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbbUri

	odbbUri = &moadb

	dbbUri.Connect()

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
	res, id := odbbUri.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidUri = id
	}
}

func TestMySQLOauthDB_AddClientRedirectURI(t *testing.T) {
	var ur odb.ClientRedirectURI
	ur.ClientID = cidUri
	ur.URI = "someuri"
	res, id := odbbUri.AddClientRedirectURI(nil, &ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		rdidi = id
	}
}

func TestMySQLOauthDB_GetClientRedirectURIList(t *testing.T) {
	res := odbbUri.GetClientRedirectURIList(cidUri)
	fmt.Println("uri res: ", res)
	if res == nil || (*res)[0].ClientID != cidUri {
		t.Fail()
	}
}

func TestMySQLOauthDB_GetClientRedirectURI(t *testing.T) {
	res := odbbUri.GetClientRedirectURI(cidUri, "someuri")
	fmt.Println("uri res by id: ", res)
	if res == nil || (*res).ClientID != cidUri {
		t.Fail()
	}
}

func TestMySQLOauthDB_DeleteClientRedirectURI(t *testing.T) {
	res := odbbUri.DeleteClientRedirectURI(nil, rdidi)
	fmt.Println("uri  delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDB_DeleteClient(t *testing.T) {
	suc := odbbUri.DeleteClient(cidUri)
	if !suc {
		t.Fail()
	}
}
