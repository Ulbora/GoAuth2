package managers

import (
	"fmt"
	"testing"
	"time"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

func TestOauthManagerToken_GetAuthCodeToken(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, _ := m.GetAuthCodeToken(&tkr)
	fmt.Println("auth code tkn suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenNoAccessToken(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"0", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, err := m.GetAuthCodeToken(&tkr)
	fmt.Println("auth code tkn suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidGrantError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenUpdateFailed(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	//mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, err := m.GetAuthCodeToken(&tkr)
	fmt.Println("auth code tkn suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidGrantError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenRevolked(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{"2", "3"}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, err := m.GetAuthCodeToken(&tkr)
	fmt.Println("auth code tkn suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenBadClient(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 1
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, err := m.GetAuthCodeToken(&tkr)
	fmt.Println("auth code tkn suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenBadRedirect(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"0", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, err := m.GetAuthCodeToken(&tkr)
	fmt.Println("auth code tkn suc: ", suc)
	fmt.Println("tkn: ", tkn)
	fmt.Println("err: ", err)
	if suc || err != "invalid_grant" {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenBadSecret(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = ""
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, err := m.GetAuthCodeToken(&tkr)
	fmt.Println("auth code tkn suc: ", suc)
	fmt.Println("tkn: ", tkn)
	fmt.Println("err: ", err)
	if suc || err != "invalid_client" {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenAlreadUsed(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "true"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "5"}
	mydb.MockRow5 = &getRow5

	var getRow6 db.DbRow
	getRow6.Row = []string{"1", "somereftoken2"}
	mydb.MockRow6 = &getRow6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, err := m.GetAuthCodeToken(&tkr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeTokenNoRefresh(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var tt = time.Now()
	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	mydb.MockRow3 = &getRow3

	var getRow4 db.DbRow
	getRow4.Row = []string{}
	mydb.MockRow4 = &getRow4

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 4

	mydb.MockUpdateSuccess1 = true

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", "someacctoken2", nowTime, "0"}
	mydb.MockRow5 = &getRow5

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn, _ := m.GetAuthCodeToken(&tkr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerToken_GetCredToken(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "client_credentials", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ctr CredentialsTokenReq
	ctr.ClientID = 2
	ctr.Secret = "12345"
	suc, tkn, _ := m.GetCredentialsToken(&ctr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if !suc || tkn.TokenType != "bearer" {
		t.Fail()
	}
}

func TestOauthManagerToken_GetCredTokenDeleteFailed(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "client_credentials", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	//mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ctr CredentialsTokenReq
	ctr.ClientID = 2
	ctr.Secret = "12345"
	suc, tkn, err := m.GetCredentialsToken(&ctr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != accessDeniedError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetCredTokenGrantOff(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "client_credentials111", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ctr CredentialsTokenReq
	ctr.ClientID = 2
	ctr.Secret = "12345"
	suc, tkn, err := m.GetCredentialsToken(&ctr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != accessDeniedError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetCredTokenBadSecret(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "client_credentials", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ctr CredentialsTokenReq
	ctr.ClientID = 2
	ctr.Secret = ""
	suc, tkn, err := m.GetCredentialsToken(&ctr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeWithRefToken(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "12345", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("12345")
	pl.ClientID = 2
	pl.Subject = codeGrantType
	pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = codeGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("12345"), "code")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 2
	rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, _ := m.GetAuthCodeAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if !suc || tkn.TokenType != "bearer" {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeWithRefTokenTokenNotValid(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "12345", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("12345")
	pl.ClientID = 2
	pl.Subject = codeGrantType
	pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = codeGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("12345"), "code2")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 2
	rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, err := m.GetAuthCodeAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeWithRefTokenNoTokenKey(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", ""}
	mydb.MockRow4 = &mGetRow4

	var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "12345", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("12345")
	pl.ClientID = 2
	pl.Subject = codeGrantType
	pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = codeGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("12345"), "code")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 2
	rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, err := m.GetAuthCodeAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeWithRefTokenBadSecret(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "false", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "12345", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("12345")
	pl.ClientID = 2
	pl.Subject = codeGrantType
	pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = codeGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("12345"), "code")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 2
	rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, err := m.GetAuthCodeAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetAuthCodeWithRefTokenBadClient(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "12345", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("12345")
	pl.ClientID = 2
	pl.Subject = codeGrantType
	pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = codeGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("12345"), "code")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 0
	rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, err := m.GetAuthCodeAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidRequestError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetPasseordToken(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "password", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "tester1", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow2 db.DbRow
	getRow2.Row = []string{"2", "someacctoken2", nowTime, "1"}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "somereftoken2"}
	mydb.MockRow3 = &getRow3

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var mGetRow5 db.DbRow
	mGetRow5.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow5 = &mGetRow5

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ptr PasswordTokenReq
	ptr.ClientID = 2
	ptr.Username = "tester1"
	suc, tkn, _ := m.GetPasswordToken(&ptr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if !suc || tkn.TokenType != "bearer" {
		t.Fail()
	}
}

func TestOauthManagerToken_GetPasseordTokenAddFailed(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "password", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "tester1", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow2 db.DbRow
	getRow2.Row = []string{"2", "someacctoken2", nowTime, "1"}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "somereftoken2"}
	mydb.MockRow3 = &getRow3

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var mGetRow5 db.DbRow
	mGetRow5.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow5 = &mGetRow5

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ptr PasswordTokenReq
	ptr.ClientID = 2
	ptr.Username = "tester1"
	suc, tkn, err := m.GetPasswordToken(&ptr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != accessDeniedError {
		t.Fail()
	}
}
func TestOauthManagerToken_GetPasseordTokenDeleteFailed(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "password", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "tester1", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow2 db.DbRow
	getRow2.Row = []string{"2", "someacctoken2", nowTime, "1"}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "somereftoken2"}
	mydb.MockRow3 = &getRow3

	//mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var mGetRow5 db.DbRow
	mGetRow5.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow5 = &mGetRow5

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ptr PasswordTokenReq
	ptr.ClientID = 2
	ptr.Username = "tester1"
	suc, tkn, err := m.GetPasswordToken(&ptr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != accessDeniedError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetPasseordTokenGrantOff(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "password1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "tester1", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow2 db.DbRow
	getRow2.Row = []string{"2", "someacctoken2", nowTime, "1"}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "somereftoken2"}
	mydb.MockRow3 = &getRow3

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var mGetRow5 db.DbRow
	mGetRow5.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow5 = &mGetRow5

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ptr PasswordTokenReq
	ptr.ClientID = 2
	ptr.Username = "tester1"
	suc, tkn, err := m.GetPasswordToken(&ptr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != accessDeniedError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetPasseordTokenDisabledClient(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "12345", "test", "test", "test", "false", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "password", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	//get cred grant for del
	var rows2 [][]string
	row2 := []string{"1", "2", "tester1", "2"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow2 db.DbRow
	getRow2.Row = []string{"2", "someacctoken2", nowTime, "1"}
	mydb.MockRow2 = &getRow2

	var getRow3 db.DbRow
	getRow3.Row = []string{"1", "somereftoken2"}
	mydb.MockRow3 = &getRow3

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true

	//ClientRoleAllowedURIList
	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	//access token key
	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	var mGetRow5 db.DbRow
	mGetRow5.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow5 = &mGetRow5

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var ptr PasswordTokenReq
	ptr.ClientID = 2
	ptr.Username = "tester1"
	suc, tkn, err := m.GetPasswordToken(&ptr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetPasswordWithRefToken(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	//var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "tester1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = passwordGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = passwordGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("tester1"), "password")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	// mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 2
	//rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, _ := m.GetPasswordAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if !suc || tkn.TokenType != "bearer" {
		t.Fail()
	}
}

func TestOauthManagerToken_GetPasswordWithRefTokenBadKey(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", ""}
	mydb.MockRow4 = &mGetRow4

	//var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "tester1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = passwordGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = passwordGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("tester1"), "password")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	// mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 2
	//rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, err := m.GetPasswordAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidClientError {
		t.Fail()
	}
}

func TestOauthManagerToken_GetPasswordWithRefTokenBadClient(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "12345", "test", "test", "test", "true", "false"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow4 = &mGetRow4

	//var tt = time.Now()
	// var getRow4 db.DbRow
	// getRow4.Row = []string{"1", "2", "3", tt.Format("2006-01-02 15:04:05"), "3", "13445bb", "false"}
	// mydb.MockRow4 = &getRow4
	var rows [][]string
	row1 := []string{"1", "2", "tester1", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = passwordGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = passwordGrantType
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	var nowTime = time.Now().Format(odb.TimeFormat)
	var getRow5 db.DbRow
	getRow5.Row = []string{"2", accessToken, nowTime, "5"}
	mydb.MockRow5 = &getRow5

	rtoken := man.GenerateRefreshToken(2, hashUser("tester1"), "password")
	fmt.Println("rtoken: ", rtoken)

	var mGetRow6 db.DbRow
	mGetRow6.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow6 = &mGetRow6

	var mGetRow7 db.DbRow
	mGetRow7.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow7 = &mGetRow7

	mydb.MockUpdateSuccess1 = true
	// mydb.MockUpdateSuccess2 = true

	var rtr RefreshTokenReq
	rtr.ClientID = 0
	//rtr.Secret = "12345"
	rtr.RefreshToken = rtoken
	suc, tkn, err := m.GetPasswordAccesssTokenWithRefreshToken(&rtr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if suc || err != invalidRequestError {
		t.Fail()
	}
}
