package models

type Medicamento struct {
	Idmedicamento int     `json:"idmedicamento"`
	Nome          *string `json:"nome"`
	Quantidade    *string `json:"quantidade"`
	Tipo          *string `json:"tipo"`
	Fabricante    *string `json:"fabricante"`
}
