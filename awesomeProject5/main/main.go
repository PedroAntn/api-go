package main

import (
	"awesomeProject5/database"
	"awesomeProject5/routes"
	"log"
	"net/http"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	routes.SetupRoutes(db, mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
