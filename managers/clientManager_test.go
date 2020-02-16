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

func TestOauthManagerClient_AddClient(t *testing.T) {

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
	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 2
	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 2

	// var rows [][]string
	// row1 := []string{"1", "code", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

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

	var c Client
	c.Name = "tester"
	c.WebSite = "www"
	c.Email = "test@test.com"
	c.Enabled = true
	c.Paid = false
	var uris []ClientRedirectURI
	var uri ClientRedirectURI
	uri.URI = "test"
	uris = append(uris, uri)
	c.RedirectURIs = &uris

	suc, id := m.AddClient(&c)
	if !suc || id == 0 {
		t.Fail()
	}
}

func TestOauthManagerClient_UpdateClient(t *testing.T) {

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
	mydb.MockUpdateSuccess1 = true
	// mydb.MockInsertID1 = 2
	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 2

	// var rows [][]string
	// row1 := []string{"1", "code", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

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

	var c Client
	c.ClientID = 22
	c.Name = "tester"
	c.WebSite = "www"
	c.Email = "test@test.com"
	c.Secret = "12345"
	c.Enabled = true
	c.Paid = false
	var uris []ClientRedirectURI
	var uri ClientRedirectURI
	uri.URI = "test"
	uris = append(uris, uri)
	c.RedirectURIs = &uris

	suc := m.UpdateClient(&c)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerClient_UpdateClient2(t *testing.T) {

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
	mydb.MockUpdateSuccess1 = true
	// mydb.MockInsertID1 = 2
	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 2

	// var rows [][]string
	// row1 := []string{"1", "code", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

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

	var c Client
	c.ClientID = 22
	c.Name = "tester"
	c.WebSite = "www"
	c.Email = "test@test.com"
	//c.Secret = "12345"
	c.Enabled = true
	c.Paid = false
	// uri := []string{"testuri"}
	// c.RedirectURIs = &uri

	suc := m.UpdateClient(&c)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerClient_GetClient(t *testing.T) {

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

	var rows [][]string
	row1 := []string{"3", "testurl", "2"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "secret", "tester", "tester.com", "t@t.com", "true", "false"}
	mydb.MockRow1 = &mGetRow

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

	c := m.GetClient(1)
	fmt.Println("client in get test: ", c)
	fmt.Println("client uris in get test: ", c.RedirectURIs)
	if c.ClientID != 2 || (*c.RedirectURIs)[0].ID != 3 {
		t.Fail()
	}
}

func TestOauthManagerClient_GetClients(t *testing.T) {

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

	var rows [][]string
	row1 := []string{"3", "secret", "tester", "tester.com", "t@t.com", "true", "false"}
	rows = append(rows, row1)
	row2 := []string{"4", "secret", "tester", "tester.com", "t@t.com", "true", "false"}
	rows = append(rows, row2)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	// var mGetRow db.DbRow
	// mGetRow.Row = []string{"2", "secret", "tester", "tester.com", "t@t.com", "true", "false"}
	// mydb.MockRow1 = &mGetRow

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

	cs := m.GetClientList()
	fmt.Println("clients in get test: ", cs)

	if len(*cs) != 2 || (*cs)[0].ClientID != 3 {
		t.Fail()
	}
}

func TestOauthManagerClient_GetClientSearch(t *testing.T) {

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

	var rows [][]string
	row1 := []string{"3", "secret", "tester", "tester.com", "t@t.com", "true", "false"}
	rows = append(rows, row1)
	row2 := []string{"4", "secret", "tester", "tester.com", "t@t.com", "true", "false"}
	rows = append(rows, row2)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	// var mGetRow db.DbRow
	// mGetRow.Row = []string{"2", "secret", "tester", "tester.com", "t@t.com", "true", "false"}
	// mydb.MockRow1 = &mGetRow

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

	cs := m.GetClientSearchList("tester")
	fmt.Println("clients in get test: ", cs)

	if len(*cs) != 2 || (*cs)[0].ClientID != 3 {
		t.Fail()
	}
}

func TestOauthManagerClient_DeleteClient(t *testing.T) {

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
	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	// mydb.MockInsertID1 = 2
	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 2

	// var rows [][]string
	// row1 := []string{"1", "code", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

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

	suc := m.DeleteClient(2)
	if !suc {
		t.Fail()
	}
}
