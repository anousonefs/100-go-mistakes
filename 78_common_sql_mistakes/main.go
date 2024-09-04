package main

import "database/sql"

func main() {

}

func callPing() error {
	// sql.Open may just validate its arguments without creating a connection to the database.
	db, err := sql.Open("mysql", "")
	if err != nil {
		return err
	}
	// db.Ping or PingContext forces the code to establish a connection that ensures that the data source name is valid and the database is reachable.
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

/*
Each of these parameters is an exported method of *sql.DB:
 1. SetMaxOpenConns—Maximum number of open connections to the database (default value: unlimited)
 2. SetMaxIdleConns—Maximum number of idle connections (default value: 2)
 3. SetConnMaxIdleTime—Maximum amount of time a connection can be idle before it’s closed (default value: unlimited)
 4. SetConnMaxLifetime—Maximum amount of time a connection can be held open before it’s closed (default value: unlimited)
*/

// 10.4.3 Not using prepared statements: for Efficiency and Security
// 10.4.4 Mishandling null values: use pointer or sql.Null***
// 10.4.5 Not handling row iteration errors: use rows.Err()
