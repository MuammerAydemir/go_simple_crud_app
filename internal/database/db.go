package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/database-name"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Database connection error:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Database ping error:", err)
		return nil, err
	}

	return db, nil
}
