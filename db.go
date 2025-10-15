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
	sdn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false", os.Getenv("db_user"), os.Getenv("db_password"), os.Getenv("db_host"), os.Getenv("db_name"))
	fmt.Println("Connecting to database with:", sdn)
	db, err = sql.Open("mysql", sdn)
	if err != nil {
		panic(err)
	}
	return err
}
