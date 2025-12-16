package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Connect() *sql.DB {
	conn, err := sql.Open("sqlite", "database.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func Close(conn *sql.DB) {
	if err := conn.Close(); err != nil {
		log.Fatalln(err)
	}
}
