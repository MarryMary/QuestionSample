package Model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/DEVAPP")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseDB(db *sql.DB) {
	defer db.Close()
}
