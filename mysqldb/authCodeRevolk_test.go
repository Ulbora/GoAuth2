package mysqldb

import (
	"fmt"
	"testing"

	//"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbAcRv db.Database
var odbAcRv odb.Oauth2DB
var cidAcRv int64 = 1
var acIDAcRv int64 = 2

//var spID2AcRv int64

func TestMySQLOauthDBAcRv_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAcRv = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var getRow db.DbRow
	getRow.Row = []string{"1", "2"}
	mydb.MockRow1 = &getRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbAcRv

	odbAcRv = &moadb

	dbAcRv.Connect()
}

// func TestMySQLOauthDBAcRv_AddClientNullUri(t *testing.T) {
// 	var c odb.Client
// 	c.Secret = "12345"
// 	c.Name = "tester"
// 	c.Email = "bob@bob.com"
// 	c.WebSite = "www.bob.com"
// 	c.Enabled = true
// 	c.Paid = false

// 	fmt.Println("before db add")
// 	res, id := odbAcRv.AddClient(&c, nil)
// 	fmt.Println("client add res: ", res)
// 	fmt.Println("client id: ", id)
// 	if !res || id == 0 {
// 		t.Fail()
// 	} else {
// 		cidAcRv = id
// 	}
// }

// func TestMySQLOauthDBAcRv_AddAuthorizationCode(t *testing.T) {

// 	var rt odb.RefreshToken
// 	rt.Token = "somereftoken2"

// 	var at odb.AccessToken
// 	at.Token = "someacctoken"
// 	at.Expires = time.Now()

// 	var ac odb.AuthorizationCode
// 	ac.ClientID = cidAcRv
// 	ac.UserID = "1234"
// 	ac.Expires = time.Now()
// 	ac.RandonAuthCode = "13445"

// 	res, id := odbAcRv.AddAuthorizationCode(&ac, &at, &rt, nil)

// 	if !res || id < 1 {
// 		t.Fail()
// 	} else {
// 		acIDAcRv = id
// 	}
// }

func TestMySQLOauthDBAcRv_AddAuthCodeRevolk(t *testing.T) {
	var rv odb.AuthCodeRevolk
	rv.AuthorizationCode = acIDAcRv
	res, id := odbAcRv.AddAuthCodeRevolk(nil, &rv)
	fmt.Println("revolk id: ", id)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcRv_AddAuthCodeRevolkTx(t *testing.T) {

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockInsertSuccess1 = true
	mdbx.MockInsertID1 = 1

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mdbx.MockTestRow = &mTestRow

	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	var l lg.Logger
	moadbtx.Log = &l
	//moadbtx.DB = &mtx
	var odbbUri2TX = &moadbtx

	var rv odb.AuthCodeRevolk
	rv.AuthorizationCode = acIDAcRv
	res, id := odbbUri2TX.AddAuthCodeRevolk(&mtx, &rv)
	fmt.Println("revolk id: ", id)
	if !res {
		t.Fail()
	}
}
func TestMySQLOauthDBAcRv_GetAuthCodeRevolk(t *testing.T) {
	rv := odbAcRv.GetAuthCodeRevolk(acIDAcRv)
	fmt.Println("revolk : ", rv)
	if rv == nil {
		t.Fail()
	}
}
func TestMySQLOauthDBAcRv_DeleteAuthCodeRevolk(t *testing.T) {
	res := odbAcRv.DeleteAuthCodeRevolk(nil, acIDAcRv)
	if !res {
		t.Fail()
	}
}

// func TestMySQLOauthDBAcRv_DeleteAuthorizationCode(t *testing.T) {
// 	res := odbAcRv.DeleteAuthorizationCode(cidAcRv, "1234")
// 	if !res {
// 		t.Fail()
// 	}
// }

// func TestMySQLOauthDBAcRv_DeleteClient(t *testing.T) {
// 	suc := odbAcRv.DeleteClient(cidAcRv)
// 	if !suc {
// 		t.Fail()
// 	}
// }
