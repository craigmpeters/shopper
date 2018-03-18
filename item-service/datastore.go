package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"os"
)

func createConnection() (*gorm.DB, error) {
	// Build a SQL Connection String
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"mssql",
		fmt.Sprintf("server=%s;item id=%s;password=%s;port=%s;database=%s;",
			host, user, password, port, database,
		),
	)
}
