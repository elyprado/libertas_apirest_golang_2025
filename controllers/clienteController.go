package controllers

import (
	"clientes-go/config"
	"clientes-go/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SetCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func GetClientes(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em GetClientes.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT idcliente, nome, email, telefone, endereco FROM cliente")
	if err != nil {
		log.Printf("Erro ao executar a consulta GetClientes: %v", err)
		http.Error(w, "Erro ao carregar clientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clientes []models.Cliente

	for rows.Next() {
		var cliente models.Cliente
		err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone, &cliente.Endereco)
		if err != nil {
			log.Printf("Erro ao ler resultados GetClientes: %v", err)
			http.Error(w, "Erro ao processar dados de clientes", http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração dos resultados GetClientes: %v", err)
		http.Error(w, "Erro ao finalizar processamento de clientes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clientes)
}

func GetClienteByID(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em GetClienteByID.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do cliente não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido em GetClienteByID '%s': %v", idStr, err)
		http.Error(w, "ID do cliente inválido", http.StatusBadRequest)
		return
	}

	query := "SELECT idcliente, nome, email, telefone, endereco FROM cliente WHERE idcliente = ?"
	row := db.QueryRow(query, id)

	var cliente models.Cliente
	err = row.Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone, &cliente.Endereco)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Cliente não encontrado", http.StatusNotFound)
			return
		}
		log.Printf("Erro ao ler resultado GetClienteByID para ID %d: %v", id, err)
		http.Error(w, "Erro ao buscar cliente", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cliente)
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em CreateCliente.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		log.Printf("Erro ao decodificar JSON em CreateCliente: %v", err)
		http.Error(w, "Dados inválidos do cliente", http.StatusBadRequest)
		return
	}

	if cliente.Nome == "" || cliente.Email == "" || cliente.Telefone == "" || cliente.Endereco == "" {
		http.Error(w, "Nome, Email, Telefone e Endereço são obrigatórios", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO cliente (nome, email, telefone, endereco) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, cliente.Nome, cliente.Email, cliente.Telefone, cliente.Endereco)
	if err != nil {
		log.Printf("Erro ao inserir cliente: %v", err)
		http.Error(w, "Erro ao criar o cliente", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erro ao obter LastInsertId: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Cliente criado com sucesso", "id": id})
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em UpdateCliente.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do cliente não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido para atualização '%s': %v", idStr, err)
		http.Error(w, "ID do cliente inválido", http.StatusBadRequest)
		return
	}

	var cliente models.Cliente
	err = json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		log.Printf("Erro ao decodificar JSON em UpdateCliente: %v", err)
		http.Error(w, "Dados inválidos do cliente", http.StatusBadRequest)
		return
	}

	if cliente.Nome == "" || cliente.Email == "" || cliente.Telefone == "" || cliente.Endereco == "" {
		http.Error(w, "Nome, Email, Telefone e Endereço são obrigatórios", http.StatusBadRequest)
		return
	}

	query := "UPDATE cliente SET nome = ?, email = ?, telefone = ?, endereco = ? WHERE idcliente = ?"
	result, err := db.Exec(query, cliente.Nome, cliente.Email, cliente.Telefone, cliente.Endereco, id)
	if err != nil {
		log.Printf("Erro ao atualizar cliente ID %d: %v", id, err)
		http.Error(w, "Erro ao atualizar o cliente", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter RowsAffected em UpdateCliente: %v", err)
	} else if rowsAffected == 0 {
		http.Error(w, "Cliente não encontrado ou nenhum dado alterado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente atualizado com sucesso"})
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em DeleteCliente.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do cliente não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido para deleção '%s': %v", idStr, err)
		http.Error(w, "ID do cliente inválido", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM cliente WHERE idcliente = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao excluir cliente ID %d: %v", id, err)
		http.Error(w, "Erro ao excluir o cliente", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter RowsAffected em DeleteCliente: %v", err)
	} else if rowsAffected == 0 {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente excluído com sucesso"})
}

func CorsPreflight(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)
	w.WriteHeader(http.StatusOK)
}
