package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DB() *sql.DB {
	if db != nil {
		return db
	}

	st := "user=postgres password=Bill@+2999 dbname=grade_tracker sslmode=disable"

	var err error
	db, err := sql.Open("postgres", st)
	if err != nil {
		panic(err)
	}
	return db
}