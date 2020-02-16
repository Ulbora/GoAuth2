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
	var l lg.Logger
	m.Log = &l
	moadb.Log = &l
	m.Db = odbAu
	token := m.GenerateRefreshToken(125, "tester1", "code")
	if token == "" {
		t.Fail()
	} else {
		fmt.Println("refreshToken in test: ", token)
	}

	valid, val := m.ValidateJwt(token, "6g651dfg6gf6")
	if !valid || val.UserID != "tester1" {
		t.Fail()
	}
}

func TestOauthManagerRefToken_GenerateRefreshTokenEnvVar(t *testing.T) {

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
	var tkp TokenParams
	tkp.RefreshTokenKey = "12345"
	tkp.Issuer = "testIss"
	tkp.Audience = "testAud"
	m.TokenParams = &tkp
	var l lg.Logger
	m.Log = &l
	moadb.Log = &l
	m.Db = odbAu
	token := m.GenerateRefreshToken(125, "tester1", "code")
	if token == "" {
		t.Fail()
	} else {
		fmt.Println("refreshToken in test: ", token)
	}

	valid, val := m.ValidateJwt(token, "12345")
	if !valid || val.UserID != "tester1" {
		t.Fail()
	}
}
