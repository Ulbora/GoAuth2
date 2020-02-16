package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

func TestOauthManagerAllowedURI_AddClientAllowedURI(t *testing.T) {

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

	mydb.MockInsertID1 = 2
	mydb.MockInsertSuccess1 = true

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
	var cu ClientAllowedURI
	cu.URI = "testuri"
	cu.ClientID = 2
	suc, id := m.AddClientAllowedURI(&cu)
	if !suc || id != 2 {
		t.Fail()
	}
}

func TestOauthManagerAllowedURI_UpdateClientAllowedURI(t *testing.T) {

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

	//mydb.MockInsertID1 = 2
	mydb.MockUpdateSuccess1 = true

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
	var cu ClientAllowedURI
	cu.ID = 1
	cu.URI = "testuri"
	cu.ClientID = 2
	suc := m.UpdateClientAllowedURI(&cu)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerAllowedURI_GetClientAllowedURI(t *testing.T) {

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

	//mydb.MockInsertID1 = 2
	//mydb.MockUpdateSuccess1 = true

	// var rows [][]string
	// row1 := []string{"1", "code", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

	var mGetRow db.DbRow
	mGetRow.Row = []string{"2", "testUri", "3"}
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

	cu := m.GetClientAllowedURI(2)
	if cu.ID != 2 {
		t.Fail()
	}
}

func TestOauthManagerAllowedURI_GetClientAllowedURIList(t *testing.T) {

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

	//mydb.MockInsertID1 = 2
	//mydb.MockUpdateSuccess1 = true

	var rows [][]string
	row1 := []string{"1", "testuri1", "2"}
	rows = append(rows, row1)
	row2 := []string{"2", "testuri2", "2"}
	rows = append(rows, row2)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	// var mGetRow db.DbRow
	// mGetRow.Row = []string{"2", "testUri", "3"}
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

	cul := m.GetClientAllowedURIList(2)
	if len(*cul) != 2 || (*cul)[1].ID != 2 {
		t.Fail()
	}
}

func TestOauthManagerAllowedURI_DeleteClientAllowedURI(t *testing.T) {

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

	//mydb.MockInsertID1 = 2
	mydb.MockDeleteSuccess1 = true

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

	suc := m.DeleteClientAllowedURI(2)
	if !suc {
		t.Fail()
	}
}
