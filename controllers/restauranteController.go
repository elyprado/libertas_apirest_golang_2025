package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"restaurante-go/config"
	"restaurante-go/models"
	"strconv"

	"github.com/gorilla/mux"
)

func SetCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func GetRestaurantes(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em GetRestaurantes.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT idrestaurante, nome, telefone, endereco, tipo_cozinha FROM restaurante")
	if err != nil {
		log.Printf("Erro ao executar a consulta GetRestaurantes: %v", err)
		http.Error(w, "Erro ao carregar restaurantes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var restaurantes []models.Restaurante

	for rows.Next() {
		var restaurante models.Restaurante
		err := rows.Scan(&restaurante.ID, &restaurante.Nome, &restaurante.Telefone, &restaurante.Endereco, &restaurante.TipoCozinha)
		if err != nil {
			log.Printf("Erro ao ler resultados GetRestaurantes: %v", err)
			http.Error(w, "Erro ao processar dados de restaurantes", http.StatusInternalServerError)
			return
		}
		restaurantes = append(restaurantes, restaurante)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração dos resultados GetRestaurantes: %v", err)
		http.Error(w, "Erro ao finalizar processamento de restaurantes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(restaurantes)
}

func GetRestauranteByID(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em GetRestauranteByID.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do restaurante não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido em GetRestauranteByID '%s': %v", idStr, err)
		http.Error(w, "ID do restaurante inválido", http.StatusBadRequest)
		return
	}

	query := "SELECT idrestaurante, nome, telefone, endereco, tipo_cozinha FROM restaurante WHERE idrestaurante = ?"
	row := db.QueryRow(query, id)

	var restaurante models.Restaurante
	err = row.Scan(&restaurante.ID, &restaurante.Nome, &restaurante.Telefone, &restaurante.Endereco, &restaurante.TipoCozinha)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Restaurante não encontrado", http.StatusNotFound)
			return
		}
		log.Printf("Erro ao ler resultado GetRestauranteByID para ID %d: %v", id, err)
		http.Error(w, "Erro ao buscar restaurante", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(restaurante)
}

func CreateRestaurante(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em CreateRestaurante.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	var restaurante models.Restaurante
	err := json.NewDecoder(r.Body).Decode(&restaurante)
	if err != nil {
		log.Printf("Erro ao decodificar JSON em CreateRestaurante: %v", err)
		http.Error(w, "Dados inválidos do restaurante", http.StatusBadRequest)
		return
	}

	if restaurante.Nome == "" || restaurante.Telefone == "" || restaurante.Endereco == "" || restaurante.TipoCozinha == "" {
		http.Error(w, "Nome, Telefone, Endereço e Tipo de Cozinha são obrigatórios", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO restaurante (nome, telefone, endereco, tipo_cozinha) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, restaurante.Nome, restaurante.Telefone, restaurante.Endereco, restaurante.TipoCozinha)
	if err != nil {
		log.Printf("Erro ao inserir restaurante: %v", err)
		http.Error(w, "Erro ao criar o restaurante", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erro ao obter LastInsertId: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Restaurante criado com sucesso", "id": id})
}

func UpdateRestaurante(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em UpdateRestaurante.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do restaurante não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido para atualização '%s': %v", idStr, err)
		http.Error(w, "ID do restaurante inválido", http.StatusBadRequest)
		return
	}

	var restaurante models.Restaurante
	err = json.NewDecoder(r.Body).Decode(&restaurante)
	if err != nil {
		log.Printf("Erro ao decodificar JSON em UpdateRestaurante: %v", err)
		http.Error(w, "Dados inválidos do restaurante", http.StatusBadRequest)
		return
	}

	if restaurante.Nome == "" || restaurante.Telefone == "" || restaurante.Endereco == "" || restaurante.TipoCozinha == "" {
		http.Error(w, "Nome, Telefone, Endereço e Tipo de Cozinha são obrigatórios", http.StatusBadRequest)
		return
	}

	query := "UPDATE restaurante SET nome = ?, telefone = ?, endereco = ?, tipo_cozinha = ? WHERE idrestaurante = ?"
	result, err := db.Exec(query, restaurante.Nome, restaurante.Telefone, restaurante.Endereco, restaurante.TipoCozinha, id)
	if err != nil {
		log.Printf("Erro ao atualizar restaurante ID %d: %v", id, err)
		http.Error(w, "Erro ao atualizar o restaurante", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter RowsAffected em UpdateRestaurante: %v", err)
	} else if rowsAffected == 0 {
		http.Error(w, "Restaurante não encontrado ou nenhum dado alterado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Restaurante atualizado com sucesso"})
}

func DeleteRestaurante(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em DeleteRestaurante.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do restaurante não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido para deleção '%s': %v", idStr, err)
		http.Error(w, "ID do restaurante inválido", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM restaurante WHERE idrestaurante = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao excluir restaurante ID %d: %v", id, err)
		http.Error(w, "Erro ao excluir o restaurante", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter RowsAffected em DeleteRestaurante: %v", err)
	} else if rowsAffected == 0 {
		http.Error(w, "Restaurante não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Restaurante excluído com sucesso"})
}

func CorsPreflight(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)
	w.WriteHeader(http.StatusOK)
}
