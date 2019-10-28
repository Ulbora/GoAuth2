package managers

import (
	"fmt"
	"testing"

	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

func TestOauthManagerAccessToken_GenerateAccessToken(t *testing.T) {

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
	var pl Payload
	pl.TokenType = codeGrantType
	pl.UserID = "tester1"
	pl.ClientID = 125
	pl.Subject = "code"
	//pl.Issuer = tokenIssuer
	//pl.Audience = tokenAudience
	pl.ExpiresInMinute = codeAccessTokenLifeInMinutes //(60 * time.Minute) => (60 * 60) => 3600 minutes => 1 hours
	pl.Grant = codeGrantType
	var ruList []RoleURI
	var ru1 RoleURI
	ru1.ClientRoleID = 1
	ru1.Role = "user"
	ru1.ClientAllowedURIID = 2
	ru1.ClientAllowedURI = "test.com"
	ru1.ClientID = 5
	ruList = append(ruList, ru1)

	var ru2 RoleURI
	ru2.ClientRoleID = 12
	ru2.Role = "user"
	ru2.ClientAllowedURIID = 21
	ru2.ClientAllowedURI = "test2.com"
	ru2.ClientID = 5
	ruList = append(ruList, ru2)

	pl.RoleURIs = ruList

	pl.ScopeList = []string{"web", "rest"}
	//pl.SecretKey = rtk
	token := m.GenerateAccessToken(&pl)
	if token == "" {
		t.Fail()
	} else {
		fmt.Println("accessToken in test: ", token)
	}

	valid, val := m.Validate(token, "6g651dfg6gf6")
	if !valid || val.UserID != "tester1" {
		t.Fail()
	}
}
