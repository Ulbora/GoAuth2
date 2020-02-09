// +build integration move to top

package mysqldb

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbAti db.Database
var odbAti odb.Oauth2DB
var idAti int64
var refTkIdi int64

//var cidRti int64

func TestMySQLOauthDBAcTokeni_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAti = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbAti

	odbAti = &moadb

	dbAti.Connect()
}

func TestMySQLOauthDBAcTokeni_AddAccessToken(t *testing.T) {
	var tk odb.AccessToken
	tk.Token = "someacctoken"
	tk.Expires = time.Now()
	res, id := odbAti.AddAccessToken(nil, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAti = id
	}
}

func TestMySQLOauthDBAcTokeni_UpdateAccessToken(t *testing.T) {
	var tk odb.AccessToken
	tk.ID = idAti
	tk.Token = "someacctoken2"
	tk.Expires = time.Now()
	res := odbAti.UpdateAccessToken(nil, &tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcTokeni_GetAccessToken(t *testing.T) {
	res := odbAti.GetAccessToken(idAti)
	fmt.Println("access token: ", res)
	fmt.Println("access token refTokenId: ", res.RefreshTokenID)
	if res == nil || (*res).Token != "someacctoken2" {
		t.Fail()
	}
}

func TestMySQLOauthDBAcTokeni_DeleteAccessToken(t *testing.T) {
	res := odbAti.DeleteAccessToken(nil, idAti)
	fmt.Println("del access token: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcTokeni_AddRefreshToken(t *testing.T) {
	var tk odb.RefreshToken
	tk.Token = "somereftoken"
	res, id := odbAti.AddRefreshToken(nil, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		refTkIdi = id
	}
}

func TestMySQLOauthDBAcTokeni_AddAccessToken2(t *testing.T) {
	var tk odb.AccessToken
	tk.Token = "someacctoken"
	tk.Expires = time.Now()
	tk.RefreshTokenID = refTkIdi
	res, id := odbAti.AddAccessToken(nil, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idAti = id
	}
}

func TestMySQLOauthDBAcTokeni_UpdateAccessToken2(t *testing.T) {
	var tk odb.AccessToken
	tk.ID = idAti
	tk.Token = "someacctoken2"
	tk.Expires = time.Now()
	tk.RefreshTokenID = refTkIdi
	res := odbAti.UpdateAccessToken(nil, &tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcTokeni_GetAccessToken2(t *testing.T) {
	res := odbAti.GetAccessToken(idAti)
	fmt.Println("access token: ", res)
	fmt.Println("access token refTokenId: ", res.RefreshTokenID)
	if res == nil || (*res).Token != "someacctoken2" {
		t.Fail()
	}
}

func TestMySQLOauthDBAcTokeni_DeleteAccessToken2(t *testing.T) {
	res := odbAti.DeleteAccessToken(nil, idAti)
	fmt.Println("del access token: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBAcTokeni_DeleteRefreshToken(t *testing.T) {
	res := odbAti.DeleteRefreshToken(nil, refTkIdi)
	fmt.Println("del ref token: ", res)
	if !res {
		t.Fail()
	}
}
