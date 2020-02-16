package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

func TestOauthManagerImplicit_AuthorizeImplicit(t *testing.T) {

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

	var rows [][]string
	row1 := []string{"1", "implicit", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "3", "user", "2", "test"}
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
	row4 := []string{"1", "3", "user", "2", "test"}
	rows4 = append(rows4, row4)
	var dbrows4 db.DbRows
	dbrows4.Rows = rows4
	mydb.MockRows4 = &dbrows4

	var rows5 [][]string
	row5 := []string{"4", "somerole", "1", "someurl", "2"}
	rows5 = append(rows5, row5)
	var dbrows5 db.DbRows
	dbrows5.Rows = rows5
	mydb.MockRows5 = &dbrows5

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow3 = &mGetRow3

	// var tt = time.Now()
	// var mGetRow3 db.DbRow
	// mGetRow3.Row = []string{"2", "testUri", tt.Format("2006-01-02 15:04:05"), "2"}
	// mydb.MockRow3 = &mGetRow3

	// var mGetRow4 db.DbRow
	// mGetRow4.Row = []string{"2", "testUri"}
	// mydb.MockRow4 = &mGetRow4

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 2

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 20

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 5

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	moadb.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var imp Implicit
	imp.ClientID = 2
	imp.UserID = "123"
	imp.Scope = "web"
	imp.RedirectURI = "google.com"
	imp.CallbackURI = "google.com"

	suc, rtn := m.AuthorizeImplicit(&imp)
	fmt.Println("rtn: ", rtn)
	if !suc || rtn.ID != 20 || rtn.Token == "" {
		t.Fail()
	}

}

