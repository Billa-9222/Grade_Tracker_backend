package httpResponser

import (
	"encoding/json"
	responses "gradeTracker/pkg/errors"
	"net/http"
)

type Response  struct {
	Message string `json:"message"`
	Payload interface{} `json:"payload"`
}

func (r *Response) Converter(rw http.ResponseWriter) {
	statusCode := responses.StatusCodes[r.Message]
	if statusCode == 0 {
		statusCode = 500
		r.Message = responses.ErrInternalServer.Error()
	}
	res, _ := json.Marshal(r)
	rw.WriteHeader(statusCode)
	rw.Write(res)
}