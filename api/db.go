package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB settings
const (
	dbHost = "172.17.0.2"
	dbPort = 5432
	dbUser = "postgres"
	dbName = "konsultant"
)

// ConnectToDb connects to local postges DB
func ConnectToDb() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
