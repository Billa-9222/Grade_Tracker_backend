package storage

import (
	"database/sql"
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
	responses "gradeTracker/pkg/errors"
)

func GetGradeByID(id int) (result entities.Grades, err error) {
	db := database.DB()

	row := db.QueryRow("select id, student_id, subject, score from grades where id=$1", id)
	err = row.Scan(&result.ID, &result.Student_ID, &result.Subject, &result.Score)
	if err == nil {
		if err == sql.ErrNoRows {
			return result, responses.ErrGradeNotFound
		}
		return result, err
	}

	return result, nil
}	
	