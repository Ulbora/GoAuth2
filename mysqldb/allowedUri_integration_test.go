// +build integration move to top

package mysqldb

import (
	"fmt"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"testing"
)

var dbAui db.Database
var odbAui odb.Oauth2DB
var cidAui int64
var idAui int64

func TestMySQLOauthDBi_ConnectAllowURI(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAui = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbAui

	odbAui = &moadb

	dbAui.Connect()
}

func TestMySQLOauthDBi_AddClientInAlowedUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbAui.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidAui = id
	}
}

func TestMySQLOauthDBi_AddClientAllowedURI(t *testing.T) {
	var ur odb.ClientAllowedURI
	ur.ClientID = cidAui
	ur.URI = "someuri"
	res, id := odbAui.AddClientAllowedURI(&ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAui = id
	}
}

func TestMySQLOauthDBi_UpdateClientAllowedURI(t *testing.T) {
	var ur odb.ClientAllowedURI
	ur.ID = idAui
	ur.URI = "someuri2"
	res := odbAui.UpdateClientAllowedURI(&ur)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBi_GetClientAllowedURIByID(t *testing.T) {
	res := odbAui.GetClientAllowedURIByID(idAui)
	fmt.Println("allowed uri res by id: ", res)
	if res == nil || (*res).ClientID != cidAui || (*res).URI != "someuri2" {
		t.Fail()
	}
}

func TestMySQLOauthDBi_GetClientAllowedURIList(t *testing.T) {
	res := odbAui.GetClientAllowedURIList(cidAui)
	fmt.Println("allowed uri list res: ", res)
	if res == nil || (*res)[0].ClientID != cidAui {
		t.Fail()
	}
}

func TestMySQLOauthDBi_GetClientAllowedURI(t *testing.T) {
	res := odbAui.GetClientAllowedURI(cidAui, "someuri2")
	fmt.Println("allowed uri res: ", res)
	if res == nil || (*res).ClientID != cidAui || (*res).URI != "someuri2" {
		t.Fail()
	}
}
func TestMySQLOauthDBi_DeleteClientAllowedURI(t *testing.T) {
	res := odbAui.DeleteClientAllowedURI(idAui)
	fmt.Println("allowed uri  delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBi_DeleteClientInAllowedURI(t *testing.T) {
	suc := odbAui.DeleteClient(cidAui)
	if !suc {
		t.Fail()
	}
}
