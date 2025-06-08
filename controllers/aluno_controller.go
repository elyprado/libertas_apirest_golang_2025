package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAlunos(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	row, erro := db.Query("SELECT nome, idade, curso, matricula, email FROM aluno")
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var alunos []models.Aluno
	for row.Next() {
		var aluno models.Aluno
		erro := row.Scan(&aluno.Nome, &aluno.Idade, &aluno.Curso, &aluno.Matricula, &aluno.Email)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		alunos = append(alunos, aluno)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

func GetAlunoById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	row, erro := db.Query("SELECT nome, idade, curso, matricula, email FROM aluno WHERE matricula = ?", id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var aluno models.Aluno
	if row.Next() {
		erro := row.Scan(&aluno.Nome, &aluno.Idade, &aluno.Curso, &aluno.Matricula, &aluno.Email)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

func CreateAluno(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var aluno models.Aluno
	erro = json.NewDecoder(r.Body).Decode(&aluno)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	query := "INSERT INTO aluno (nome, idade, curso, matricula, email) VALUES (?, ?, ?, ?, ?)"
	_, erro = db.Exec(query, aluno.Nome, aluno.Idade, aluno.Curso, aluno.Matricula, aluno.Email)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "sucesso"})
}

func UpdateAluno(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var aluno models.Aluno
	erro = json.NewDecoder(r.Body).Decode(&aluno)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	params := mux.Vars(r)
	id := params["id"]

	query := "UPDATE aluno SET nome=?, idade=?, curso=?, email=? WHERE matricula=?"
	_, erro = db.Exec(query, aluno.Nome, aluno.Idade, aluno.Curso, aluno.Email, id)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "sucesso"})
}

func DeleteAluno(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM aluno WHERE matricula=?"
	_, erro = db.Exec(query, id)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "sucesso"})
}
