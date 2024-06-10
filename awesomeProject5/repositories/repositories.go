package repositories

import (
	"awesomeProject5/models"
	"database/sql"
)

func BuscarProdutos(db *sql.DB) ([]models.Livro, error) {
	produtos := []models.Livro{}
	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		return nil, err
	}
	defer selectDeTodosOsProdutos.Close()

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome_livro, autor_livro, sinopse_livro string

		err = selectDeTodosOsProdutos.Scan(&id, &nome_livro, &autor_livro, &sinopse_livro)
		if err != nil {
			return nil, err
		}

		produto := models.Livro{
			ID:      id,
			Nome:    nome_livro,
			Autor:   autor_livro,
			Sinopse: sinopse_livro,
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}

func CreateProdict(db *sql.DB, nome, autor, sinopse string) error {
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, autor, sinopse) values($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = insereDadosNoBanco.Exec(nome, autor, sinopse)
	return err
}

func DeleteProduct(db *sql.DB, id string) error {
	delete, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		return err
	}
	_, err = delete.Exec(id)
	return err
}

func EditProduct(db *sql.DB, id string) (models.Livro, error) {
	productDB, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		return models.Livro{}, err
	}
	defer productDB.Close()

	productUpdate := models.Livro{}

	for productDB.Next() {
		var id int
		var nome, autor, sinopse string

		err = productDB.Scan(&id, &nome, &autor, &sinopse)
		if err != nil {
			return models.Livro{}, err
		}

		productUpdate.ID = id
		productUpdate.Nome = nome
		productUpdate.Autor = autor
		productUpdate.Sinopse = sinopse
	}

	return productUpdate, nil
}

func UpdateProduct(db *sql.DB, id int, nome, autor, sinopse string) error {
	updateProduct, err := db.Prepare("update produtos set nome=$1, autor=$2, sinopse=$3 where id=$4")
	if err != nil {
		return err
	}
	_, err = updateProduct.Exec(nome, autor, sinopse, id)
	return err
}

func LoginCliente(db *sql.DB, email, senha string) (*models.Cliente, error) {
	cliente := &models.Cliente{}
	err := db.QueryRow("SELECT * FROM clientes WHERE email=$1 AND senha=$2", email, senha).Scan(&cliente.ID, &cliente.Email, &cliente.NomeCliente, &cliente.Senha)
	if err != nil {
		return nil, err
	}
	return cliente, nil
}
