package managers

import (
	"fmt"
	"testing"
	"time"

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
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn := m.GetAuthCodeToken(&tkr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
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
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn := m.GetAuthCodeToken(&tkr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
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
	man.Db = odbAu
	var m Manager
	m = &man

	var tkr AuthCodeTokenReq
	tkr.ClientID = 2
	tkr.Secret = "12345"
	tkr.Code = "5555"
	tkr.RedirectURI = "google.com"
	suc, tkn := m.GetAuthCodeToken(&tkr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
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
	man.Db = odbAu
	var m Manager
	m = &man
	var ctr CredentialsTokenReq
	ctr.ClientID = 2
	ctr.Secret = "12345"
	suc, tkn := m.GetCredentialsToken(&ctr)
	fmt.Println("suc: ", suc)
	fmt.Println("tkn: ", tkn)
	if !suc || tkn.TokenType != "bearer" {
		t.Fail()
	}
}
