package service

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/storage"
	responses "gradeTracker/pkg/errors"
)

func UpdateGrade(grade entities.Grades) (err error) {
	prev, err := storage.GetGradeByID(grade.ID)
	if err != nil {
		return responses.ErrStudentNotFound
	}

	if grade.Subject == "" {
		grade.Subject = prev.Subject
	}

	if grade.Score == 0 {
		grade.Score = prev.Score
	}

	if grade.Score < 0 {
		return responses.ErrBadRequest
	}

	if grade.Subject == prev.Subject && grade.Score == prev.Score{
		return responses.ErrBadRequest
	}

	return storage.UpdateGrade(grade)

}