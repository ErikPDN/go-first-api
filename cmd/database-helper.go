package main

import (
	"database/sql"
	"go-api/internal/database"
)

var db *sql.DB

func initDatabase() error {
	var err error
	config := database.DefaultPostgresConfig()
	db, err = database.NewPostgresConnection(config)
	return err
}

func closeDatabase() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
