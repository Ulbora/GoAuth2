package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbrou db.Database
var odbrou odb.Oauth2DB
var cidrou int64

//uil id
var idrouiUid int64

//role id
var idrouiRoid int64

var idrou int64

func TestMySQLOauthDBRoleUri_ConnectAllowURI(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbrou = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow
	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var rows [][]string
	row1 := []string{"4", "1"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rowsb [][]string
	row1b := []string{"4", "somerole", "1", "someurl", "2"}
	rowsb = append(rowsb, row1b)
	var dbrowsb db.DbRows
	dbrowsb.Rows = rowsb
	mydb.MockRows2 = &dbrowsb

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbrou

	odbrou = &moadb

	dbrou.Connect()
}

func TestMySQLOauthDBRoleUri_AddClientRoleURI(t *testing.T) {
	cidrou = 2
	idrouiUid = 1
	idrouiRoid = 4

	var r odb.ClientRoleURI
	r.ClientAllowedURIID = idrouiUid
	r.ClientRoleID = idrouiRoid
	res := odbrou.AddClientRoleURI(&r)
	// fmt.Println("role uri id: ", id)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUri_GetClientRoleURIList(t *testing.T) {
	res := odbrou.GetClientRoleAllowedURIList(idrouiRoid)
	fmt.Println("Role URI list res: ", res)
	fmt.Println("Role URI list res.ClientRoleID: ", (*res)[0].ClientRoleID)
	if res == nil || (*res)[0].ClientRoleID != idrouiRoid {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUri_GetClientRoleURIListByClient(t *testing.T) {
	res := odbrou.GetClientRoleAllowedURIListByClientID(cidrou)
	fmt.Println("Role URI list by client res: ", res)
	if res == nil || (*res)[0].ClientRoleID != idrouiRoid {
		t.Fail()
	}
}

func TestMySQLOauthDBRoleUri_DeleteClientRoleURI(t *testing.T) {
	var r odb.ClientRoleURI
	r.ClientAllowedURIID = idrouiUid
	r.ClientRoleID = idrouiRoid
	res := odbrou.DeleteClientRoleURI(&r)
	if !res {
		t.Fail()
	}
}
