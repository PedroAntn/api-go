package repositories

import (
	"database/sql"
	"log"
)

type Produto struct {
	Id      int
	Nome    string
	Autor   string
	Sinopse string
}

type Cliente struct {
	Id    int
	Email string
}

func BuscarProdutos(db *sql.DB) []Produto {
	rows, err := db.Query("SELECT id, nome, autor, sinopse FROM produtos")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var p Produto
		if err := rows.Scan(&p.Id, &p.Nome, &p.Autor, &p.Sinopse); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		produtos = append(produtos, p)
	}
	return produtos
}

func CreateProduct(db *sql.DB, nome, autor, sinopse string) {
	_, err := db.Exec("INSERT INTO produtos (nome, autor, sinopse) VALUES ($1, $2, $3)", nome, autor, sinopse)
	if err != nil {
		log.Printf("Error inserting product: %v", err)
	}
}

func DeleteProduct(db *sql.DB, id string) {
	_, err := db.Exec("DELETE FROM produtos WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting product: %v", err)
	}
}

func EditProduct(db *sql.DB, id string) Produto {
	row := db.QueryRow("SELECT id, nome, autor, sinopse FROM produtos WHERE id = $1", id)
	var p Produto
	err := row.Scan(&p.Id, &p.Nome, &p.Autor, &p.Sinopse)
	if err != nil {
		log.Printf("Error scanning row: %v", err)
	}
	return p
}

func UpdateProduct(db *sql.DB, id int, nome, autor, sinopse string) {
	_, err := db.Exec("UPDATE produtos SET nome = $1, autor = $2, sinopse = $3 WHERE id = $4", nome, autor, sinopse, id)
	if err != nil {
		log.Printf("Error updating product: %v", err)
	}
}

func LoginCliente(db *sql.DB, email, password string) *Cliente {
	row := db.QueryRow("SELECT id, email FROM clientes WHERE email = $1 AND password = $2", email, password)
	var c Cliente
	err := row.Scan(&c.Id, &c.Email)
	if err != nil {
		log.Printf("Error logging in: %v", err)
		return nil
	}
	return &c
}
