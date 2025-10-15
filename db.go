package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBConnect() error {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("db_user"), os.Getenv("db_password"), os.Getenv("db_host"), os.Getenv("db_name")))
	if err != nil {
		panic(err)
	}
	return err
}
