package routes

import (
	"awesomeProject4/controllers"
	"database/sql"
	"net/http"
)

func SetupRoutes(db *sql.DB, mux *http.ServeMux) {
	pc := controllers.NewProductController(db)

	// Associar os métodos do controlador às rotas
	mux.HandleFunc("/", pc.Index)
	mux.HandleFunc("/criar-produto", pc.CreateProduct)
	mux.HandleFunc("/delete", pc.DeleteProduct)
	mux.HandleFunc("/edit", pc.EditProduct)
	mux.HandleFunc("/update", pc.UpdateProduct)
	mux.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Insert.html")
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Login.html")
	})

	// Servir arquivos estáticos
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}
