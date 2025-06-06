package models

type Produto struct {
    ID           int     `json:"id"`
    Descricao    string  `json:"descricao"`
    Fornecedor   string  `json:"fornecedor"`
    Estoque      int     `json:"estoque"`     
    Valor        float64 `json:"valor"`       
    Detalhes     string  `json:"detalhes"`
}

