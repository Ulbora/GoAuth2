// +build integration move to top

package mysqldb

import (
	"fmt"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"testing"
)

var dbSoi db.Database
var odbSoi odb.Oauth2DB
var cidSoi int64
var idSoi int64

func TestMySQLOauthDBi_ConnectScope(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbSoi = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbSoi

	odbSoi = &moadb

	dbSoi.Connect()
}

func TestMySQLOauthDBi_AddClientInScope(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbSoi.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidSoi = id
	}
}

func TestMySQLOauthDBi_AddClientScope(t *testing.T) {
	var ur odb.ClientScope
	ur.ClientID = cidSoi
	ur.Scope = "somescope"
	res, id := odbSoi.AddClientScope(&ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idSoi = id
	}
}

func TestMySQLOauthDBi_GetClientScopeList(t *testing.T) {
	res := odbSoi.GetClientScopeList(cidSoi)
	fmt.Println("scope list res: ", res)
	if res == nil || (*res)[0].ClientID != cidSoi {
		t.Fail()
	}
}

func TestMySQLOauthDBi_DeleteClientScope(t *testing.T) {
	res := odbSoi.DeleteClientScope(idSoi)
	fmt.Println("scope  delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBi_DeleteClientInScope(t *testing.T) {
	suc := odbSoi.DeleteClient(cidSoi)
	if !suc {
		t.Fail()
	}
}
