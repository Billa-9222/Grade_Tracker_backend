package handler

import (
	"gradeTracker/internal/service"
	responses "gradeTracker/pkg/errors"
	httpResponser "gradeTracker/pkg/http"
	"net/http"
	"strconv"
)

func GetGrades(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	var (
		studentID, gradeID int
		err error
	)

	idIs := r.URL.Query().Get("id")
	if idIs != "" {
		gradeID, err = strconv.Atoi(idIs)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return 
		}

		grade, err := service.GetGradeByID(gradeID)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return 
		}
		resp.Message = responses.ErrSuccess.Error()
		resp.Payload = grade
		return 
	}

	stuIdIs := r.URL.Query().Get("student_id")
	if stuIdIs != "" {
		studentID, err = strconv.Atoi(stuIdIs)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return 
		}
	}

	grades, err := service.GetGrades(studentID)
	if err != nil {
		resp.Message = err.Error()
		return 
	}
	resp.Message = responses.ErrSuccess.Error()
	resp.Payload = grades
}