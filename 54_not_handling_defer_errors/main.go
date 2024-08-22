package main

import (
	"database/sql"
	"log"
)

func main() {

}

const query = "..."

func getBalance(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close() // not handler error
	return 0, nil
}

// handler defer error
func getBalance2(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()
	return 0, nil
}

// not compile
/* func getBalance3(db *sql.DB, clientID string) (float32, error) { */
/* 	rows, err := db.Query(query, clientID) */
/* 	if err != nil { */
/* 		return 0, err */
/* 	} */
/* 	defer func() { */
/* 		err := rows.Close() */
/* 		if err != nil { */
/* 			return  err */
/* 		} */
/* 	}() */
/* 	return 0, nil */
/* } */

// bug
func getBalance4(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = rows.Close() // alway return error from rows.Close
	}()
	return 0, nil
}

// errors should always be handled
func getBalance5(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil {
			if closeErr != nil {
				log.Printf("failed to close rows: %v", err)
			}
			return
		}
		err = closeErr
	}()
	return 0, nil
}
