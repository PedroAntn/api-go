package routes

import (
	"awesomeProject5/controllers"
	"database/sql"
	"net/http"
)

func SetupRoutes(db *sql.DB, mux *http.ServeMux) {
	pc := controllers.NewProductController(db)
	mux.HandleFunc("/", pc.Index)
	mux.HandleFunc("/criar-produto", pc.CreateProduct)
	mux.HandleFunc("/insert", pc.Insert)
	mux.HandleFunc("/delete", pc.DeleteProduct)
	mux.HandleFunc("/edit", pc.EditProduct)
	mux.HandleFunc("/update", pc.UpdateProduct)
	mux.HandleFunc("/login", pc.Login)
}
