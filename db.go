package main

import (
	"database/sql"
	"fmt"
	"log"
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

func InsertInitialAccountData(email string, accountId string) {
	// Here status will default to NOT_CONNECTED
	if _, err := db.Exec("INSERT INTO accounts(email, account_id, status) VALUES (?, ?, ?)", email, accountId, "NOT_CONNECTED"); err != nil {
		log.Print(err.Error())
	}
}

func UpdateAccountData(accountId string, status string) {
	// To be called from webhook handler and update records with the status update that we receive
	if _, err := db.Exec("UPDATE accounts SET status = ? where account_Id = ?", status, accountId); err != nil {
		log.Print(err.Error())
	}
}

func SelectAccountData(email string) Account {
	res := db.QueryRow("SELECT id, email, account_id, status FROM accounts where email = ?", email)
	acc := Account{}
	if err := res.Scan(&acc.Id, &acc.Email, &acc.AccountId, &acc.Status); err != nil {
		log.Print(err.Error())
	}
	return acc
}
