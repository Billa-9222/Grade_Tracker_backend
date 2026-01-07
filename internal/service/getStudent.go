package service

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/storage"
) 


func GetStudent(student entities.Students, limit, offset int) (students []entities.Students, err error) {
	if student.ID > 0 {
	    students, err = storage.GetStudentByID(student.ID)
		if err != nil{
			return nil, err
		}
		return students, nil
	}

	 if student.Name != "" || student.Surname != "" || student.Class != "" {
		students, err= storage.GetStudentByFilters(student.Name, student.Surname, student.Class, limit, offset)
		if err != nil {
			return nil, err
		}
		return students, nil

	} 

	students, err = storage.GetAllStudents(limit, offset)
	if err != nil {
		return nil, err
	}
	return students, nil
  }
