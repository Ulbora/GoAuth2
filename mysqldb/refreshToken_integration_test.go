// +build integration move to top

package mysqldb

import (
	"fmt"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"testing"
)

var dbRti db.Database
var odbRti odb.Oauth2DB
var idRti int64

//var cidRti int64

func TestMySQLOauthDBReTokeni_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbRti = &mydb

	var moadb MySQLOauthDB
	moadb.DB = dbRti

	odbRti = &moadb

	dbRti.Connect()
}

func TestMySQLOauthDBReTokeni_AddRefreshToken(t *testing.T) {
	var tk odb.RefreshToken
	tk.Token = "somereftoken"
	res, id := odbRti.AddRefreshToken(&tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idRti = id
	}
}

func TestMySQLOauthDBReTokeni_UpdateRefreshToken(t *testing.T) {
	var tk odb.RefreshToken
	tk.ID = idRti
	tk.Token = "somereftoken2"
	res := odbRti.UpdateRefreshToken(&tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBReTokeni_GetRefreshToken(t *testing.T) {
	res := odbRti.GetRefreshToken(idRti)
	fmt.Println("ref token: ", res)
	if res == nil || (*res).Token != "somereftoken2" {
		t.Fail()
	}
}

func TestMySQLOauthDBReTokeni_DeleteRefreshToken(t *testing.T) {
	res := odbRti.DeleteRefreshToken(idRti)
	fmt.Println("del ref token: ", res)
	if !res {
		t.Fail()
	}
}
