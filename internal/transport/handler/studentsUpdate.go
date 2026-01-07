package handler

import (
	"encoding/json"
	"gradeTracker/internal/entities"
	"gradeTracker/internal/service"
	responses "gradeTracker/pkg/errors"
	httpResponser "gradeTracker/pkg/http"
	"log"
	"net/http"
)

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	var student entities.Students
	jsonBody := json.NewDecoder(r.Body)
	err := jsonBody.Decode(&student)
	if err != nil {
		log.Println(err)
		return 
	}

	err = service.UpdateStudent(student)
	if err != nil {
		resp.Message = err.Error()
		return 
	}

	resp.Message = responses.ErrSuccess.Error()
	resp.Payload = student

}