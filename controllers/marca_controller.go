package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMarcas(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	row, erro := db.Query("SELECT idmarca, nome, nicho, cnpj, site FROM marcas_5P")
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var marcas []models.Marca
	for row.Next() {
		var marca models.Marca
		erro := row.Scan(&marca.Idmarca, &marca.Nome,
			&marca.Nicho, &marca.Cnpj, &marca.Site)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		marcas = append(marcas, marca)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marcas)
}

func GetMarcaById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	params := mux.Vars(r)
	id := params["id"]

	row, erro := db.Query("SELECT idmarca, nome, nicho, cnpj, site FROM marcas_5P WHERE idmarca = ?", id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var marca models.Marca
	if row.Next() {
		erro := row.Scan(&marca.Idmarca, &marca.Nome,
			&marca.Nicho, &marca.Cnpj, &marca.Site)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marca)
}

func CreateMarca(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var marca models.Marca
	erro = json.NewDecoder(r.Body).Decode(&marca)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	query := "INSERT INTO marcas_5P (nome, nicho, cnpj, site) VALUES (?, ?, ?, ?)"

	_, erro = db.Exec(query, marca.Nome, marca.Nicho,
		marca.Cnpj, marca.Site)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func UpdateMarca(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var marca models.Marca
	erro = json.NewDecoder(r.Body).Decode(&marca)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	params := mux.Vars(r)
	id := params["id"]

	query := "UPDATE marcas_5P SET nome=?,nicho=?,cnpj=?,site=? WHERE idmarca=?"

	_, erro = db.Exec(query, marca.Nome, marca.Nicho,
		marca.Cnpj, marca.Site, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func DeleteMarca(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM marcas_5P WHERE idmarca=?"

	_, erro = db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}
