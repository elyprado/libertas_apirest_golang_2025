package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFilmes(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	row, erro := db.Query("SELECT idfilme, nome, classificacao, genero, ano, autor FROM filmes")
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var filmes []models.Filme
	for row.Next() {
		var filme models.Filme
		erro := row.Scan(&filme.Idfilme, &filme.Nome,
			&filme.Classificacao, &filme.Genero, &filme.Ano, &filme.Autor)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		filmes = append(filmes, filme)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filmes)
}

func GetfilmeById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	params := mux.Vars(r)
	id := params["id"]

	row, erro := db.Query("SELECT idfilme, nome, classificacao, genero, ano, autor FROM filmes WHERE idfilme = ?", id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var filme models.Filme
	if row.Next() {
		erro := row.Scan(&filme.Idfilme, &filme.Nome,
			&filme.Classificacao, &filme.Genero, &filme.Ano, &filme.Autor)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filme)
}

func Createfilme(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var filme models.Filme
	erro = json.NewDecoder(r.Body).Decode(&filme)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	query := "INSERT INTO filmes (nome,classificacao,genero,ano,autor) VALUES (?, ?, ?, ?,?)"

	_, erro = db.Exec(query, filme.Nome, filme.Classificacao,
		filme.Genero, filme.Ano, filme.Autor)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func Updatefilme(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var filme models.Filme
	erro = json.NewDecoder(r.Body).Decode(&filme)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	params := mux.Vars(r)
	id := params["id"]

	query := "UPDATE filmes SET nome=?,classificacao=?,genero=?,ano=?,autor=? WHERE idfilme=?"

	_, erro = db.Exec(query, filme.Nome, filme.Classificacao,
		filme.Genero, filme.Ano, filme.Autor, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func Deletefilme(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM filmes WHERE idfilme=?"

	_, erro = db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}