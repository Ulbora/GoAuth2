package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbCgt db.Database
var odbCgt odb.Oauth2DB
var cidCgt int64 = 1
var idCgt int64

func TestMySQLOauthDBCgt_ConnectClientGrantType(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbCgt = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var rows [][]string
	row1 := []string{"1", "tester5", "1"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	moadb.DB = dbCgt

	odbCgt = &moadb

	dbCgt.Connect()
}

func TestMySQLOauthDBCgt_AddClientGrantType(t *testing.T) {
	var ur odb.ClientGrantType
	ur.ClientID = cidCgt
	ur.GrantType = "someGrantType"
	res, id := odbCgt.AddClientGrantType(&ur)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idCgt = id
	}
}

func TestMySQLOauthDBCgt_GetClientGrantTypeList(t *testing.T) {
	res := odbCgt.GetClientGrantTypeList(cidCgt)
	fmt.Println("grant type list res: ", res)
	if res == nil || (*res)[0].ClientID != cidCgt {
		t.Fail()
	}
}

func TestMySQLOauthDBCgt_DeleteClientGrantType(t *testing.T) {
	res := odbCgt.DeleteClientGrantType(idCgt)
	fmt.Println("client grant type  delete: ", res)
	if !res {
		t.Fail()
	}
}
