package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
	responses "gradeTracker/pkg/errors"
	"strconv"
) 

func GetStudentByFilters(name, surname, class string, limit, offset int) (result []entities.Students, err error) {
	db := database.DB()

	query := "select id, name, surname, class from students where 1=1"
	var info []interface{}
	placeholder := 1

	if name != "" {
		query += " and name ILIKE $" + strconv.Itoa(placeholder)
		info = append(info, name)
		placeholder++
	}
	if surname != "" {
		query += " and surname ILIKE $" + strconv.Itoa(placeholder)
		info = append(info, surname)
		placeholder++
	}
	if class != "" {
		query += " and class ILIKE $" + strconv.Itoa(placeholder)
		info = append(info, class)
		placeholder++
	}

	query += " limit $" + strconv.Itoa(placeholder) + " offset $" + strconv.Itoa(placeholder+1)
	info = append(info, limit, offset)

	rows, err := db.Query(query, info...)
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

	if len(result) == 0 {
		return nil, responses.ErrStudentNotFound
	}

	return result, nil
}