package storage

import (
	"database/sql"
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
)

func GetStudentByFullInfo(student *entities.Students) (exists bool, err error) {
	db := database.DB()

	row := db.QueryRow("select id from students where name=$1 and surname=$2 and class=$3", student.Name, student.Surname, student.Class)

	var id int
	err = row.Scan(&id)

	if err == sql.ErrNoRows{
		return false, nil 
	}
	 
	if err != nil {
		return false, err
	}

	return true, nil
}