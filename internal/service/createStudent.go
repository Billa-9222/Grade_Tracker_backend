package service

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/storage"
	responses "gradeTracker/pkg/errors"
)

func CreateStudent(student *entities.Students) (err error) {
	if student.Name == "" || student.Surname == "" || student.Class == "" {
		err = responses.ErrBadRequest
		return 
	}

	if !isValidName(student.Name) {
		return responses.ErrBadRequest
	}

	if !isValidSurname(student.Surname) {
		return responses.ErrBadRequest
	}

	if !isValidClass(student.Class) {
		return responses.ErrBadRequest
	}

		exists, err := storage.GetStudentByFullInfo(student)
	if err != nil {
		return err
	}

	if exists {
		return responses.ErrStudentExists
	}

	
	err = storage.CreateStudent(student)
	if err != nil {
		return err
	}

	return nil
}