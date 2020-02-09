package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbRl db.Database
var odbRl odb.Oauth2DB
var cidRl int64
var idRl int64

func TestMySQLOauthDB_ConnectRole(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbRl = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var rows [][]string
	row1 := []string{"1", "tester5", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbRl

	odbRl = &moadb

	dbRl.Connect()
}

func TestMySQLOauthDB_AddClientRole(t *testing.T) {
	cidRl = 2
	var r odb.ClientRole
	r.ClientID = cidRl
	r.Role = "someRole"
	res, id := odbRl.AddClientRole(&r)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idRl = id
	}
}

func TestMySQLOauthDB_GetClientRoleList(t *testing.T) {
	res := odbRl.GetClientRoleList(cidRl)
	fmt.Println("Role list res: ", res)
	if res == nil || (*res)[0].ClientID != cidRl {
		t.Fail()
	}
}

func TestMySQLOauthDB_DeleteClientRole(t *testing.T) {
	res := odbRl.DeleteClientRole(idRl)
	fmt.Println("role  delete: ", res)
	if !res {
		t.Fail()
	}
}
