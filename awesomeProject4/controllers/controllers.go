package controllers

import (
	"awesomeProject4/repositories"
	"database/sql"
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
	t, _ := template.New("Index").ParseFiles("templates/Index.html")
	t.Execute(w, produtos)
}

func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	repositories.CreateProdict(pc.db, nome, autor, sinopse)
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	repositories.DeleteProduct(pc.db, id)
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) EditProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := repositories.EditProduct(pc.db, id)
	t, _ := template.New("Edit").ParseFiles("templates/Edit.html")
	t.Execute(w, product)
}

func (pc *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	repositories.UpdateProduct(pc.db, id, nome, autor, sinopse)
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) Insert(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	repositories.CreateProdict(pc.db, nome, autor, sinopse)
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	senha := r.FormValue("senha")
	cliente := repositories.LoginCliente(pc.db, email, senha)
	if cliente != nil {
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/Login.html")
	}
}
