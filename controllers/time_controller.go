package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := r.URL.Query()

	var filtros []string
	var args []interface{}

	if id := params.Get("idtime"); id != "" {
		filtros = append(filtros, "idtime = ?")
		args = append(args, id)
	}
	if nome := params.Get("nome"); nome != "" {
		filtros = append(filtros, "nome LIKE ?")
		args = append(args, "%"+nome+"%")
	}
	if cidade := params.Get("cidade"); cidade != "" {
		filtros = append(filtros, "cidade LIKE ?")
		args = append(args, "%"+cidade+"%")
	}
	if estado := params.Get("estado"); estado != "" {
		filtros = append(filtros, "estado LIKE ?")
		args = append(args, "%"+estado+"%")
	}
	if fundacao := params.Get("fundacao"); fundacao != "" {
		filtros = append(filtros, "fundacao = ?")
		args = append(args, fundacao)
	}
	if estadio := params.Get("estadio"); estadio != "" {
		filtros = append(filtros, "estadio LIKE ?")
		args = append(args, "%"+estadio+"%")
	}

	query := "SELECT idtime, nome, cidade, estado, fundacao, estadio FROM time_futebol"
	if len(filtros) > 0 {
		query += " WHERE " + strings.Join(filtros, " OR ")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var times []models.Time
	for rows.Next() {
		var time models.Time
		err := rows.Scan(&time.Idtime, &time.Nome, &time.Cidade, &time.Estado, &time.Fundacao, &time.Estadio)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		times = append(times, time)
	}

	if len(times) == 0 {
		http.Error(w, "Nenhum time encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(times)
}

func GetTimeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	params := mux.Vars(r)
	id := params["idtime"]
	if id == "" {
		http.Error(w, "ID do time não fornecido", http.StatusBadRequest)
		return
	}
	query := "SELECT idtime, nome, cidade, estado, fundacao, estadio FROM time_futebol WHERE idtime = ?"
	row := db.QueryRow(query, id)
	var time models.Time
	err = row.Scan(&time.Idtime, &time.Nome, &time.Cidade, &time.Estado, &time.Fundacao, &time.Estadio)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Time não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(time)
}

func CreateTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var time models.Time
	err = json.NewDecoder(r.Body).Decode(&time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO time_futebol (nome, cidade, estado, fundacao, estadio) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(query, time.Nome, time.Cidade, time.Estado, time.Fundacao, time.Estadio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Time criado com sucesso"})
}

func UpdateTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var time models.Time
	err = json.NewDecoder(r.Body).Decode(&time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id := params["idtime"]

	query := "UPDATE time_futebol SET nome = ?, cidade = ?, estado = ?, fundacao = ?, estadio = ? WHERE idtime = ?"
	_, err = db.Exec(query, time.Nome, time.Cidade, time.Estado, time.Fundacao, time.Estadio, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Time atualizado com sucesso"})
}

func DeleteTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["idtime"]
	if id == "" {
		http.Error(w, "ID do time não fornecido", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM time_futebol WHERE idtime = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Time excluído com sucesso"})
}
