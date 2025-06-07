package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetImoveis(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, erro := db.Query("SELECT id, endereco, cep, valor, contato, STATUS FROM imoveis")
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var imoveis []models.Imovel
	for rows.Next() {
		var i models.Imovel
		if err := rows.Scan(&i.ID, &i.Endereco, &i.CEP, &i.Valor, &i.Contato, &i.STATUS); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		imoveis = append(imoveis, i)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(imoveis)
}

func GetImovelById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := mux.Vars(r)["id"]
	var i models.Imovel
	err := db.QueryRow("SELECT id, endereco, cep, valor, contato, STATUS FROM imoveis WHERE id = ?", id).
		Scan(&i.ID, &i.Endereco, &i.CEP, &i.Valor, &i.Contato, &i.STATUS)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
}

func CreateImovel(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var i models.Imovel
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO imoveis (endereco, cep, valor, contato, STATUS) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, i.Endereco, i.CEP, i.Valor, i.Contato, i.STATUS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Imóvel cadastrado com sucesso"})
}

func UpdateImovel(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := mux.Vars(r)["id"]
	var i models.Imovel
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE imoveis SET endereco = ?, cep = ?, valor = ?, contato = ?, STATUS = ? WHERE id = ?"
	_, err := db.Exec(query, i.Endereco, i.CEP, i.Valor, i.Contato, i.STATUS, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Imóvel atualizado com sucesso"})
}

func DeleteImovel(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := mux.Vars(r)["id"]
	_, err := db.Exec("DELETE FROM imoveis WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Imóvel excluído com sucesso"})
}
