// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbbi db.Database
var odbbi odb.Oauth2DB
var cidi int64
var cid2i int64

func TestMySQLDBi_Connect(t *testing.T) {

	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbbi = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbbi

	odbbi = &moadb

	dbbi.Connect()

}

func TestMySQLDBi_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbbi.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidi = id
	}
}

func TestMySQLDBi_AddClient(t *testing.T) {
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
	res, id := odbbi.AddClient(&c, &uis)
	cid2i = id
	fmt.Println("res: ", res)
	fmt.Println("id in addclient int test: ", id)
	if !res || id == 0 {
		t.Fail()
	}
}

func TestMySQLDBi_UpdateClient(t *testing.T) {
	var c odb.Client
	c.Secret = "555555"
	c.Name = "tester5"
	c.Email = "bob5@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = false
	c.Paid = false
	c.ClientID = cidi
	suc := odbbi.UpdateClient(&c)
	if !suc {
		t.Fail()
	}
}

func TestMySQLDBi_GetClient(t *testing.T) {
	c := odbbi.GetClient(cidi)
	fmt.Println("client found: ", c)
	if c.Name != "tester5" {
		t.Fail()
	}
}

func TestMySQLDBi_GetClients(t *testing.T) {
	cs := odbbi.GetClients()
	fmt.Println("client found: ", cs)
	for _, c := range *cs {
		fmt.Println("client found in getClients: ", c)
	}
	if len(*cs) == 0 {
		t.Fail()
	}

}

func TestMySQLDBi_SearchClients(t *testing.T) {
	cs := odbbi.SearchClients("test")
	fmt.Println("client found in search: ", cs)
	for _, c := range *cs {
		fmt.Println("client found in searchClients: ", c)
	}
	if len(*cs) == 0 {
		t.Fail()
	}

}

func TestMySQLDBi_DeleteClient(t *testing.T) {
	suc := odbbi.DeleteClient(cidi)
	if !suc {
		t.Fail()
	}
}

func TestMySQLDBi_DeleteClient2(t *testing.T) {
	suc := odbbi.DeleteClient(cid2i)
	if !suc {
		t.Fail()
	}
}
