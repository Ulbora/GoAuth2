package oauth2mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	 "fmt"
	 //"github.com/Ulbora/oauth2server/domain"
)
var db *sql.DB
var err error
var returnCode int = 0
func Initialize() int{
	
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