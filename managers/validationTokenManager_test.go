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

func TestOauthManagerTakenVal_ValidateAccessToken(t *testing.T) {

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

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	man.Db = odbAu
	var l lg.Logger
	man.Log = &l
	moadb.Log = &l
	var m Manager
	m = &man

	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = passwordGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = passwordGrantType
	pl.Issuer = tokenIssuer
	pl.ScopeList = []string{"write"}
	//pl.RoleURIs = *m.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	fmt.Println("accessToken: ", accessToken)

	var at ValidateAccessTokenReq
	at.UserID = hashUser("tester1")
	at.Hashed = true
	at.AccessToken = accessToken
	at.ClientID = 2
	suc := m.ValidateAccessToken(&at)
	fmt.Println("suc: ", suc)
	//fmt.Println("tkn: ", tkn)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerTakenVal_ValidateAccessToken2(t *testing.T) {

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
	row := []string{"4", "somerole", "1", "someurl", "2"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

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

	roleURIList := man.Db.GetClientRoleAllowedURIListByClientID(2)
	fmt.Println("roleURIList", roleURIList)
	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = passwordGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = passwordGrantType
	pl.Issuer = tokenIssuer
	pl.ScopeList = []string{"web"}

	pl.RoleURIs = *man.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	fmt.Println("accessToken: ", accessToken)

	var at ValidateAccessTokenReq
	at.UserID = "tester1"
	at.Hashed = false
	at.AccessToken = accessToken
	at.ClientID = 2
	at.Scope = "web"
	at.Role = "somerole"
	at.URI = "someurl"
	suc := m.ValidateAccessToken(&at)
	fmt.Println("suc: ", suc)
	//fmt.Println("tkn: ", tkn)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerTakenVal_ValidateAccessToken3(t *testing.T) {

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
	row := []string{"4", "somerole", "1", "someurl", "2"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

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

	roleURIList := man.Db.GetClientRoleAllowedURIListByClientID(2)
	fmt.Println("roleURIList", roleURIList)
	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = implicitGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = implicitGrantType
	pl.Issuer = tokenIssuer
	pl.ScopeList = []string{"web"}

	pl.RoleURIs = *man.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	fmt.Println("accessToken: ", accessToken)

	var at ValidateAccessTokenReq
	at.UserID = "tester1"
	at.Hashed = false
	at.AccessToken = accessToken
	at.ClientID = 2
	at.Scope = "web"
	at.Role = "somerole"
	at.URI = "someurl"
	suc := m.ValidateAccessToken(&at)
	fmt.Println("suc: ", suc)
	//fmt.Println("tkn: ", tkn)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerTakenVal_ValidateAccessToken3a(t *testing.T) {

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
	row := []string{"4", "somerole", "1", "someurl", "2"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

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

	roleURIList := man.Db.GetClientRoleAllowedURIListByClientID(2)
	fmt.Println("roleURIList", roleURIList)
	var pl Payload
	pl.TokenType = accessTokenType
	pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = implicitGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = implicitGrantType
	pl.Issuer = tokenIssuer
	pl.ScopeList = []string{"write"}

	pl.RoleURIs = *man.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	fmt.Println("accessToken: ", accessToken)

	var at ValidateAccessTokenReq
	at.UserID = "tester1"
	at.Hashed = false
	at.AccessToken = accessToken
	at.ClientID = 2
	at.Scope = "read"
	at.Role = "somerole"
	at.URI = "someurl"
	suc := m.ValidateAccessToken(&at)
	fmt.Println("suc: ", suc)
	//fmt.Println("tkn: ", tkn)
	if !suc {
		t.Fail()
	}
}

func TestOauthManagerTakenVal_ValidateAccessToken4(t *testing.T) {

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
	row := []string{"4", "somerole", "1", "someurl", "2"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var mGetRow1 db.DbRow
	mGetRow1.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow1 = &mGetRow1

	var mGetRow2 db.DbRow
	mGetRow2.Row = []string{"2", "6g651dfg6gf6"}
	mydb.MockRow2 = &mGetRow2

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

	roleURIList := man.Db.GetClientRoleAllowedURIListByClientID(2)
	fmt.Println("roleURIList", roleURIList)
	var pl Payload
	pl.TokenType = accessTokenType
	//pl.UserID = hashUser("tester1")
	pl.ClientID = 2
	pl.Subject = clientGrantType
	pl.ExpiresInMinute = passwordGrantAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = clientGrantType
	pl.Issuer = tokenIssuer
	pl.ScopeList = []string{"web"}

	pl.RoleURIs = *man.populateRoleURLList(roleURIList)
	//pl.ScopeList = *scopeStrList
	accessToken := man.GenerateAccessToken(&pl)

	fmt.Println("accessToken: ", accessToken)

	var at ValidateAccessTokenReq
	//at.UserID = "tester1"
	at.Hashed = false
	at.AccessToken = accessToken
	at.ClientID = 2
	at.Scope = "web"
	at.Role = "somerole"
	at.URI = "someurl"
	suc := m.ValidateAccessToken(&at)
	fmt.Println("suc: ", suc)
	//fmt.Println("tkn: ", tkn)
	if !suc {
		t.Fail()
	}
}
