package service

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/storage"
	responses "gradeTracker/pkg/errors"
)

func UpdateStudent(student entities.Students) (err error) {
	students, err := storage.GetStudentByID(student.ID)
	if err != nil {
		return responses.ErrStudentNotFound
	}
	prev := students[0]

	if student.Name == "" {
		student.Name = prev.Name
	}

	if student.Surname == "" {
		student.Surname = prev.Surname
	}

	if student.Class == "" {
		student.Class = prev.Class
	}

	if student.Name == prev.Name && student.Surname == prev.Surname && student.Class == prev.Class {
		return responses.ErrBadRequest
	}

	return storage.UpdateStudent(student)
}