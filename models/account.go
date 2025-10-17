package models

import (
	"encoding/json"
	"log"
	"os"
)

type DBAccountModel struct {
	GetAccountByEmail  string
	InsertAccount      string
	UpdateAccountState string
}

var (
	AccountModel *DBAccountModel
)

func GetAccountDSLs() *DBAccountModel {
	f, err := os.ReadFile("sql/account.json")
	if err != nil {
		log.Panic(err)
	}
	m := DBAccountModel{}
	json.Unmarshal(f, &m)
	return &m
}
