package service

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/storage"
	responses "gradeTracker/pkg/errors"
)

func GetGrades(studentID int) (grades []entities.Grades, err error) {
	if studentID == 0 {
		return storage.GetAllGrades()
	}

	exists, err := storage.CheckStudentExists(studentID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, responses.ErrStudentNotFound
	}

	return storage.GetGradesByStudentID(studentID)
}