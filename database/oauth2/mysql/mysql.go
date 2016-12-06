package oauth2mysql

import (
	"database/sql"
	"fmt"
	"github.com/Ulbora/oauth2server/domain"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var returnCode int = 0

func Initialize() int {

	db, err = sql.Open("mysql", "admin:admin@/ulbora_oauth2_server")
	if err != nil {
		returnCode = 1
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		returnCode = 1
		panic(err.Error())
	}
	return returnCode
}

func AddClient(c *domain.Client) bool {
	rtn := true
	fmt.Println(c.Name)
	stmtIns, err := db.Prepare(ClientInsertQuery) // ? = placeholder
	if err != nil {
		rtn = false
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(c.Secret, c.RedirectUrl, c.Name, c.WebSite, c.Email, c.Enabled) // Insert tuples (i, i^2)
	if err != nil {
		rtn = false
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return rtn
}

func CloseDb() int {
	rtn := 1
	db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		rtn = 0
	}
	return rtn
}
