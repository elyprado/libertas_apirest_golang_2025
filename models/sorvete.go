package models

type Sorvete struct {
	ID         int     `json:"id"`
	Sabor      string  `json:"sabor"`
	Preco      float32 `json:"preco"`
	Tipo       string  `json:"tipo"`
	Disponivel bool    `json:"disponivel"`
	Descricao  string  `json:"descricao"`
}
