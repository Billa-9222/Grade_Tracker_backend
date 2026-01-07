package service

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/storage"
	responses "gradeTracker/pkg/errors"
)


func AssignGrades(grades *entities.Grades) (err error) {
	if grades.Student_ID == 0 || grades.Subject == "" || grades.Score == 0 {
		err = responses.ErrBadRequest
		return 
	}
	
	exists, err := storage.CheckStudentExists(grades.Student_ID)
	if err != nil {
		return err
	}
	if !exists {
		return responses.ErrStudentNotFound
	}

	err = storage.AssignGrades(grades)
	if err != nil {
		return err
	}

	return nil
}