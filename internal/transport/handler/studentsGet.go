package handler

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/service"
	responses "gradeTracker/pkg/errors"
	httpResponser "gradeTracker/pkg/http"
	"net/http"
	"strconv"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	var student entities.Students
	var err error
	idIs := r.URL.Query().Get("id")
	if idIs != "" {
		student.ID, err = strconv.Atoi(idIs)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return
		}
	}
	student.Name = r.URL.Query().Get("name")
	student.Surname = r.URL.Query().Get("surname")
	student.Class = r.URL.Query().Get("class")

	limit := 10
	offset := 0

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return 
		}
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return 
		}
	}

	students, err := service.GetStudent(student, limit, offset)
	if err != nil {
		resp.Message = err.Error()
		return 
	}
	resp.Message = responses.ErrSuccess.Error()
	resp.Payload = students
}