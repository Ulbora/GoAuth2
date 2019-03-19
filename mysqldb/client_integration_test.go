package mysqldb

// +bbuild integration move to top

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbb db.Database
var odbb odb.Oauth2DB
var cid int64
var cid2 int64

func TestMySQLDB_Connect(t *testing.T) {

	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbb = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbb

	odbb = &moadb

	dbb.Connect()

}

func TestMySQLDB_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbb.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cid = id
	}
}

func TestMySQLDB_AddClient(t *testing.T) {
	var c odb.Client
	c.Secret = "1234567"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false
	var uis []odb.ClientRedirectURI
	var u1 odb.ClientRedirectURI
	u1.URI = "addSomething"
	uis = append(uis, u1)

	var u2 odb.ClientRedirectURI
	u2.URI = "addSomething2"
	uis = append(uis, u2)

	fmt.Println("before db add")
	res, id := odbb.AddClient(&c, &uis)
	cid2 = id
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	}
}

func TestMySQLDB_UpdateClient(t *testing.T) {
	var c odb.Client
	c.Secret = "555555"
	c.Name = "tester5"
	c.Email = "bob5@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = false
	c.Paid = false
	c.ClientID = cid
	suc := odbb.UpdateClient(&c)
	if !suc {
		t.Fail()
	}
}

func TestMySQLDB_DeleteClient(t *testing.T) {
	suc := odbb.DeleteClient(cid)
	if !suc {
		t.Fail()
	}
}

func TestMySQLDB_DeleteClient2(t *testing.T) {
	suc := odbb.DeleteClient(cid2)
	if !suc {
		t.Fail()
	}
}
