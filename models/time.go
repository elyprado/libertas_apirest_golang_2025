package models

type Time struct {
	Idtime   *int    `json:"idtime"`
	Nome     *string `json:"nome"`
	Cidade   *string `json:"cidade"`
	Estado   *string `json:"estado"`
	Fundacao *int    `json:"fundacao"`
	Estadio  *string `json:"estadio"`
}
