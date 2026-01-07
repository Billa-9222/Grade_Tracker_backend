package main

import (
	"fmt"
	"gradeTracker/internal/transport/router"
	"gradeTracker/pkg/database"

	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	
	db := database.DB()
	defer db.Close()

	fmt.Println("listening localhost:2025")
	http.ListenAndServe("localhost:2025", router.NewRouterComp())
}
