package handler

import (
	"encoding/json"
	"gradeTracker/internal/entities"
	"gradeTracker/internal/service"
	responses "gradeTracker/pkg/errors"
	httpResponser "gradeTracker/pkg/http"
	"net/http"
)

func UpdateGrades(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	var grade entities.Grades
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&grade)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return 
	}

	err = service.UpdateGrade(grade)
	if err != nil {
		resp.Message = err.Error()
		return 
	}

	resp.Message = responses.ErrSuccess.Error()
	resp.Payload = grade
}