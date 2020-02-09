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

var dbRli db.Database
var odbRli odb.Oauth2DB
var cidRli int64
var idRli int64

func TestMySQLOauthDBi_ConnectRole(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbRli = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbRli

	odbRli = &moadb

	dbRli.Connect()
}

func TestMySQLOauthDBi_AddClientInRole(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbRli.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidRli = id
	}
}

func TestMySQLOauthDBi_AddClientRole(t *testing.T) {
	var r odb.ClientRole
	r.ClientID = cidRli
	r.Role = "someRole"
	res, id := odbRli.AddClientRole(&r)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idRli = id
	}
}

func TestMySQLOauthDBi_GetClientRoleList(t *testing.T) {
	res := odbRli.GetClientRoleList(cidRli)
	fmt.Println("Role list res: ", res)
	if res == nil || (*res)[0].ClientID != cidRli {
		t.Fail()
	}
}

func TestMySQLOauthDBi_DeleteClientRole(t *testing.T) {
	res := odbRli.DeleteClientRole(idRli)
	fmt.Println("role  delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBi_DeleteClientInRole(t *testing.T) {
	suc := odbRli.DeleteClient(cidRli)
	if !suc {
		t.Fail()
	}
}
