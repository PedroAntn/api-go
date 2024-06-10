package main

import (
	"awesomeProject4/database"
	"awesomeProject4/routes"
	"net/http"
)

func main() {
	db := database.ConnectToDB()
	defer db.Close()

	mux := http.NewServeMux()
	routes.SetupRoutes(db, mux)

	http.ListenAndServe(":8080", mux)
}
