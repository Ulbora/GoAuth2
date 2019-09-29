// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbKeyi db.Database
var odbKeyi odb.Oauth2DB
var cidKeyi int64
var idKeyi int64

func TestMySQLOauthDBKeyi_ConnectAllowURI(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbKeyi = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbKeyi

	odbKeyi = &moadb

	dbKeyi.Connect()
}

func TestMySQLOauthDBKeyi_GetAccessTokenKey(t *testing.T) {
	key := odbKeyi.GetAccessTokenKey()
	fmt.Println("access token key: ", key)
	if key == "" {
		t.Fail()
	}
}

func TestMySQLOauthDBKeyi_GetRefreshTokenKey(t *testing.T) {
	key := odbKeyi.GetRefreshTokenKey()
	fmt.Println("refresh token key: ", key)
	if key == "" {
		t.Fail()
	}
}

func TestMySQLOauthDBKeyi_GetSessionKey(t *testing.T) {
	key := odbKeyi.GetSessionKey()
	fmt.Println("session key: ", key)
	if key == "" {
		t.Fail()
	}
}
