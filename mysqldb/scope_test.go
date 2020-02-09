package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbSo db.Database
var odbSo odb.Oauth2DB
var cidSo int64
var idSo int64

func TestMySQLOauthDB_ConnectScope(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbSo = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var rows [][]string
	row1 := []string{"1", "somescope", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbSo

	odbSo = &moadb

	dbSo.Connect()
}

func TestMySQLOauthDB_AddClientScope(t *testing.T) {
	cidSo = 2
	var ur odb.ClientScope
	ur.ClientID = cidSo
	ur.Scope = "somescope"
	res, id := odbSo.AddClientScope(&ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idSo = id
	}
}

func TestMySQLOauthDB_GetClientScopeList(t *testing.T) {
	res := odbSo.GetClientScopeList(cidSo)
	fmt.Println("scope list res: ", res)
	if res == nil || (*res)[0].ClientID != cidSo {
		t.Fail()
	}
}

func TestMySQLOauthDB_DeleteClientScope(t *testing.T) {
	res := odbSo.DeleteClientScope(idSo)
	fmt.Println("scope  delete: ", res)
	if !res {
		t.Fail()
	}
}
