package models

type Livro struct {
	ID      int    `json:"id"`
	Nome    string `json:"nome"`
	Autor   string `json:"autor"`
	Sinopse string `json:"sinopse"`
}

type Cliente struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	NomeCliente string `json:"nome_cliente"`
	Senha       string `json:"senha"`
}

type Funcionario struct {
	ID              int    `json:"id"`
	Email           string `json:"email"`
	NomeFuncionario string `json:"nome_funcionario"`
	Senha           string `json:"senha"`
}
