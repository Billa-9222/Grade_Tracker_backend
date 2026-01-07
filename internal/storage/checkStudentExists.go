package storage

import (
	"database/sql"
	"gradeTracker/pkg/database"
)

var DB *sql.DB

func CheckStudentExists(studentID int) (exists bool, err error) {
    DB = database.DB()
	err = DB.QueryRow("select exists (select 1 from students where id=$1)", studentID,).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil

} 