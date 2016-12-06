package oauth2mysql

import (
	"database/sql"
	"database/sql/driver"
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

func AddClient(c *domain.Client) uint64 {
	var insertId int64 = 0
	var rtnId uint64 = 0
	fmt.Println(c.Name)
	stmtIns, err := db.Prepare(ClientInsertQuery) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var result driver.Result
	result, err = stmtIns.Exec(c.Secret, c.RedirectUrl, c.Name, c.WebSite, c.Email, c.Enabled) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insertId, _ = result.LastInsertId()
	rtnId = uint64(insertId)
	return rtnId
}

func GetClient(cid uint64) *domain.Client {
	stmtOut, err := db.Prepare(ClientReadRecordQuery) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	c := domain.Client{}
	err = stmtOut.QueryRow(cid).Scan(&c.ClientId, &c.Secret, &c.RedirectUrl, &c.Name, &c.WebSite, &c.Email, &c.Enabled) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return &c
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
