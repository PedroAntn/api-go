package models

type Livro struct {
	Id      int
	Nome    string
	Autor   string
	Sinopse string
}

type Cliente struct {
	Id          int
	Email       string
	NomeCliente string
	Senha       string
}

type Funcionario struct {
	Id              int
	Email           string
	NomeFuncionario string
	Senha           string
}
