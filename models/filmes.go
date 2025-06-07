package models

type Filme struct {
    Idfilme           int    `json:"id"`
    Nome         string `json:"nome"`
    Classificacao string `json:"classificacao"`
    Genero       string `json:"genero"`
    Ano          int    `json:"ano"`
    Autor        string `json:"autor"`
}
