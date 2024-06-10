package controllers

import (
	"awesomeProject5/repositories"
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
	produtos, err := repositories.BuscarProdutos(pc.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := template.New("Index").ParseFiles("templates/Index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, produtos)
}

func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	err := repositories.CreateProdict(pc.db, nome, autor, sinopse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := repositories.DeleteProduct(pc.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) EditProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product, err := repositories.EditProduct(pc.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := template.New("Edit").ParseFiles("templates/Edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, product)
}

func (pc *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	err = repositories.UpdateProduct(pc.db, id, nome, autor, sinopse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) Insert(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	autor := r.FormValue("autor")
	sinopse := r.FormValue("sinopse")
	err := repositories.CreateProdict(pc.db, nome, autor, sinopse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 301)
}

func (pc *ProductController) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	senha := r.FormValue("senha")
	cliente, err := repositories.LoginCliente(pc.db, email, senha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cliente == nil {
		http.ServeFile(w, r, "templates/Login.html")
		return
	}
	http.Redirect(w, r, "/", 301)
}
