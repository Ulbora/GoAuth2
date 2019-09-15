package mysqldb

// +bbuild integration move to top

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbAt db.Database
var odbAt odb.Oauth2DB
var idAt int64
var refTkId int64

//var cidRti int64

func TestMySQLOauthDBAcToken_Connect(t *testing.T) {
	refTkId = 5
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAt = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 2

	mydb.MockUpdateSuccess1 = true
	mydb.MockUpdateSuccess2 = true

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true

	var nowTime = time.Now().Format(odb.TimeFormat)

	fmt.Println("time now: ", nowTime)

	var getRow db.DbRow
	getRow.Row = []string{"1", "someacctoken2", nowTime, ""}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow2 = &getRow2

	var moadb MySQLOauthDB
	moadb.DB = dbAt

	odbAt = &moadb

	dbAt.Connect()
}

func TestMySQLOauthDBAcToken_AddAccessToken(t *testing.T) {
	var tk odb.AccessToken
	tk.Token = "someacctoken"
	tk.Expires = time.Now()
	res, id := odbAt.AddAccessToken(nil, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAt = id
	}
}

func TestMySQLOauthDBAcToken_AddAccessTokenTx(t *testing.T) {
	var tk odb.AccessToken
	tk.Token = "someacctoken"
	tk.Expires = time.Now()

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockInsertSuccess1 = true
	mdbx.MockInsertID1 = 1
	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res, id := odbbUri2TX.AddAccessToken(&mtx, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAt = id
	}
}

func TestMySQLOauthDBAcToken_UpdateAccessToken(t *testing.T) {
	var tk odb.AccessToken
	tk.ID = idAt
	tk.Token = "someacctoken2"
	tk.Expires = time.Now()
	res := odbAt.UpdateAccessToken(nil, &tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcToken_UpdateAccessTokenTx(t *testing.T) {
	var tk odb.AccessToken
	tk.ID = idAt
	tk.Token = "someacctoken2"
	tk.Expires = time.Now()

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockUpdateSuccess1 = true

	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res := odbbUri2TX.UpdateAccessToken(&mtx, &tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcToken_GetAccessToken(t *testing.T) {
	res := odbAt.GetAccessToken(idAt)
	fmt.Println("access token: ", res)
	fmt.Println("access token refTokenId: ", res.RefreshTokenID)
	if res == nil || (*res).Token != "someacctoken2" {
		t.Fail()
	}
}

func TestMySQLOauthDBAcToken_DeleteAccessToken(t *testing.T) {
	res := odbAt.DeleteAccessToken(nil, idAt)
	fmt.Println("del access token: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcToken_DeleteAccessTokenTx(t *testing.T) {

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockDeleteSuccess1 = true

	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res := odbbUri2TX.DeleteAccessToken(&mtx, idAt)
	fmt.Println("del access token: ", res)
	if !res {
		t.Fail()
	}
}

// func TestMySQLOauthDBAcToken_AddRefreshToken(t *testing.T) {
// 	var tk odb.RefreshToken
// 	tk.Token = "somereftoken"
// 	res, id := odbAt.AddRefreshToken(&tk)
// 	if !res || id <= 0 {
// 		t.Fail()
// 	} else {
// 		refTkIdi = id
// 	}
// }

func TestMySQLOauthDBAcToken_AddAccessToken2(t *testing.T) {
	var tk odb.AccessToken
	tk.Token = "someacctoken"
	tk.Expires = time.Now()
	tk.RefreshTokenID = refTkId
	res, id := odbAt.AddAccessToken(nil, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAt = id
	}
}

func TestMySQLOauthDBAcToken_AddAccessToken2Tx(t *testing.T) {
	var tk odb.AccessToken
	tk.Token = "someacctoken"
	tk.Expires = time.Now()
	tk.RefreshTokenID = refTkId

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockInsertSuccess1 = true
	mdbx.MockInsertID1 = 1
	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res, id := odbbUri2TX.AddAccessToken(&mtx, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAt = id
	}
}

func TestMySQLOauthDBAcToken_UpdateAccessToken2(t *testing.T) {
	var tk odb.AccessToken
	tk.ID = idAt
	tk.Token = "someacctoken2"
	tk.Expires = time.Now()
	tk.RefreshTokenID = refTkId
	res := odbAt.UpdateAccessToken(nil, &tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcToken_UpdateAccessToken2Tx(t *testing.T) {
	var tk odb.AccessToken
	tk.ID = idAt
	tk.Token = "someacctoken2"
	tk.Expires = time.Now()
	tk.RefreshTokenID = refTkId

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockUpdateSuccess1 = true
	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res := odbbUri2TX.UpdateAccessToken(&mtx, &tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcToken_GetAccessToken2(t *testing.T) {
	res := odbAt.GetAccessToken(idAt)
	fmt.Println("access token: ", res)
	fmt.Println("access token refTokenId: ", res.RefreshTokenID)
	if res == nil || (*res).Token != "someacctoken2" {
		t.Fail()
	}
}

func TestMySQLOauthDBAcToken_DeleteAccessToken2(t *testing.T) {
	res := odbAt.DeleteAccessToken(nil, idAt)
	fmt.Println("del access token: ", res)
	if !res {
		t.Fail()
	}
}

// func TestMySQLOauthDBAcToken_DeleteRefreshToken(t *testing.T) {
// 	res := odbAt.DeleteRefreshToken(refTkIdi)
// 	fmt.Println("del ref token: ", res)
// 	if !res {
// 		t.Fail()
// 	}
// }
