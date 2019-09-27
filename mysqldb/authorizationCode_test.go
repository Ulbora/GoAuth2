package mysqldb

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbAc db.Database
var odbAc odb.Oauth2DB
var cidAc int64
var spIDAc int64
var spID2Ac int64

func TestMySQLOauthDBAC_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	var getRow db.DbRow
	getRow.Row = []string{"1", "somereftoken2"}
	mydb.MockRow1 = &getRow

	mydb.MockUpdateSuccess1 = true

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	dbAc.Connect()
}

func TestMySQLOauthDBAC_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbAc.AddClient(&c, nil)
	fmt.Println("client add res: ", res)
	fmt.Println("client id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidAc = id
	}
}

func TestMySQLOauthDBAC_AddAuthorizationCode(t *testing.T) {

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAc
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445a"

	res, id := odbAc.AddAuthorizationCode(&ac, &at, &rt, nil)
	if !res || id < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_AddAuthorizationCodeNoRefresh(t *testing.T) {

	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	var getRow db.DbRow
	getRow.Row = []string{"1", "somereftoken2"}
	mydb.MockRow1 = &getRow

	mydb.MockUpdateSuccess1 = true

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	var rt odb.RefreshToken
	rt.Token = ""

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAc
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445a"

	res, id := odbAc.AddAuthorizationCode(&ac, &at, &rt, nil)
	if !res || id < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_AddAuthorizationCodeAuthCodeFail(t *testing.T) {

	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = false
	mydb.MockInsertID2 = 0

	mydb.MockInsertSuccess3 = false
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	var getRow db.DbRow
	getRow.Row = []string{"1", "somereftoken2"}
	mydb.MockRow1 = &getRow

	mydb.MockUpdateSuccess1 = true

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	var rt odb.RefreshToken
	rt.Token = ""

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAc
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445a"

	res, id := odbAc.AddAuthorizationCode(&ac, &at, &rt, nil)
	if res || id != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_AddAuthorizationCodeRefTokenFail(t *testing.T) {

	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = false
	mydb.MockInsertID1 = 0

	mydb.MockInsertSuccess2 = false
	mydb.MockInsertID2 = 0

	mydb.MockInsertSuccess3 = false
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	var getRow db.DbRow
	getRow.Row = []string{"1", "somereftoken2"}
	mydb.MockRow1 = &getRow

	mydb.MockUpdateSuccess1 = true

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	var rt odb.RefreshToken
	rt.Token = "aaaaa"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAc
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445a"

	res, id := odbAc.AddAuthorizationCode(&ac, &at, &rt, nil)
	if res || id != 0 {
		t.Fail()
	}
}

// func TestMySQLOauthDBACi_DeleteAuthorizationCodeScope1(t *testing.T) {
// 	res := odbAci.DeleteAuthorizationCode(cidAci, "1234")
// 	if !res {
// 		t.Fail()
// 	}
// }

func TestMySQLOauthDBAC_AddAuthorizationCodeScope(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockInsertSuccess5 = true
	mydb.MockInsertID5 = 1

	var rows [][]string
	row1 := []string{"1", "test1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAc
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445b"
	var scope = []string{"test1", "test2"}

	res, id := odbAc.AddAuthorizationCode(&ac, &at, &rt, &scope)
	if !res || id < 1 {
		t.Fail()
	} else {
		spID2Ac = id
	}
}

func TestMySQLOauthDBAC_AddAuthorizationCodeScopeFailscope(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	// mydb.MockInsertSuccess5 = true
	// mydb.MockInsertID5 = 1

	var rows [][]string
	row1 := []string{"1", "test1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ac odb.AuthorizationCode
	ac.ClientID = cidAc
	ac.UserID = "1234"
	ac.Expires = time.Now()
	ac.RandonAuthCode = "13445b"
	var scope = []string{"test1", "test2"}

	res, _ := odbAc.AddAuthorizationCode(&ac, &at, &rt, &scope)
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_GetAuthCodeScopeList(t *testing.T) {
	res := odbAc.GetAuthorizationCodeScopeList(spID2Ac)
	fmt.Println("auth code scope in get: ", res)
	if res == nil || (*res)[0].Scope != "test1" {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_AddAuthCodeRevolk(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var rows [][]string
	row1 := []string{"1", "test1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	var rv odb.AuthCodeRevolk
	rv.AuthorizationCode = spID2Ac
	res, id := odbAc.AddAuthCodeRevolk(nil, &rv)
	fmt.Println("revolk id: ", id)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_UpdateAuthCode(t *testing.T) {
	var ac odb.AuthorizationCode
	ac.RandonAuthCode = "13445bb"
	ac.AlreadyUsed = true
	ac.AuthorizationCode = spID2Ac
	res := odbAc.UpdateAuthorizationCode(&ac)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_UpdateAuthCodeToken(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "someacctoken2", "false", ""}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "someacctoken2", tt.Format("2006-01-02 15:04:05"), "3", ""}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "someacctoken2", "false", ""}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{"1", "someacctokenupd", tt.Format("2006-01-02 15:04:05"), "3", ""}
	mydb.MockRow4 = &getRow4

	var rows [][]string
	row1 := []string{"1", "test1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	ac := odbAc.GetAuthorizationCodeByCode("13445bb")

	fmt.Println("auth code in update token: ", ac)
	var rt odb.RefreshToken
	rt.Token = "somereftoken2upd"
	rfs, rfid := odbAc.AddRefreshToken(nil, &rt)
	fmt.Println("new refresh token: ", rfs)
	if rfs {
		at := odbAc.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at in update token: ", at)
		at.Token = "someacctokenupd"
		at.Expires = time.Now()
		at.RefreshTokenID = rfid
		tt := time.Now()
		ac.Expires = tt
		res := odbAc.UpdateAuthorizationCodeAndToken(ac, at)
		fmt.Println("auth code update token suc: ", res)
		ac2 := odbAc.GetAuthorizationCodeByCode("13445bb")
		fmt.Println("auth2 code in update token: ", ac2)
		fmt.Println("tt in update token: ", tt.UTC())
		fmt.Println("expires in update token: ", ac2.Expires)
		at2 := odbAc.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at2 in update token: ", at2)
		if !res || at2.Token != "someacctokenupd" {
			t.Fail()
		}
	}

}

func TestMySQLOauthDBAC_UpdateAuthCodeTokenFail(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = false

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "someacctoken2", "false", ""}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "someacctoken2", tt.Format("2006-01-02 15:04:05"), "3", ""}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "someacctoken2", "false", ""}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{"1", "someacctokenupd", tt.Format("2006-01-02 15:04:05"), "3", ""}
	mydb.MockRow4 = &getRow4

	var rows [][]string
	row1 := []string{"1", "test1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	ac := odbAc.GetAuthorizationCodeByCode("13445bb")

	fmt.Println("auth code in update token: ", ac)
	var rt odb.RefreshToken
	rt.Token = "somereftoken2upd"
	rfs, rfid := odbAc.AddRefreshToken(nil, &rt)
	fmt.Println("new refresh token: ", rfs)
	if rfs {
		at := odbAc.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at in update token: ", at)
		at.Token = "someacctokenupd"
		at.Expires = time.Now()
		at.RefreshTokenID = rfid
		tt := time.Now()
		ac.Expires = tt
		res := odbAc.UpdateAuthorizationCodeAndToken(ac, at)
		fmt.Println("auth code update token suc: ", res)
		ac2 := odbAc.GetAuthorizationCodeByCode("13445bb")
		fmt.Println("auth2 code in update token: ", ac2)
		fmt.Println("tt in update token: ", tt.UTC())
		fmt.Println("expires in update token: ", ac2.Expires)
		at2 := odbAc.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at2 in update token: ", at2)
		if res {
			t.Fail()
		}
	}

}

func TestMySQLOauthDBAC_UpdateAuthCodeTokenFail2(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockUpdateSuccess1 = false

	mydb.MockUpdateSuccess2 = false

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "someacctoken2", "false", ""}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "someacctoken2", tt.Format("2006-01-02 15:04:05"), "3", ""}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "someacctoken2", "false", ""}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{"1", "someacctokenupd", tt.Format("2006-01-02 15:04:05"), "3", ""}
	mydb.MockRow4 = &getRow4

	var rows [][]string
	row1 := []string{"1", "test1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	ac := odbAc.GetAuthorizationCodeByCode("13445bb")

	fmt.Println("auth code in update token: ", ac)
	var rt odb.RefreshToken
	rt.Token = "somereftoken2upd"
	rfs, rfid := odbAc.AddRefreshToken(nil, &rt)
	fmt.Println("new refresh token: ", rfs)
	if rfs {
		at := odbAc.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at in update token: ", at)
		at.Token = "someacctokenupd"
		at.Expires = time.Now()
		at.RefreshTokenID = rfid
		tt := time.Now()
		ac.Expires = tt
		res := odbAc.UpdateAuthorizationCodeAndToken(ac, at)
		fmt.Println("auth code update token suc: ", res)
		ac2 := odbAc.GetAuthorizationCodeByCode("13445bb")
		fmt.Println("auth2 code in update token: ", ac2)
		fmt.Println("tt in update token: ", tt.UTC())
		fmt.Println("expires in update token: ", ac2.Expires)
		at2 := odbAc.GetAccessToken(ac.AccessTokenID)
		fmt.Println("at2 in update token: ", at2)
		if res {
			t.Fail()
		}
	}

}

func TestMySQLOauthDBAC_GetAuthCodeByCode(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true", ""}
	mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{"1", "test1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "2", "test1", "13445bb", "true"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	res := odbAc.GetAuthorizationCodeByCode("13445bb")
	fmt.Println("auth code in get: ", res)
	if res == nil || res.RandonAuthCode != "13445bb" || res.AlreadyUsed != true {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_GetAuthCodeByClient(t *testing.T) {
	res := odbAc.GetAuthorizationCode(cidAc, "1234")
	fmt.Println("auth code in get by client: ", res)
	if len(*res) < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_GetAuthCodeByScope(t *testing.T) {
	res := odbAc.GetAuthorizationCodeByScope(cidAc, "1234", "test1")
	fmt.Println("auth code in get by scope: ", res)
	if len(*res) < 1 || (*res)[0].Scope != "test1" {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_DeleteAuthorizationCode(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockDeleteSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockDeleteSuccess5 = true

	mydb.MockDeleteSuccess6 = true

	mydb.MockDeleteSuccess7 = true

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var rows [][]string
	row1 := []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "2", "test1", "13445bb", "true"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	res := odbAc.DeleteAuthorizationCode(cidAc, "1234")
	fmt.Println("auth code delete: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_DeleteAuthorizationCodeFail1(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockDeleteSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockDeleteSuccess5 = false

	mydb.MockDeleteSuccess6 = true

	mydb.MockDeleteSuccess7 = true

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var rows [][]string
	row1 := []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "2", "test1", "13445bb", "true"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	res := odbAc.DeleteAuthorizationCode(cidAc, "1234")
	fmt.Println("auth code delete: ", res)
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_DeleteAuthorizationCodeFail2(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockDeleteSuccess4 = false
	mydb.MockInsertID4 = 1

	mydb.MockDeleteSuccess5 = true

	mydb.MockDeleteSuccess6 = true

	mydb.MockDeleteSuccess7 = true

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var rows [][]string
	row1 := []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "2", "test1", "13445bb", "true"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	res := odbAc.DeleteAuthorizationCode(cidAc, "1234")
	fmt.Println("auth code delete: ", res)
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_DeleteAuthorizationCodeFail3(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = false
	mydb.MockInsertID3 = 1

	mydb.MockDeleteSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockDeleteSuccess5 = true

	mydb.MockDeleteSuccess6 = true

	mydb.MockDeleteSuccess7 = true

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var rows [][]string
	row1 := []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "2", "test1", "13445bb", "true"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	res := odbAc.DeleteAuthorizationCode(cidAc, "1234")
	fmt.Println("auth code delete: ", res)
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_DeleteAuthorizationCodeFail4(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = false

	mydb.MockDeleteSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockDeleteSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockDeleteSuccess5 = true

	mydb.MockDeleteSuccess6 = true

	mydb.MockDeleteSuccess7 = true

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var rows [][]string
	row1 := []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "2", "test1", "13445bb", "true"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	res := odbAc.DeleteAuthorizationCode(cidAc, "1234")
	fmt.Println("auth code delete: ", res)
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBAC_DeleteAuthorizationCodeFail5(t *testing.T) {
	var mydb mdb.MyDBMock
	dbAc = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = false

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = true
	mydb.MockInsertID3 = 1

	mydb.MockDeleteSuccess4 = true
	mydb.MockInsertID4 = 1

	mydb.MockDeleteSuccess5 = true

	mydb.MockDeleteSuccess6 = true

	mydb.MockDeleteSuccess7 = true

	mydb.MockUpdateSuccess1 = true

	mydb.MockUpdateSuccess2 = true

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var rows [][]string
	row1 := []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "2", "test1", "13445bb", "true"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var moadb MySQLOauthDB
	moadb.DB = dbAc

	odbAc = &moadb

	res := odbAc.DeleteAuthorizationCode(cidAc, "1234")
	fmt.Println("auth code delete fail5: ", res)
	if res {
		t.Fail()
	}
}
func TestMySQLOauthDBAC_DeleteClient(t *testing.T) {
	suc := odbAc.DeleteClient(cidAc)
	if !suc {
		t.Fail()
	}
}
