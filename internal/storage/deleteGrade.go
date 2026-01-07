package storage

import "gradeTracker/pkg/database"

func DeleteGrade(gradeID int) (err error) {
	_, err = database.DB().Exec("delete from grades where id=$1", gradeID)
	if err != nil {
		return 
	}
	return 
}