package main

import (
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func createConnection() (*gorm.DB, error)  {
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"mssql",
		fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			user,password, host, port, database,
		),
	)

}
