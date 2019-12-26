package mysqldb

import (
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbCg db.Database
var odbCg odb.Oauth2DB
var cidCg int64 = 1
var spIDCg int64
var spID2Cg int64

func TestMySQLOauthDBCg_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbCg = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbCg

	odbCg = &moadb

	dbCg.Connect()
}

func TestMySQLOauthDBCg_AddCredentialsGrant(t *testing.T) {

	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	var moadb MySQLOauthDB
	moadb.DB = dbPg

	odbCg = &moadb

	dbCg.Connect()

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.CredentialsGrant
	pwg.ClientID = cidCg
	res, id := odbCg.AddCredentialsGrant(&pwg, &at)
	if !res || id < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBCg_AddCredentialsGrantFail1(t *testing.T) {

	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	//mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	var moadb MySQLOauthDB
	moadb.DB = dbPg

	odbCg = &moadb

	dbCg.Connect()

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.CredentialsGrant
	pwg.ClientID = cidCg
	res, id := odbCg.AddCredentialsGrant(&pwg, &at)
	if res || id != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBCg_AddCredentialsGrantFail2(t *testing.T) {

	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	//mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	//mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	var moadb MySQLOauthDB
	moadb.DB = dbPg

	odbCg = &moadb

	dbCg.Connect()

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.CredentialsGrant
	pwg.ClientID = cidCg
	res, id := odbCg.AddCredentialsGrant(&pwg, &at)
	if res || id != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBCg_DeleteCredentialsGrant(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	var rows [][]string
	row := []string{"1", "2", "4"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbPg

	odbCg = &moadb

	dbCg.Connect()

	res := odbCg.DeleteCredentialsGrant(cidCg)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBCg_DeleteCredentialsGrant2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	var rows [][]string
	row := []string{}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbPg

	odbCg = &moadb

	dbCg.Connect()

	res := odbCg.DeleteCredentialsGrant(cidCg)
	if !res {
		t.Fail()
	}
}
func TestMySQLOauthDBCg_DeleteCredentialsGrantFail1(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	//mydb.MockDeleteSuccess2 = true

	var rows [][]string
	row := []string{"1", "2", "4"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbPg

	odbCg = &moadb

	dbCg.Connect()

	res := odbCg.DeleteCredentialsGrant(cidCg)
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBCg_DeleteCredentialsGrantFail2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	//mydb.MockDeleteSuccess1 = true

	//mydb.MockDeleteSuccess2 = true

	var rows [][]string
	row := []string{"1", "2", "4"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbPg

	odbCg = &moadb

	dbCg.Connect()

	res := odbCg.DeleteCredentialsGrant(cidCg)
	if res {
		t.Fail()
	}
}
