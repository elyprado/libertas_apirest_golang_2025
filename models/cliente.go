package models

type Cliente struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}
