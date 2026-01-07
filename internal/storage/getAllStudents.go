package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
	responses "gradeTracker/pkg/errors"
)

func GetAllStudents(limit, offset int) (result []entities.Students, err error) {
	db := database.DB()

	rows, err := db.Query("select id, name, surname, class from students limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student entities.Students
		err = rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Class)
		if err != nil {
			return nil, err
		}
		result = append(result, student)
	}

	if len(result)== 0 {
		return nil, responses.ErrStudentNotFound
	}

	return result, nil
}