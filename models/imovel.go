package models

type Imovel struct {
	ID      int     `json:"id"`
	Endereco string `json:"endereco"`
	CEP     string `json:"cep"`
	Valor   float64 `json:"valor"`
	Contato string  `json:"contato"`
	STATUS  string  `json:"STATUS"`
}
