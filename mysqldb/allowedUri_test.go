package mysqldb

import (
	"fmt"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"testing"
)

var dbAu db.Database
var odbAu odb.Oauth2DB
var cidAu int64
var idAu int64

func TestMySQLOauthDB_ConnectAllowURI(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockUpdateSuccess1 = true

	var getRow db.DbRow
	getRow.Row = []string{"1", "someuri2", "2"}
	mydb.MockRow1 = &getRow
	mydb.MockRow2 = &getRow

	var rows [][]string
	row1 := []string{"1", "tester5", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	dbAu.Connect()
}

func TestMySQLOauthDB_AddClientAllowedURI(t *testing.T) {
	cidAu = 2
	var ur odb.ClientAllowedURI
	ur.ClientID = cidAu
	ur.URI = "someuri"
	res, id := odbAu.AddClientAllowedURI(&ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAu = id
	}
}

func TestMySQLOauthDB_UpdateClientAllowedURI(t *testing.T) {
	var ur odb.ClientAllowedURI
	ur.ID = idAu
	ur.URI = "someuri2"
	res := odbAu.UpdateClientAllowedURI(&ur)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDB_GetClientAllowedURIByID(t *testing.T) {
	res := odbAu.GetClientAllowedURIByID(idAu)
	fmt.Println("allowed uri res by id: ", res)
	if res == nil || (*res).ClientID != cidAu || (*res).URI != "someuri2" {
		t.Fail()
	}
}

func TestMySQLOauthDB_GetClientAllowedURIList(t *testing.T) {
	res := odbAu.GetClientAllowedURIList(cidAu)
	fmt.Println("allowed uri list res: ", res)
	if res == nil || (*res)[0].ClientID != cidAu {
		t.Fail()
	}
}

func TestMySQLOauthDB_GetClientAllowedURI(t *testing.T) {
	res := odbAu.GetClientAllowedURI(cidAu, "someuri2")
	fmt.Println("allowed uri res: ", res)
	if res == nil || (*res).ClientID != cidAu || (*res).URI != "someuri2" {
		t.Fail()
	}
}

func TestMySQLOauthDB_DeleteClientAllowedURI(t *testing.T) {
	res := odbAu.DeleteClientAllowedURI(idAu)
	fmt.Println("allowed uri  delete: ", res)
	if !res {
		t.Fail()
	}
}
