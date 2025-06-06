package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProdutos(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, descricao, fornecedor, estoque, valor, detalhes FROM produtos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []models.Produto
	for rows.Next() {
		var produto models.Produto
		err := rows.Scan(&produto.ID, &produto.Descricao, &produto.Fornecedor, &produto.Estoque, &produto.Valor, &produto.Detalhes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, produto)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func GetProdutoById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	row := db.QueryRow("SELECT id, descricao, fornecedor, estoque, valor, detalhes FROM produtos WHERE id = ?", id)

	var produto models.Produto
	err = row.Scan(&produto.ID, &produto.Descricao, &produto.Fornecedor, &produto.Estoque, &produto.Valor, &produto.Detalhes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var produto models.Produto
	err = json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO produtos (descricao, fornecedor, estoque, valor, detalhes) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(query, produto.Descricao, produto.Fornecedor, produto.Estoque, produto.Valor, produto.Detalhes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "produto inserido com sucesso"})
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var produto models.Produto
	err = json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE produtos SET descricao = ?, fornecedor = ?, estoque = ?, valor = ?, detalhes = ? WHERE id = ?"
	_, err = db.Exec(query, produto.Descricao, produto.Fornecedor, produto.Estoque, produto.Valor, produto.Detalhes, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "produto atualizado com sucesso"})
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM produtos WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "produto removido com sucesso"})
}
