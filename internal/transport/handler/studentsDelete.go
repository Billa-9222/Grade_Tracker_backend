package handler

import (
	"gradeTracker/internal/service"
	httpResponser "gradeTracker/pkg/http"
	"net/http"
	"strconv"
)


func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		resp.Message = err.Error()
		return 
	}

	err = service.DeleteStudent(id)
	if err != nil {
		resp.Message = err.Error()
		return 
	}
}