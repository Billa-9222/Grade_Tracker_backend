package service

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/storage"
)

func GetGradeByID(id int) (entities.Grades, error) {
    grade, err := storage.GetGradeByID(id)
    if err != nil {
        return entities.Grades{}, err
    }

    return grade, nil
}
