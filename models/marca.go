package models

type Marca struct {
	Idmarca int     `json:"idmarca"`
	Nome  	*string `json:"nome"`
	Nicho   *string `json:"nicho"`
	Cnpj    *string `json:"cnpj"`
	Site  	*string `json:"site"`
}