func TestOauthManagerImplicit_AuthorizeImplicitNewScope(t *testing.T) {

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

	var rows [][]string
	row1 := []string{"1", "implicit", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var rows2 [][]string
	row2 := []string{"1", "3", "user", "2", "test"}
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
	row4 := []string{"1", "3", "user", "2", "test"}
	rows4 = append(rows4, row4)
	var dbrows4 db.DbRows
	dbrows4.Rows = rows4
	mydb.MockRows4 = &dbrows4

	var rows5 [][]string
	row5 := []string{"4", "somerole", "1", "someurl", "2"}
	rows5 = append(rows5, row5)
	var dbrows5 db.DbRows
	dbrows5.Rows = rows5
	mydb.MockRows5 = &dbrows5

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow3 = &mGetRow3

	// var tt = time.Now()
	// var mGetRow3 db.DbRow
	// mGetRow3.Row = []string{"2", "testUri", tt.Format("2006-01-02 15:04:05"), "2"}
	// mydb.MockRow3 = &mGetRow3

	// var mGetRow4 db.DbRow
	// mGetRow4.Row = []string{"2", "testUri"}
	// mydb.MockRow4 = &mGetRow4

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 2

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 20

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 5

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	moadb.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var imp Implicit
	imp.ClientID = 2
	imp.UserID = "123"
	imp.Scope = "cloud"
	imp.RedirectURI = "google.com"
	imp.CallbackURI = "google.com"

	suc, rtn := m.AuthorizeImplicit(&imp)
	fmt.Println("rtn: ", rtn)
	if !suc || rtn.ID != 20 || rtn.Token == "" {
		t.Fail()
	}

}

func TestOauthManagerImplicit_AuthorizeImplicitNew(t *testing.T) {

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

	var rows [][]string
	row1 := []string{"1", "implicit", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	// var rows2 [][]string
	// row2 := []string{"1", "3", "user", "2", "test"}
	// rows2 = append(rows2, row2)
	// var dbrows2 db.DbRows
	// dbrows2.Rows = rows2
	// mydb.MockRows2 = &dbrows2

	var rows2 [][]string
	//row2 := []string{"1", "3", "user", "2", "test"}
	//rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	// var rows3 [][]string
	// row3 := []string{"1", "web", "2"}
	// rows3 = append(rows3, row3)
	// var dbrows3 db.DbRows
	// dbrows3.Rows = rows3
	// mydb.MockRows3 = &dbrows3

	// var rows4 [][]string
	// row4 := []string{"1", "3", "user", "2", "test"}
	// rows4 = append(rows4, row4)
	// var dbrows4 db.DbRows
	// dbrows4.Rows = rows4
	// mydb.MockRows4 = &dbrows4

	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow3 = &mGetRow3

	// var tt = time.Now()
	// var mGetRow3 db.DbRow
	// mGetRow3.Row = []string{"2", "testUri", tt.Format("2006-01-02 15:04:05"), "2"}
	// mydb.MockRow3 = &mGetRow3

	// var mGetRow4 db.DbRow
	// mGetRow4.Row = []string{"2", "testUri"}
	// mydb.MockRow4 = &mGetRow4

	// mydb.MockDeleteSuccess1 = true
	// mydb.MockDeleteSuccess2 = true
	// mydb.MockDeleteSuccess3 = true

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 2

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 20

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 5

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	moadb.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var imp Implicit
	imp.ClientID = 2
	imp.UserID = "123"
	imp.Scope = "cloud"
	imp.RedirectURI = "google.com"
	imp.CallbackURI = "google.com"

	suc, rtn := m.AuthorizeImplicit(&imp)
	fmt.Println("rtn: ", rtn)
	if !suc || rtn.ID != 20 || rtn.Token == "" {
		t.Fail()
	}

}

func TestOauthManagerImplicit_AuthorizeImplicitClientDisabled(t *testing.T) {

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
	mGetRow.Row = []string{"2", "testUri", "test", "test", "test", "false", "false"}
	mydb.MockRow1 = &mGetRow

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "testUri", "2"}
	mydb.MockRow2 = &mGetRow2

	var rows [][]string
	row1 := []string{"1", "implicit", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	// var rows2 [][]string
	// row2 := []string{"1", "3", "user", "2", "test"}
	// rows2 = append(rows2, row2)
	// var dbrows2 db.DbRows
	// dbrows2.Rows = rows2
	// mydb.MockRows2 = &dbrows2

	var rows2 [][]string
	//row2 := []string{"1", "3", "user", "2", "test"}
	//rows2 = append(rows2, row2)
	var dbrows2 db.DbRows
	dbrows2.Rows = rows2
	mydb.MockRows2 = &dbrows2

	// var rows3 [][]string
	// row3 := []string{"1", "web", "2"}
	// rows3 = append(rows3, row3)
	// var dbrows3 db.DbRows
	// dbrows3.Rows = rows3
	// mydb.MockRows3 = &dbrows3

	// var rows4 [][]string
	// row4 := []string{"1", "3", "user", "2", "test"}
	// rows4 = append(rows4, row4)
	// var dbrows4 db.DbRows
	// dbrows4.Rows = rows4
	// mydb.MockRows4 = &dbrows4

	var rows3 [][]string
	row3 := []string{"4", "somerole", "1", "someurl", "2"}
	rows3 = append(rows3, row3)
	var dbrows3 db.DbRows
	dbrows3.Rows = rows3
	mydb.MockRows3 = &dbrows3

	var mGetRow3 db.DbRow
	mGetRow3.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow3 = &mGetRow3

	// var tt = time.Now()
	// var mGetRow3 db.DbRow
	// mGetRow3.Row = []string{"2", "testUri", tt.Format("2006-01-02 15:04:05"), "2"}
	// mydb.MockRow3 = &mGetRow3

	// var mGetRow4 db.DbRow
	// mGetRow4.Row = []string{"2", "testUri"}
	// mydb.MockRow4 = &mGetRow4

	// mydb.MockDeleteSuccess1 = true
	// mydb.MockDeleteSuccess2 = true
	// mydb.MockDeleteSuccess3 = true

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 2

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 20

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 5

	mydb.MockInsertSuccess4 = true
	mydb.MockInsertID4 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	moadb.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var imp Implicit
	imp.ClientID = 2
	imp.UserID = "123"
	imp.Scope = "cloud"
	imp.RedirectURI = "google.com"
	imp.CallbackURI = "google.com"

	suc, rtn := m.AuthorizeImplicit(&imp)
	fmt.Println("rtn: ", rtn)
	if suc {
		t.Fail()
	}

}

func TestOauthManagerImplicit_CheckImplicitApplicationAuthorization(t *testing.T) {

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

	// var mGetRow db.DbRow
	// mGetRow.Row = []string{"2", "testUri", "test", "test", "test", "false", "false"}
	// mydb.MockRow1 = &mGetRow

	// var mGetRow2 db.DbRow
	// mGetRow2.Row = []string{"2", "testUri", "2"}
	// mydb.MockRow2 = &mGetRow2

	// var rows [][]string
	// row1 := []string{"1", "implicit", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

	var rows1 [][]string
	row1 := []string{"1", "3", "user", "2", "test"}
	rows1 = append(rows1, row1)
	var dbrows1 db.DbRows
	dbrows1.Rows = rows1
	mydb.MockRows1 = &dbrows1

	// var rows2 [][]string
	// //row2 := []string{"1", "3", "user", "2", "test"}
	// //rows2 = append(rows2, row2)
	// var dbrows2 db.DbRows
	// dbrows2.Rows = rows2
	// mydb.MockRows2 = &dbrows2

	// var rows3 [][]string
	// row3 := []string{"1", "web", "2"}
	// rows3 = append(rows3, row3)
	// var dbrows3 db.DbRows
	// dbrows3.Rows = rows3
	// mydb.MockRows3 = &dbrows3

	// var rows4 [][]string
	// row4 := []string{"1", "3", "user", "2", "test"}
	// rows4 = append(rows4, row4)
	// var dbrows4 db.DbRows
	// dbrows4.Rows = rows4
	// mydb.MockRows4 = &dbrows4

	// var rows3 [][]string
	// row3 := []string{"4", "somerole", "1", "someurl", "2"}
	// rows3 = append(rows3, row3)
	// var dbrows3 db.DbRows
	// dbrows3.Rows = rows3
	// mydb.MockRows3 = &dbrows3

	// var mGetRow3 db.DbRow
	// mGetRow3.Row = []string{"2", "6g651dfg6gf6"}
	// mydb.MockRow3 = &mGetRow3

	// var tt = time.Now()
	// var mGetRow3 db.DbRow
	// mGetRow3.Row = []string{"2", "testUri", tt.Format("2006-01-02 15:04:05"), "2"}
	// mydb.MockRow3 = &mGetRow3

	// var mGetRow4 db.DbRow
	// mGetRow4.Row = []string{"2", "testUri"}
	// mydb.MockRow4 = &mGetRow4

	// mydb.MockDeleteSuccess1 = true
	// mydb.MockDeleteSuccess2 = true
	// mydb.MockDeleteSuccess3 = true

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 2

	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 20

	// mydb.MockInsertSuccess3 = true
	// mydb.MockInsertID3 = 5

	// mydb.MockInsertSuccess4 = true
	// mydb.MockInsertID4 = 6

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	moadb.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man
	var imp Implicit
	imp.ClientID = 2
	imp.UserID = "123"
	imp.Scope = "cloud"
	imp.RedirectURI = "google.com"
	imp.CallbackURI = "google.com"

	suc := m.CheckImplicitApplicationAuthorization(&imp)
	fmt.Println("authorized: ", suc)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerImplicit_ValidateAuthCodeClientAndCallback(t *testing.T) {

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
	mGetRow1.Row = []string{"5", "someurl.com", "25"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"25", "secret", "testname", "testWebSite", "test", "true", "false"}
	mydb.MockRow2 = &mGetRow2

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	var l lg.Logger
	man.Log = &l
	moadb.Log = &l
	man.Db = odbAu
	var m Manager
	m = &man

	var imp Implicit
	imp.ClientID = 25
	imp.UserID = "123"
	imp.Scope = "cloud"
	imp.RedirectURI = "someurl.com"
	imp.CallbackURI = "google.com"
	ip := m.ValidateImplicitClientAndCallback(&imp)
	fmt.Println("auth: ", ip)
	if !ip.Valid || ip.ClientName != "testname" || ip.WebSite != "testWebSite" {
		t.Fail()
	}
}
