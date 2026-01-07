package errors

import "errors"

var (
	ErrInternalServer = errors.New("Internal server error")
	ErrBadRequest = errors.New("Incorrect request")
	ErrSuccess = errors.New("Successful")
	ErrAccessDenied = errors.New("Access denied")
	ErrStudentExists = errors.New("Student already exists")
	ErrStudentNotFound = errors.New("Student not found")
	ErrMethodNotAllowed = errors.New("This method is not allowed")
	ErrGradeNotFound = errors.New("Grade not found")
)

var StatusCodes map[string]int = map[string]int{
	ErrInternalServer.Error(): 500,
	ErrBadRequest.Error(): 400,
	ErrSuccess.Error(): 200,
	ErrAccessDenied.Error(): 403,
	ErrStudentExists.Error(): 400,
	ErrStudentNotFound.Error(): 404,
	ErrMethodNotAllowed.Error(): 405,
	ErrGradeNotFound.Error(): 404,
}