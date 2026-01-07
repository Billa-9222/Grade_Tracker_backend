package handler

import (
	"gradeTracker/internal/entities"
	"gradeTracker/internal/service"
	responses "gradeTracker/pkg/errors"
	httpResponser "gradeTracker/pkg/http"
	"net/http"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	var student entities.Students

	student.Name = r.URL.Query().Get("name")
	student.Surname = r.URL.Query().Get("surname")
	student.Class = r.URL.Query().Get("class")

	err := service.CreateStudent(&student)
	if err != nil {
		resp.Message = err.Error()
		return 
	}

	resp.Message = responses.ErrSuccess.Error()
	resp.Payload = student

}