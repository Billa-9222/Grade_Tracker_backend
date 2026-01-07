package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
)

func AssignGrades(grade *entities.Grades) (err error) {
    db := database.DB()
    err = db.QueryRow(`insert into grades (student_id, subject, score) values ($1, $2, $3) returning id`,grade.Student_ID, grade.Subject, grade.Score,).Scan(&grade.ID)
    return err
}