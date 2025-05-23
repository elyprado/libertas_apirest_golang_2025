package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetVendedores(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	row, erro := db.Query("SELECT idvendedor, nome, email, cpf, telefone, estado FROM vendedores_matheus")
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var vendedores []models.Vendedor
	for row.Next() {
		var vendedor models.Vendedor
		erro := row.Scan(&vendedor.Idvendedor, &vendedor.Nome,
			&vendedor.Email, &vendedor.Cpf, &vendedor.Telefone, &vendedor.Estado)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		vendedores = append(vendedores, vendedor)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedores)
}

func GetVendedorById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	params := mux.Vars(r)
	id := params["id"]

	row, erro := db.Query("SELECT idvendedor, nome, email, cpf, telefone, estado FROM vendedores_matheus WHERE idvendedor = ?", id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var vendedor models.Vendedor
	if row.Next() {
		erro := row.Scan(&vendedor.Idvendedor, &vendedor.Nome,
			&vendedor.Email, &vendedor.Cpf, &vendedor.Telefone, &vendedor.Estado)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedor)
}

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var vendedor models.Vendedor
	erro = json.NewDecoder(r.Body).Decode(&vendedor)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	query := "INSERT INTO vendedores_matheus (nome,email,cpf,telefone, estado) VALUES (?, ?, ?, ?, ?)"

	_, erro = db.Exec(query, vendedor.Nome, vendedor.Email,
		vendedor.Cpf, vendedor.Telefone, vendedor.Estado)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func UpdateVendedor(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var vendedor models.Vendedor
	erro = json.NewDecoder(r.Body).Decode(&vendedor)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	params := mux.Vars(r)
	id := params["id"]

	query := "UPDATE vendedores_matheus SET nome=?,email=?,cpf=?,telefone=?,estado=? WHERE idvendedor=?"

	_, erro = db.Exec(query, vendedor.Nome, vendedor.Email,
		vendedor.Cpf, vendedor.Telefone, vendedor.Estado, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func DeleteVendedor(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM vendedores_matheus WHERE idvendedor=?"

	_, erro = db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}
