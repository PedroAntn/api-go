package routes

import (
	"awesomeProject4/controllers"
	"awesomeProject4/repositories"
	"database/sql"
	"net/http"
	"strconv"
	"text/template"
)

func SetupRoutes(db *sql.DB, mux *http.ServeMux) {
	pc := controllers.NewProductController(db)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		produtos := repositories.BuscarProdutos(db)
		t, _ := template.New("Index").ParseFiles("templates/Index.html")
		t.Execute(w, produtos)
	})
	mux.HandleFunc("/criar-produto", pc.CreateProduct)
	mux.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Insert.html")
	})
	mux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		repositories.DeleteProduct(db, id)
		http.Redirect(w, r, "/", 301)
	})
	mux.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		product := repositories.EditProduct(db, id)
		t, _ := template.New("Edit").ParseFiles("templates/Edit.html")
		t.Execute(w, product)
	})
	mux.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.FormValue("id"))
		nome := r.FormValue("nome")
		autor := r.FormValue("autor")
		sinopse := r.FormValue("sinopse")
		repositories.UpdateProduct(db, id, nome, autor, sinopse)
		http.Redirect(w, r, "/", 301)
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Login.html")
	})
}
