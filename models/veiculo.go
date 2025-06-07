package models

type Veiculo struct {
	Idveiculo int     `json:"idveiculo"`
	Modelo    *string `json:"modelo"`
	Marca	  *string `json:"marca"`
	Ano       *string `json:"ano"`
	Cor  	  *string `json:"cor"`
	Preco  	  *string `json:"preco"`
}
