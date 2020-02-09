package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbKey db.Database
var odbKey odb.Oauth2DB
var cidKey int64
var idKey int64

func TestMySQLOauthDBKey_ConnectAllowURI(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbKey = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbKey

	odbKey = &moadb

	dbKey.Connect()
}

func TestMySQLOauthDBKey_GetAccessTokenKey(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbKey = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var getRow db.DbRow
	getRow.Row = []string{"1", "somekey"}
	mydb.MockRow1 = &getRow

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbKey

	odbKey = &moadb

	dbKey.Connect()

	key := odbKey.GetAccessTokenKey()
	fmt.Println("access token key: ", key)
	if key == "" {
		t.Fail()
	}
}

func TestMySQLOauthDBKey_GetRefreshTokenKey(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbKey = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var getRow db.DbRow
	getRow.Row = []string{"1", "somekey"}
	mydb.MockRow1 = &getRow

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbKey

	odbKey = &moadb

	dbKey.Connect()

	key := odbKey.GetRefreshTokenKey()
	fmt.Println("refresh token key: ", key)
	if key == "" {
		t.Fail()
	}
}

func TestMySQLOauthDBKey_GetSessionKey(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbKey = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var getRow db.DbRow
	getRow.Row = []string{"1", "somekey"}
	mydb.MockRow1 = &getRow

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbKey

	odbKey = &moadb

	dbKey.Connect()

	key := odbKey.GetSessionKey()
	fmt.Println("session key: ", key)
	if key == "" {
		t.Fail()
	}
}
