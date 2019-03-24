package mysqldb

import (
	"fmt"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"testing"
)

var dbRt db.Database
var odbRt odb.Oauth2DB
var idRt int64

//var cidRti int64

func TestMySQLOauthDBReToken_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbRt = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var getRow db.DbRow
	getRow.Row = []string{"1", "somereftoken2"}
	mydb.MockRow1 = &getRow

	mydb.MockUpdateSuccess1 = true

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	moadb.DB = dbRt

	odbRt = &moadb

	dbRt.Connect()
}

func TestMySQLOauthDBReToken_AddRefreshToken(t *testing.T) {
	var tk odb.RefreshToken
	tk.Token = "somereftoken"
	res, id := odbRt.AddRefreshToken(&tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idRt = id
	}
}

func TestMySQLOauthDBReToken_UpdateRefreshToken(t *testing.T) {
	var tk odb.RefreshToken
	tk.ID = idRt
	tk.Token = "somereftoken2"
	res := odbRt.UpdateRefreshToken(&tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBReToken_GetRefreshToken(t *testing.T) {
	res := odbRt.GetRefreshToken(idRt)
	fmt.Println("ref token: ", res)
	if res == nil || (*res).Token != "somereftoken2" {
		t.Fail()
	}
}

func TestMySQLOauthDBReToken_DeleteRefreshToken(t *testing.T) {
	res := odbRt.DeleteRefreshToken(idRt)
	fmt.Println("del ref token: ", res)
	if !res {
		t.Fail()
	}
}
