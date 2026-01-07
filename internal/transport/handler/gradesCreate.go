package handler

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/service"
	responses "gradeTracker/pkg/errors"
	httpResponser "gradeTracker/pkg/http"
	"net/http"
	"strconv"
)

func CreateGrades(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	var (
		grades entities.Grades
		err error 
	)

	grades.Student_ID, err = strconv.Atoi(r.URL.Query().Get("student_id"))
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return 
	}
	grades.Subject = r.URL.Query().Get("subject")
	grades.Score, err = strconv.Atoi(r.URL.Query().Get("score"))
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return 
	}
	err = service.AssignGrades(&grades)
	if err != nil {
		resp.Message = err.Error()
		return 
	}
	resp.Message = responses.ErrSuccess.Error()
	resp.Payload = grades
 
}