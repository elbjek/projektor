package main

import (
	"database/sql"
)

var db *sql.DB

func main() {
	db = ConnectToDb()
	RunAPI()
}
