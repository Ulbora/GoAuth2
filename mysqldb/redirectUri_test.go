package mysqldb

import (
	//"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbbUri2 db.Database
var odbbUri2 odb.Oauth2DB
var rdid2 int64
var cidUri2 int64

func TestMySQLOauthDB2_Connect(t *testing.T) {

	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbbUri2 = &mydb

	//mydb.MockTestRow
	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow
	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var getRow db.DbRow
	getRow.Row = []string{"1", "test", "2"}
	mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{"1", "tester5", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true

	var moadb MySQLOauthDB
	moadb.DB = dbbUri2

	odbbUri2 = &moadb

	dbbUri2.Connect()

}

func TestMySQLOauthDB2_AddClientRedirectURI(t *testing.T) {
	var ur odb.ClientRedirectURI
	ur.ClientID = 4
	ur.URI = "someuri"
	res, id := odbbUri2.AddClientRedirectURI(nil, &ur)
	if !res || id <= 0 {
		t.Fail()
	}
}

func TestMySQLOauthDB2_AddClientRedirectURITx(t *testing.T) {
	var ur odb.ClientRedirectURI
	ur.ClientID = 4
	ur.URI = "someuri"

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockInsertSuccess1 = true
	mdbx.MockInsertID1 = 1
	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res, id := odbbUri2TX.AddClientRedirectURI(&mtx, &ur)
	if !res || id <= 0 {
		t.Fail()
	}
}

func TestMySQLOauthDB2_GetClientRedirectURI(t *testing.T) {
	var cid int64 = 2
	res := odbbUri2.GetClientRedirectURI(cid, "someuri")
	if res == nil {
		t.Fail()
	}
}

func TestMySQLOauthDB2_GetClientRedirectURIList(t *testing.T) {
	var cid int64 = 2
	res := odbbUri2.GetClientRedirectURIList(cid)
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestMySQLOauthDB2_DeleteClientRedirectURI(t *testing.T) {
	var id int64 = 2
	res := odbbUri2.DeleteClientRedirectURI(nil, id)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDB2_DeleteClientRedirectURIAll(t *testing.T) {
	var id int64 = 2
	res := odbbUri2.DeleteClientAllRedirectURI(nil, id)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDB2_DeleteClientRedirectURITx(t *testing.T) {
	var id int64 = 2

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockDeleteSuccess1 = true
	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res := odbbUri2TX.DeleteClientRedirectURI(&mtx, id)
	if !res {
		t.Fail()
	}
}
