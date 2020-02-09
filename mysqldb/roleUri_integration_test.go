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

var dbroui db.Database
var odbroui odb.Oauth2DB
var cidroui int64

//uil id
var idrouiUidi int64

//role id
var idrouiRoidi int64

var idroui int64

func TestMySQLOauthDBRoleUrii_ConnectAllowURI(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbroui = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbroui

	odbroui = &moadb

	dbroui.Connect()
}

func TestMySQLOauthDBRoleUrii_AddClient(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbroui.AddClient(&c, nil)
	fmt.Println("res: ", res)
	fmt.Println("id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidroui = id
	}
}

func TestMySQLOauthDBRoleUrii_AddClientAllowedURI(t *testing.T) {
	var ur odb.ClientAllowedURI
	ur.ClientID = cidroui
	ur.URI = "someuri"
	res, id := odbroui.AddClientAllowedURI(&ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idrouiUidi = id
	}
}

func TestMySQLOauthDBRoleUrii_AddClientRole(t *testing.T) {
	var r odb.ClientRole
	r.ClientID = cidroui
	r.Role = "someRole"
	res, id := odbroui.AddClientRole(&r)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idrouiRoidi = id
	}
}
func TestMySQLOauthDBRoleUrii_AddClientRoleURI(t *testing.T) {
	var r odb.ClientRoleURI
	r.ClientAllowedURIID = idrouiUidi
	r.ClientRoleID = idrouiRoidi
	res := odbroui.AddClientRoleURI(&r)
	// fmt.Println("role uri id: ", id)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUrii_GetClientRoleURIList(t *testing.T) {
	res := odbroui.GetClientRoleAllowedURIList(idrouiRoidi)
	fmt.Println("Role URI list res: ", res)
	if res == nil || (*res)[0].ClientRoleID != idrouiRoidi {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUrii_GetClientRoleURIListByClient(t *testing.T) {
	res := odbroui.GetClientRoleAllowedURIListByClientID(cidroui)
	fmt.Println("Role URI list by client res: ", res)
	if res == nil || (*res)[0].ClientRoleID != idrouiRoidi {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUrii_DeleteClientRoleURI(t *testing.T) {
	var r odb.ClientRoleURI
	r.ClientAllowedURIID = idrouiUidi
	r.ClientRoleID = idrouiRoidi
	res := odbroui.DeleteClientRoleURI(&r)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUrii_DeleteClientAllowedURI(t *testing.T) {
	res := odbroui.DeleteClientAllowedURI(idrouiUidi)
	fmt.Println("allowed uri  delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUrii_DeleteClientRole(t *testing.T) {
	res := odbroui.DeleteClientRole(idrouiRoidi)
	fmt.Println("role  delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUrii_DeleteClient(t *testing.T) {
	suc := odbroui.DeleteClient(cidroui)
	if !suc {
		t.Fail()
	}
}
