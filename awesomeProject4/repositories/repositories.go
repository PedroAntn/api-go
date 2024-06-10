package repositories

import (
	"awesomeProject4/models"
	"database/sql"
)

func BuscarProdutos(db *sql.DB) []models.Livro {
	produtos := []models.Livro{}
	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	defer selectDeTodosOsProdutos.Close()

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome_livro, autor_livro, sinopse_livro string

		err = selectDeTodosOsProdutos.Scan(&id, &nome_livro, &autor_livro, &sinopse_livro)
		if err != nil {
			panic(err.Error())
		}

		produto := models.Livro{
			Id:      id,
			Nome:    nome_livro,
			Autor:   autor_livro,
			Sinopse: sinopse_livro,
		}

		produtos = append(produtos, produto)
	}

	return produtos
}

func CreateProdict(db *sql.DB, nome, autor, sinopse string) {
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, autor, sinopse) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, autor, sinopse)
}

func DeleteProduct(db *sql.DB, id string) {
	delete, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
}

func EditProduct(db *sql.DB, id string) models.Livro {
	productDB, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	defer productDB.Close()

	productUpdate := models.Livro{}

	for productDB.Next() {
		var id int
		var nome, autor, sinopse string

		err = productDB.Scan(&id, &nome, &autor, &sinopse)
		if err != nil {
			panic(err.Error())
		}

		productUpdate.Id = id
		productUpdate.Nome = nome
		productUpdate.Autor = autor
		productUpdate.Sinopse = sinopse
	}

	return productUpdate
}

func UpdateProduct(db *sql.DB, id int, nome, autor, sinopse string) {
	updateProduct, err := db.Prepare("update produtos set nome=$1, autor=$2, sinopse=$3 where id=$4")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(nome, autor, sinopse, id)
}
func LoginCliente(db *sql.DB, email, senha string) *models.Cliente {
	cliente := &models.Cliente{}
	err := db.QueryRow("SELECT * FROM clientes WHERE email=$1 AND senha=$2", email, senha).Scan(&cliente.Id, &cliente.Email, &cliente.NomeCliente, &cliente.Senha)
	if err != nil {
		return nil
	}
	return cliente
}
