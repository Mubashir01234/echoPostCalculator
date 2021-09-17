package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
func Conc() *sql.DB {
	db, err := sql.Open("mysql", "root:01234@tcp(127.0.0.1:3306)/calculator")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	return db

}

