package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
	"log"
)

func CreateStudent(student *entities.Students) (err error) {
	db := database.DB()
	err = db.QueryRow(`insert into students(name, surname, class) values($1, $2, $3) returning id`, student.Name, student.Surname, student.Class).Scan(&student.ID)
	if err != nil {
		log.Println(err)
		return 
	}
	return 
}