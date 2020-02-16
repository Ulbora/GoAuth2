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

func TestOauthManagerGrantType_grantTypeTurnedOn(t *testing.T) {

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
	row1 := []string{"1", "code", "2"}
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
	moadb.Log = &l
	man.Db = odbAu

	res := man.grantTypeTurnedOn(1, "code")
	fmt.Println("grant turned on: ", res)
	if !res {
		t.Fail()
	}

}

func TestOauthManagerGrantType_AddGrantType(t *testing.T) {

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
	var gt ClientGrantType
	gt.GrantType = "code"
	gt.ClientID = 2
	suc, id := m.AddClientGrantType(&gt)
	if !suc || id != 2 {
		t.Fail()
	}
}

func TestOauthManagerGrantType_GetGrantTypeList(t *testing.T) {

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

	// mydb.MockInsertID1 = 2
	// mydb.MockInsertSuccess1 = true

	var rows [][]string
	row1 := []string{"1", "code", "2"}
	rows = append(rows, row1)
	row2 := []string{"2", "client", "2"}
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

	gtl := m.GetClientGrantTypeList(2)
	if len(*gtl) != 2 {
		t.Fail()
	}
}

func TestOauthManagerGrantType_DeleteGrantType(t *testing.T) {

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

	suc := m.DeleteClientGrantType(2)
	if !suc {
		t.Fail()
	}
}
