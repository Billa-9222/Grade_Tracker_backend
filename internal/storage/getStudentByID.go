package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
	responses "gradeTracker/pkg/errors"
)

func GetStudentByID(id int) (students []entities.Students, err error) {
	db := database.DB()
    rows, err := db.Query("select id, name, surname, class from students where id = $1", id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var student entities.Students
        err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Class)
        if err != nil {
            return nil, err
        }
        students = append(students, student)
    }

    if len(students) == 0 {
        return nil, responses.ErrStudentNotFound
    }
    return students, nil
}