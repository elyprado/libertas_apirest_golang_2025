package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetVeiculos(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	row, erro := db.Query("SELECT idveiculo, modelo, marca, ano, cor, preco FROM veiculo")
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var veiculos []models.Veiculo
	for row.Next() {
		var veiculo models.Veiculo
		erro := row.Scan(&veiculo.Idveiculo, &veiculo.Modelo,
			&veiculo.Marca, &veiculo.Ano, &veiculo.Cor, &veiculo.Preco)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		veiculos = append(veiculos, veiculo)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculos)
}

func GetVeiculoById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	params := mux.Vars(r)
	id := params["id"]

	row, erro := db.Query("SELECT idveiculo, modelo, marca, ano, cor, preco FROM veiculo WHERE idveiculo = ?", id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var veiculo models.Veiculo
	if row.Next() {
		erro := row.Scan(&veiculo.Idveiculo, &veiculo.Modelo,
			&veiculo.Marca, &veiculo.Ano, &veiculo.Cor, &veiculo.Preco)

		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculo)
}

func CreateVeiculo(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var veiculo models.Veiculo
	erro = json.NewDecoder(r.Body).Decode(&veiculo)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	query := "INSERT INTO veiculo (modelo, marca, ano, cor, preco) VALUES (?, ?, ?, ?, ?)"

	_, erro = db.Exec(query, veiculo.Modelo, veiculo.Marca,
		veiculo.Ano, veiculo.Cor, veiculo.Preco)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func UpdateVeiculo(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var veiculo models.Veiculo
	erro = json.NewDecoder(r.Body).Decode(&veiculo)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	params := mux.Vars(r)
	id := params["id"]

	query := "UPDATE veiculo SET modelo=?, marca=?, ano=?, cor=?, preco=? WHERE idveiculo=?"

	_, erro = db.Exec(query, veiculo.Modelo, veiculo.Marca,
		veiculo.Ano, veiculo.Cor, veiculo.Preco, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}

func DeleteVeiculo(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM veiculo WHERE idveiculo=?"

	_, erro = db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "sucesso"})
}
