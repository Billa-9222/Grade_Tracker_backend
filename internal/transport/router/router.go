package router

import (
	"gradeTracker/internal/transport/handler"
	"gradeTracker/internal/transport/mw"

	"github.com/gorilla/mux"
)

func NewRouterComp() *mux.Router {
	router := mux.NewRouter()
	
    router.Use(mw.CORS)
	router.Use(mw.Time)

	router.HandleFunc("/students", handler.CreateStudent).Methods("POST")
	router.HandleFunc("/students", handler.GetStudent).Methods("GET")
	router.HandleFunc("/students", handler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students", handler.DeleteStudent).Methods("DELETE")

	router.HandleFunc("/grades", handler.CreateGrades).Methods("POST")
	router.HandleFunc("/grades", handler.GetGrades).Methods("GET")
	router.HandleFunc("/grades", handler.UpdateGrades).Methods("PUT")
	router.HandleFunc("/grades", handler.DeleteGrades).Methods("DELETE")

	return router
}