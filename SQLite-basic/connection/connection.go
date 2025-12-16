// SQLite-basic/connection/connection.go
package connection

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite", "database.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}

	return db

}

// Close closes the database connection
func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatalln(err)
	}
}
