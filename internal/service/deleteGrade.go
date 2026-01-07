package service

import (
	"gradeTracker/internal/storage"
	responses "gradeTracker/pkg/errors"
)

func DeleteGrade(gradeID int) (err error) {
	if gradeID <= 0 {
		return responses.ErrBadRequest
	}
	
	_, err = storage.GetGradeByID(gradeID)
	if err != nil {
		return responses.ErrGradeNotFound
	}


	err = storage.DeleteGrade(gradeID)
	if err != nil {
		return err
	}

	return
}