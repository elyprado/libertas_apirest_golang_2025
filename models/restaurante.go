package models

type Restaurante struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Telefone    string `json:"telefone"`
	Endereco    string `json:"endereco"`
	TipoCozinha string `json:"tipoCozinha"`
}
