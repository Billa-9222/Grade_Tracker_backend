package handler

import (
	"gradeTracker/internal/service"
	responses "gradeTracker/pkg/errors"
	httpResponser "gradeTracker/pkg/http"
	"net/http"
	"strconv"
)

func DeleteGrades(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	id, err := strconv.Atoi(r.URL.Query().Get("id")) 
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return 
	}
	err = service.DeleteGrade(id)
	if err != nil {
		resp.Message = err.Error()
		return 
	}
	resp.Message = responses.ErrSuccess.Error()
}