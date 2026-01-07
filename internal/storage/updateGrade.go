package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
)

func UpdateGrade(grade entities.Grades) (err error) {
	db := database.DB()

	_, err = db.Exec("update grades set subject=$1, score=$2 where id=$3", grade.Subject, grade.Score, grade.ID)
	return err
}