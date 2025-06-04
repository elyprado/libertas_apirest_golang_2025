package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

func GetItens(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var dbName string
	err := db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Banco selecionado na conexão:", dbName)  // imprime no console

	rows, erro := db.Query("SELECT id, nome, marca, quantidade, preco, observacao FROM lista_compras_db")
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var itens []models.Item
	for rows.Next() {
		var item models.Item
		erro := rows.Scan(&item.ID, &item.Nome, &item.Marca, &item.Quantidade, &item.Preco, &item.Observacao)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		itens = append(itens, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itens)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	row := db.QueryRow("SELECT id, nome, marca, quantidade, preco, observacao FROM lista_compras_db WHERE id = ?", id)

	var item models.Item
	erro = row.Scan(&item.ID, &item.Nome, &item.Marca, &item.Quantidade, &item.Preco, &item.Observacao)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var item models.Item
	erro = json.NewDecoder(r.Body).Decode(&item)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO lista_compras_db (nome, marca, quantidade, preco, observacao) VALUES (?, ?, ?, ?, ?)"
	_, erro = db.Exec(query, item.Nome, item.Marca, item.Quantidade, item.Preco, item.Observacao)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "item inserido com sucesso"})
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var item models.Item
	erro = json.NewDecoder(r.Body).Decode(&item)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE lista_compras_db SET nome = ?, marca = ?, quantidade = ?, preco = ?, observacao = ? WHERE id = ?"
	_, erro = db.Exec(query, item.Nome, item.Marca, item.Quantidade, item.Preco, item.Observacao, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "item atualizado com sucesso"})
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM lista_compras_db WHERE id = ?"
	_, erro = db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "item removido com sucesso"})
}
