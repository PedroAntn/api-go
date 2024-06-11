package controllers

import (
	"awesomeProject4/repositories"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type ProductController struct {
	db *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{db: db}
}

func (pc *ProductController) Index(w http.ResponseWriter, r *http.Request) {
	produtos := repositories.BuscarProdutos(pc.db)
	t, err := template.ParseFiles("templates/Index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, produtos)
}

func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	repositories.CreateProduct(pc.db, nome, autor, sinopse)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	repositories.DeleteProduct(pc.db, id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (pc *ProductController) EditProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := repositories.EditProduct(pc.db, id)
	t, err := template.ParseFiles("templates/Edit.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, product)
}

func (pc *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Printf("Error converting id: %v", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	repositories.UpdateProduct(pc.db, id, nome, autor, sinopse)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (pc *ProductController) Insert(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	repositories.CreateProduct(pc.db, nome, autor, sinopse)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (pc *ProductController) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	cliente := repositories.LoginCliente(pc.db, email, password)
	if cliente != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		http.ServeFile(w, r, "templates/Login.html")
	}
}
