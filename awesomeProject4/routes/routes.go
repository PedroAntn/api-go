package routes

import (
	"awesomeProject4/controllers"
	"database/sql"
	"net/http"
)

func SetupRoutes(db *sql.DB, mux *http.ServeMux) {
	pc := controllers.NewProductController(db)
	mux.HandleFunc("/api/products", pc.Index)
	mux.HandleFunc("/api/products/create", pc.CreateProduct)
	mux.HandleFunc("/api/products/insert", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Insert.html")
	})
	mux.HandleFunc("/api/products/delete", pc.DeleteProduct)
	mux.HandleFunc("/api/products/edit", pc.EditProduct)
	mux.HandleFunc("/api/products/update", pc.UpdateProduct)
	mux.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Login.html")
	})
}
