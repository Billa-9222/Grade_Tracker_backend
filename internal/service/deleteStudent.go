package service

import (
	"gradeTracker/internal/storage"
	responses "gradeTracker/pkg/errors"
)

func DeleteStudent(studentID int) (err error) {
	if studentID <= 0 {
		return responses.ErrBadRequest
	}

	_, err = storage.GetStudentByID(studentID)
	if err != nil {
		return err
	}

	err = storage.DeleteStudent(studentID)
	if err != nil {
		return err
	}

	return 
}