package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

func TestOauthManagerRoleURI_AddClientRoleURI(t *testing.T) {

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
	var cr ClientRoleURI
	cr.ClientRoleID = 1
	cr.ClientAllowedURIID = 2
	suc := m.AddClientRoleURI(&cr)
	if !suc {
		t.Fail()
	}

}

func TestOauthManagerRoleURI_GetClientRoleURIList(t *testing.T) {

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
	//mydb.MockInsertSuccess1 = true

	var rows [][]string
	row1 := []string{"2", "2"}
	rows = append(rows, row1)
	row2 := []string{"2", "3"}
	rows = append(rows, row2)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

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

	rul := m.GetClientRoleAllowedURIList(2)
	if len(*rul) != 2 || (*rul)[1].ClientAllowedURIID != 3 {
		t.Fail()
	}
}

func TestOauthManagerRoleURI_DeleteClientRoleURI(t *testing.T) {

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
	var cr ClientRoleURI
	cr.ClientRoleID = 1
	cr.ClientAllowedURIID = 2
	suc := m.DeleteClientRoleURI(&cr)
	if !suc {
		t.Fail()
	}

}
