package mysqldb

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbIg db.Database
var odbIg odb.Oauth2DB
var cidIg int64
var spIDIg int64
var spID2Ig int64

func TestMySQLOauthDBIg_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

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
	mydb.MockInsertSuccess6 = true
	mydb.MockInsertID6 = 1
	mydb.MockInsertSuccess7 = true
	mydb.MockInsertID7 = 1
	mydb.MockInsertSuccess8 = true
	mydb.MockInsertID8 = 1

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true
	mydb.MockDeleteSuccess4 = true

	// var getRow db.DbRow
	// getRow.Row = []string{"1", "2", "user", "4", ""}
	// mydb.MockRow1 = &getRow

	var rows1 [][]string
	row1 := []string{"1", "scope1", "4"}
	rows1 = append(rows1, row1)
	var dbrows1 db.DbRows
	dbrows1.Rows = rows1
	mydb.MockRows1 = &dbrows1

	// var rows [][]string
	// row1 := []string{"1", "2", "user", "4", ""}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows
	// mydb.MockRows2 = &dbrows
	// mydb.MockRows3 = &dbrows

	var rows [][]string
	row := []string{"1", "2", "user", "4", ""}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	//mydb.MockRows1 = &dbrows
	mydb.MockRows2 = &dbrows
	mydb.MockRows3 = &dbrows
	mydb.MockRows4 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()
}

func TestMySQLOauthDBIg_AddClientNullUri(t *testing.T) {
	var c odb.Client
	c.Secret = "12345"
	c.Name = "tester"
	c.Email = "bob@bob.com"
	c.WebSite = "www.bob.com"
	c.Enabled = true
	c.Paid = false

	fmt.Println("before db add")
	res, id := odbIg.AddClient(&c, nil)
	fmt.Println("client add res: ", res)
	fmt.Println("client id: ", id)
	if !res || id == 0 {
		t.Fail()
	} else {
		cidIg = id
	}
}

