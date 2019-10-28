package managers

import (
	"fmt"
	"testing"

	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//var token string

func TestOauthManagerRefToken_GenerateRefreshToken(t *testing.T) {

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
	mGetRow.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var m OauthManager
	m.Db = odbAu
	token := m.GenerateRefreshToken(125, "tester1", "code")
	if token == "" {
		t.Fail()
	} else {
		fmt.Println("refreshToken in test: ", token)
	}

	valid, val := m.Validate(token, "6g651dfg6gf6")
	if !valid || val.UserID != "tester1" {
		t.Fail()
	}
}
