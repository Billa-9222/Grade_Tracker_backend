package storage

import "gradeTracker/pkg/database"

func DeleteStudent(studentID int) (err error) {
	_, err = database.DB().Exec("delete from students where id=$1", studentID)
	if err != nil {
		return
	}
	return 
}
