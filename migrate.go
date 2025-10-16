package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Migration struct {
	Up   string `json:"up"`
	Down string `json:"down"`
}

func getMigration() Migration {
	data, err := os.ReadFile("migrations/account.json")
	if err != nil {
		panic(err)
	}

	var migration Migration
	if err = json.Unmarshal(data, &migration); err != nil {
		fmt.Println(err.Error())
	}
	return migration
}

func handleMigrationFileError(err error) {
	if err != nil {
		panic(err)
	}
}

func MigrateUp() {
	migration := getMigration()
	// Execute the "up" SQL command to create the accounts table.
	_, err := db.Exec(migration.Up)
	handleMigrationFileError(err)
}

func MigrateDown() {
	migration := getMigration()
	// Execute the "down" SQL command to drop the accounts table.
	_, err := db.Exec(migration.Down)
	handleMigrationFileError(err)
}
