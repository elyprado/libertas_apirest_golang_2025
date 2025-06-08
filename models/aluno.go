package models

type Aluno struct {
	Nome      string `json:"nome"`
	Idade     int    `json:"idade"`
	Curso     string `json:"curso"`
	Matricula string `json:"matricula"`
	Email     string `json:"email"`
}
