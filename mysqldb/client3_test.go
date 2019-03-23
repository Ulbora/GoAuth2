package mysqldb

import (
	//"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbb3 db.Database
var odbb3 odb.Oauth2DB
var cid3 int64
var cid23 int64

func TestMySQLDB3_Connect(t *testing.T) {

	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	var mTestRow db.DbRow
	mTestRow.Row = []string{"e"}
	mydb.MockTestRow = &mTestRow
	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 2

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 4

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 4

	mydb.MockInsertSuccess4 = false
	mydb.MockInsertID4 = 4

	mydb.MockUpdateSuccess1 = true

	var mGetRow db.DbRow
	mGetRow.Row = []string{"1", "1235", "tester5", "some site", "some email", "true", "false"}
	mydb.MockRow1 = &mGetRow

	var rows [][]string
	row1 := []string{"1", "1235", "tester5", "some site", "some email", "true", "false"}
	rows = append(rows, row1)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	mydb.MockRows2 = &dbrows
	//mGetRows.Rows[0] = []string{"1", "1235", "tester5", "some site", "some email", "true", "false"}

	mydb.MockDeleteSuccess1 = true
	mydb.MockDeleteSuccess2 = false
	mydb.MockDeleteSuccess3 = false
	mydb.MockDeleteSuccess4 = true

	dbb3 = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbb3

	odbb3 = &moadb

	dbb3.Connect()

}

// func TestMySQLDB_AddClientNullUri(t *testing.T) {
// 	var c odb.Client
// 	c.Secret = "12345"
// 	c.Name = "tester"
// 	c.Email = "bob@bob.com"
// 	c.WebSite = "www.bob.com"
// 	c.Enabled = true
// 	c.Paid = false

// 	fmt.Println("before db add")
// 	res, id := odbb.AddClient(&c, nil)
// 	fmt.Println("res: ", res)
// 	fmt.Println("id: ", id)
// 	if res || id == 0 {
// 		t.Fail()
// 	} else {
// 		cid = id
// 	}
// }

// func TestMySQLDB3_AddClient(t *testing.T) {
// 	var c odb.Client
// 	c.Secret = "1234567"
// 	c.Name = "tester"
// 	c.Email = "bob@bob.com"
// 	c.WebSite = "www.bob.com"
// 	c.Enabled = true
// 	c.Paid = false
// 	var uis []odb.ClientRedirectURI
// 	var u1 odb.ClientRedirectURI
// 	u1.URI = "addSomething"
// 	uis = append(uis, u1)

// 	var u2 odb.ClientRedirectURI
// 	u2.URI = "addSomething2"
// 	uis = append(uis, u2)

// 	fmt.Println("before db add")
// 	res, id := odbb_2.AddClient(&c, &uis)
// 	cid_2 = id
// 	fmt.Println("res: ", res)
// 	fmt.Println("id: ", id)
// 	if !res || id == 0 {
// 		t.Fail()
// 	}
// }

// func TestMySQLDB_UpdateClient(t *testing.T) {
// 	var c odb.Client
// 	c.Secret = "555555"
// 	c.Name = "tester5"
// 	c.Email = "bob5@bob.com"
// 	c.WebSite = "www.bob.com"
// 	c.Enabled = false
// 	c.Paid = false
// 	c.ClientID = cid
// 	suc := odbb.UpdateClient(&c)
// 	if !suc {
// 		t.Fail()
// 	}
// }

// func TestMySQLDB_GetClient(t *testing.T) {
// 	c := odbb.GetClient(cid)
// 	fmt.Println("client found: ", c)
// 	if c.Name != "tester5" {
// 		t.Fail()
// 	}
// }

// func TestMySQLDB_GetClients(t *testing.T) {
// 	cs := odbb.GetClients()
// 	fmt.Println("client found: ", cs)
// 	for _, c := range *cs {
// 		fmt.Println("client found in getClients: ", c)
// 	}
// 	if len(*cs) == 0 {
// 		t.Fail()
// 	}

// }

// func TestMySQLDB_SearchClients(t *testing.T) {
// 	cs := odbb.SearchClients("tester")
// 	fmt.Println("client found in search: ", cs)
// 	for _, c := range *cs {
// 		fmt.Println("client found in searchClients: ", c)
// 	}
// 	if len(*cs) == 0 {
// 		t.Fail()
// 	}

// }

func TestMySQLDB3_DeleteClient(t *testing.T) {
	suc := odbb3.DeleteClient(cid3)
	if suc {
		t.Fail()
	}
}

// func TestMySQLDB2_DeleteClient2(t *testing.T) {
// 	suc := odbb_2.DeleteClient(cid_2)
// 	if suc {
// 		t.Fail()
// 	}
// }

// func TestMySQLDB_DeleteClient2(t *testing.T) {
// 	suc := odbb.DeleteClient(cid2)
// 	if !suc {
// 		t.Fail()
// 	}
// }
