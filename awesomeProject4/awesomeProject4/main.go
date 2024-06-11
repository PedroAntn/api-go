package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

type ProductController struct {
	db *sql.DB
}

func (pc ProductController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	produtos, err := pc.BuscarProdutos()
	if err != nil {
		log.Println("Error querying database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "Index.html", produtos)
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (pc ProductController) BuscarProdutos() ([]Produto, error) {
	rows, err := pc.db.Query("SELECT id, nome_livro, autor_livro, sinopse_livro FROM livros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var p Produto
		if err := rows.Scan(&p.ID, &p.NomeLivro, &p.AutorLivro, &p.SinopseLivro); err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return produtos, nil
}

type Produto struct {
	ID           int
	NomeLivro    string
	AutorLivro   string
	SinopseLivro string
}

func templateExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func init() {
	templatePath := "templates/Index.html"
	if !templateExists(templatePath) {
		log.Fatalf("Template file does not exist: %s", templatePath)
	}
}

func main() {
	connStr := "user=username dbname=mydb sslmode=disable" // Modifique para suas credenciais
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	router := httprouter.New()
	pc := &ProductController{db: db}

	router.GET("/", pc.Index)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
