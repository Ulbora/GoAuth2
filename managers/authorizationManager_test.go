package managers

import (
	"testing"
	"time"

	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

func TestOauthManagerAuthCode_AuthorizeAuthCode(t *testing.T) {

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
	mGetRow.Row = []string{"2", "testUri", "test", "test", "test", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	//mydb.MockInsertID1 = 2
	//mydb.MockDeleteSuccess1 = true

	var rows [][]string
	row1 := []string{"1", "code", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var tt = time.Now()
	var rows2 [][]string
	row2 := []string{"1", "3", "code", tt.Format("2006-01-02 15:04:05"), "2", "test", "false", "test"}
	rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	var rows3 [][]string
	row3 := []string{"1", "web", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	var rows4 [][]string
	row4 := []string{"1", "3", "code", tt.Format("2006-01-02 15:04:05"), "2", "test", "false", "test"}
	rows4 = append(rows4, row4)
	var dbrows4 db.DbRows
	dbrows4.Rows = rows4
	mydb.MockRows4 = &dbrows4

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "testUri", tt.Format("2006-01-02 15:04:05"), "2"}
	mydb.MockRow3 = &mGetRow3

	var mGetRow4 db.DbRow
	mGetRow4.Row = []string{"2", "testUri"}
	mydb.MockRow4 = &mGetRow4

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true
	mydb.MockDeleteSuccess4 = true
	mydb.MockDeleteSuccess5 = true

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	man.Db = odbAu
	var m Manager
	m = &man
	var ac AuthCode
	ac.ClientID = 2
	ac.UserID = "tester"
	ac.Scope = "web"
	ac.RedirectURI = "someurl.com"
	ac.CallbackURI = "someotherurl.com"
	suc, acode, cstring := m.AuthorizeAuthCode(&ac)
	if suc || acode != 0 || cstring != "" {
		t.Fail()
	}
}
