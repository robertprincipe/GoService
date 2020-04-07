package database

import (
	"fmt"
	"database/sql"
)

// InitConnection devuelve una coneción a la base de datos
func InitConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", "root", "", "northwind")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}