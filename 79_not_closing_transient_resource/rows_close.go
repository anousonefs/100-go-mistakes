package main

import (
	"database/sql"
	"log"
)

// rows.Close()
/* Forgetting to close the rows means a connection leak, which prevents the database connection from being put back into the connection pool. */

var db *sql.DB

func fetchUsers() {
	for i := 0; i < 1000; i++ {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			log.Fatal(err)
		}
		_ = rows
		// Forgetting to call rows.Close()
		// This causes the connection to stay open
	}
}
