package models

type Vendedor struct {
	Idvendedor int     `json:"idvendedor"`
	Nome      *string `json:"nome"`
	Email	 *string `json:"email"`
	Cpf     *string `json:"cpf"`
	Telefone     *string `json:"telefone"`
	Estado  *string `json:"estado"`
}
