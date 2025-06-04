package models

type Item struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Marca      string  `json:"marca"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
	Observacao string  `json:"observacao"`
}

