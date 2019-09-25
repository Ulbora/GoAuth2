// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbCgti db.Database
var odbCgti odb.Oauth2DB
var cidCgti int64
var idCgti int64

func TestMySQLOauthDBCgti_ConnectClientGrantType(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbCgti = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbCgti

	odbCgti = &moadb

	dbCgti.Connect()
}

func TestMySQLOauthDBCgti_AddClientInClientGrantType(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbCgti.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidCgti = id
	}
}
func TestMySQLOauthDBCgti_AddClientGrantType(t *testing.T) {
	var ur odb.ClientGrantType
	ur.ClientID = cidCgti
	ur.GrantType = "someGrantType"
	res, id := odbCgti.AddClientGrantType(&ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idCgti = id
	}
}

func TestMySQLOauthDBCgti_GetClientGrantTypeList(t *testing.T) {
	res := odbCgti.GetClientGrantTypeList(cidCgti)
	fmt.Println("grant type list res: ", res)
	if res == nil || (*res)[0].ClientID != cidCgti {
		t.Fail()
	}
}

func TestMySQLOauthDBCgti_DeleteClientGrantType(t *testing.T) {
	res := odbCgti.DeleteClientGrantType(idCgti)
	fmt.Println("client grant type  delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBCgti_DeleteClientInClientGrantType(t *testing.T) {
	suc := odbCgti.DeleteClient(cidCgti)
	if !suc {
		t.Fail()
	}
}