func TestMySQLOauthDBIg_AddImplicitGrantNoScope(t *testing.T) {

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ig odb.ImplicitGrant
	ig.ClientID = cidIg
	ig.UserID = "1234"
	res, igid := odbIg.AddImplicitGrant(&ig, &at, nil)
	if !res || igid <= 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_AddImplicitGrant(t *testing.T) {

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ig odb.ImplicitGrant
	ig.ClientID = cidIg
	ig.UserID = "1234"
	var scope = []string{"test1", "test2"}
	res, igid := odbIg.AddImplicitGrant(&ig, &at, &scope)
	if !res || igid <= 0 {
		t.Fail()
	} else {
		spIDIg = igid
	}

}

func TestMySQLOauthDBIg_GetImplicitGrantScopeList(t *testing.T) {
	res := odbIg.GetImplicitGrantScopeList(spIDIg)
	fmt.Println("implicit grant scope in get: ", res)
	if res == nil || (*res)[0].Scope != "scope1" {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_GetImplicitGrant(t *testing.T) {
	res := odbIg.GetImplicitGrant(cidIg, "1234")
	if len(*res) < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_GetImplicitGrantByScope(t *testing.T) {
	res := odbIg.GetImplicitGrantByScope(cidIg, "1234", "test1")
	if len(*res) != 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_AddImplicitGrantFailTk(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	//mydb.MockInsertSuccess1 = true
	//mydb.MockInsertID1 = 5

	//mydb.MockInsertSuccess2 = true
	//mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 7

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ig odb.ImplicitGrant
	ig.ClientID = cidIg
	ig.UserID = "1234"
	var scope = []string{"test1", "test2"}
	fmt.Println("last add -------------------------------------------------------")
	res, igid := odbIg.AddImplicitGrant(&ig, &at, &scope)
	if res || igid != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_AddImplicitGrantFailIg(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	//mydb.MockInsertSuccess2 = true
	//mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 7

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ig odb.ImplicitGrant
	ig.ClientID = cidIg
	ig.UserID = "1234"
	var scope = []string{"test1", "test2"}
	fmt.Println("last add -------------------------------------------------------")
	res, igid := odbIg.AddImplicitGrant(&ig, &at, &scope)
	if res || igid != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_AddImplicitGrantFailScope(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 6

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 7

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true
	mydb.MockDeleteSuccess4 = true

	// var getRow db.DbRow
	// getRow.Row = []string{"1", "2", "user", "4", ""}
	// mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{"1", "2", "user", "4", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	mydb.MockRows2 = &dbrows
	mydb.MockRows3 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var ig odb.ImplicitGrant
	ig.ClientID = cidIg
	ig.UserID = "1234"
	var scope = []string{"test1", "test2"}
	fmt.Println("last add -------------------------------------------------------")
	res, _ := odbIg.AddImplicitGrant(&ig, &at, &scope)
	fmt.Println("res in AddImplicitGrantFailScope: ", res)
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_DeleteImplicitGrant(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5

	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 6

	// mydb.MockInsertSuccess3 = true
	// mydb.MockInsertID3 = 7

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true
	mydb.MockDeleteSuccess4 = true

	// var getRow db.DbRow
	// getRow.Row = []string{"1", "2", "user", "4", ""}
	// mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{"1", "2", "user", "4", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	mydb.MockRows2 = &dbrows
	mydb.MockRows3 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	res := odbIg.DeleteImplicitGrant(cidIg, "1234")
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_DeleteImplicitGrant2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5

	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 6

	// mydb.MockInsertSuccess3 = true
	// mydb.MockInsertID3 = 7

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	mydb.MockDeleteSuccess3 = true
	mydb.MockDeleteSuccess4 = true

	// var getRow db.DbRow
	// getRow.Row = []string{"1", "2", "user", "4", ""}
	// mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	fmt.Println("rows in test", mydb.MockRows1)
	fmt.Println("rows in test len", len(mydb.MockRows1.Rows))
	//mydb.MockRows2 = &dbrows
	//mydb.MockRows3 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	res := odbIg.DeleteImplicitGrant(cidIg, "1234")
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_DeleteImplicitGrantFail1(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5

	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 6

	// mydb.MockInsertSuccess3 = true
	// mydb.MockInsertID3 = 7

	// mydb.MockDeleteSuccess1 = true
	// mydb.MockDeleteSuccess2 = true
	// mydb.MockDeleteSuccess3 = true
	// mydb.MockDeleteSuccess4 = true

	// var getRow db.DbRow
	// getRow.Row = []string{"1", "2", "user", "4", ""}
	// mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{"1", "2", "user", "4", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	mydb.MockRows2 = &dbrows
	mydb.MockRows3 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	res := odbIg.DeleteImplicitGrant(cidIg, "1234")
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_DeleteImplicitGrantFail2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5

	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 6

	// mydb.MockInsertSuccess3 = true
	// mydb.MockInsertID3 = 7

	mydb.MockDeleteSuccess1 = true
	// mydb.MockDeleteSuccess2 = true
	// mydb.MockDeleteSuccess3 = true
	// mydb.MockDeleteSuccess4 = true

	// var getRow db.DbRow
	// getRow.Row = []string{"1", "2", "user", "4", ""}
	// mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{"1", "2", "user", "4", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	mydb.MockRows2 = &dbrows
	mydb.MockRows3 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	res := odbIg.DeleteImplicitGrant(cidIg, "1234")
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBIg_DeleteImplicitGrantFail3(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbIg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5

	// mydb.MockInsertSuccess2 = true
	// mydb.MockInsertID2 = 6

	// mydb.MockInsertSuccess3 = true
	// mydb.MockInsertID3 = 7

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = true
	// mydb.MockDeleteSuccess3 = true
	// mydb.MockDeleteSuccess4 = true

	// var getRow db.DbRow
	// getRow.Row = []string{"1", "2", "user", "4", ""}
	// mydb.MockRow1 = &getRow

	var rows [][]string
	row1 := []string{"1", "2", "user", "4", ""}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	// var rows2 [][]string
	// row2 := []string{"1", "scope1", "4"}
	// rows2 = append(rows2, row2)
	// var dbrows2 db.DbRows
	// dbrows2.Rows = rows2
	// mydb.MockRows2 = &dbrows2
	//mydb.MockRows2 = &dbrows
	//mydb.MockRows3 = &dbrows

	var moadb MySQLOauthDB
	moadb.DB = dbIg

	odbIg = &moadb

	dbIg.Connect()

	res := odbIg.DeleteImplicitGrant(cidIg, "1234")
	if res {
		t.Fail()
	}
}
