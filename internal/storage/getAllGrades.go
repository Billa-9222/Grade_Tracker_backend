package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
)

func GetAllGrades() (grades []entities.Grades, err error) {
	db := database.DB()
	rows, err := db.Query("select id, student_id, subject, score from grades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var grade entities.Grades
		err = rows.Scan(&grade.ID, &grade.Student_ID, &grade.Subject, &grade.Score)
		if err != nil {
			return nil, err 
		}
		grades = append(grades, grade)
	}

	return grades, nil
}