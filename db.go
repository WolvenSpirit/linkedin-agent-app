package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	START_CONNECTION_AFTER_SECONDS = 15
)

func DBConnect() error {
	var err error

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("db_user")
	cfg.Passwd = os.Getenv("db_password")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("db_host")
	cfg.DBName = os.Getenv("db_name")
	time.Sleep(START_CONNECTION_AFTER_SECONDS * time.Second)
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Printf("Attempting to test connection encountered %s", err.Error())
	} else {
		fmt.Print("Connected to DB")
	}
	return err
}
