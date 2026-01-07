package storage

import (
	"gradeTracker/internal/entities"
	"gradeTracker/pkg/database"
)
func UpdateStudent(student entities.Students) (err error) {
    db := database.DB()
    _, err = db.Exec("update students set name=$1, surname=$2, class=$3 where id=$4", student.Name, student.Surname, student.Class, student.ID,)
    return err
}
