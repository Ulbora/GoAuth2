package mysqldb

import (
	//"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbb_uri2 db.Database
var odbb_uri2 odb.Oauth2DB
var rdid2 int64
var cid_uri2 int64

func TestMySQLOauthDB2_Connect(t *testing.T) {

	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbb_uri2 = &mydb

	//mydb.MockTestRow
	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow
	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var moadb MySQLOauthDB
	moadb.DB = dbb_uri2

	odbb_uri2 = &moadb

	dbb_uri2.Connect()

}

// func TestMySQLOauthDB2_AddClientNullUri(t *testing.T) {
// 	var c odb.Client
// 	c.Secret = "12345"
// 	c.Name = "tester"
// 	c.Email = "bob@bob.com"
// 	c.WebSite = "www.bob.com"
// 	c.Enabled = true
// 	c.Paid = false

// 	fmt.Println("before db add")
// 	res, id := odbb_uri2.AddClient(&c, nil)
// 	fmt.Println("res: ", res)
// 	fmt.Println("id: ", id)
// 	if !res || id == 0 {
// 		t.Fail()
// 	} else {
// 		cid_uri2 = id
// 	}
// }
func TestMySQLOauthDB2_AddClientRedirectURI(t *testing.T) {
	var ur odb.ClientRedirectURI
	ur.ClientID = 4
	ur.URI = "someuri"
	res, id := odbb_uri2.AddClientRedirectURI(&ur)
	if !res || id <= 0 {
		t.Fail()
	}
}

// func TestMySQLOauthDB2_DeleteClient(t *testing.T) {
// 	suc := odbb_uri2.DeleteClient(cid_uri)
// 	if !suc {
// 		t.Fail()
// 	}
// }
