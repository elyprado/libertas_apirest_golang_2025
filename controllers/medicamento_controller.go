package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMedicamentos(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	nome := r.URL.Query().Get("nome")

	query := "SELECT idmedicamento, nome, quantidade, tipo, fabricante FROM medicamento"

	var rows *sql.Rows

	if nome != "" {
		query += " WHERE nome LIKE ?"
		nome = "%" + nome + "%"
		rows, erro = db.Query(query, nome)
	} else {
		rows, erro = db.Query(query)
	}

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var medicamentos []models.Medicamento
	for rows.Next() {
		var medicamento models.Medicamento
		erro := rows.Scan(&medicamento.Idmedicamento, &medicamento.Nome,
			&medicamento.Quantidade, &medicamento.Tipo, &medicamento.Fabricante)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		medicamentos = append(medicamentos, medicamento)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicamentos)
}

func GetMedicamentoById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //executa no fim do método

	params := mux.Vars(r)
	id := params["id"]

	row, erro := db.Query("SELECT idmedicamento, nome, quantidade, tipo, fabricante FROM medicamento WHERE idmedicamento = ?", id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var medicamento models.Medicamento
	if row.Next() {
		erro := row.Scan(&medicamento.Idmedicamento, &medicamento.Nome,
			&medicamento.Quantidade, &medicamento.Tipo, &medicamento.Fabricante)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicamento)
}

func CreateMedicamento(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var medicamento models.Medicamento
	erro = json.NewDecoder(r.Body).Decode(&medicamento)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	query := "INSERT INTO medicamento (nome,quantidade,tipo,fabricante) VALUES (?, ?, ?, ?)"

	_, erro = db.Exec(query, medicamento.Nome, medicamento.Quantidade,
		medicamento.Tipo, medicamento.Fabricante)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "Medicamento cadastrado com sucesso!", "success": "true"})
}

func UpdateMedicamento(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	var medicamento models.Medicamento
	erro = json.NewDecoder(r.Body).Decode(&medicamento)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	params := mux.Vars(r)
	id := params["id"]

	query := "UPDATE medicamento SET nome=?,quantidade=?,tipo=?,fabricante=? WHERE idmedicamento=?"

	_, erro = db.Exec(query, medicamento.Nome, medicamento.Quantidade,
		medicamento.Tipo, medicamento.Fabricante, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "Medicamento atualizado com sucesso!", "success": "true"})
}

func DeleteMedicamento(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM medicamento WHERE idmedicamento=?"

	_, erro = db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{"message": "Medicamento excluído com sucesso!", "success": "true"})
}
