package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sorvetes-go/config"
	"sorvetes-go/models"
	"strconv"

	"github.com/gorilla/mux"
)

func SetCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func GetSorvetes(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)
	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em GetSorvetes.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, sabor, preco, tipo, disponivel, descricao FROM sorvete")
	if err != nil {
		log.Printf("Erro ao executar a consulta GetSorvetes: %v", err)
		http.Error(w, "Erro ao carregar sorvetes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sorvetes []models.Sorvete

	for rows.Next() {
		var sorvete models.Sorvete
		err := rows.Scan(&sorvete.ID, &sorvete.Sabor, &sorvete.Preco, &sorvete.Tipo, &sorvete.Disponivel, &sorvete.Descricao)
		if err != nil {
			log.Printf("Erro ao ler resultados GetSorvetes: %v", err)
			http.Error(w, "Erro ao processar dados de sorvetes", http.StatusInternalServerError)
			return
		}
		sorvetes = append(sorvetes, sorvete)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração dos resultados GetSorvetes: %v", err)
		http.Error(w, "Erro ao finalizar processamento de sorvetes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sorvetes)
}

func GetSorveteByID(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em GetSorveteByID.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do sorvete não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido em GetSorveteByID '%s': %v", idStr, err)
		http.Error(w, "ID do sorvete inválido", http.StatusBadRequest)
		return
	}

	query := "SELECT id, sabor, preco, tipo, disponivel, descricao FROM sorvete WHERE id = ?"
	row := db.QueryRow(query, id)

	var sorvete models.Sorvete
	err = row.Scan(&sorvete.ID, &sorvete.Sabor, &sorvete.Preco, &sorvete.Tipo, &sorvete.Disponivel, &sorvete.Descricao)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Sorvete não encontrado", http.StatusNotFound)
			return
		}
		log.Printf("Erro ao ler resultado GetSorveteByID para ID %d: %v", id, err)
		http.Error(w, "Erro ao buscar sorvete", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sorvete)
}

func CreateSorvete(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em CreateSorvete.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	var sorvete models.Sorvete
	err := json.NewDecoder(r.Body).Decode(&sorvete)
	if err != nil {
		log.Printf("Erro ao decodificar JSON em CreateSorvete: %v", err)
		http.Error(w, "Dados inválidos do sorvete", http.StatusBadRequest)
		return
	}

	if sorvete.Sabor == "" || sorvete.Preco <= 0 || sorvete.Tipo == "" {
		http.Error(w, "Sabor, Preço e Tipo são obrigatórios e Preço deve ser maior que zero", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO sorvete (sabor, preco, tipo, disponivel, descricao) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, sorvete.Sabor, sorvete.Preco, sorvete.Tipo, sorvete.Disponivel, sorvete.Descricao)
	if err != nil {
		log.Printf("Erro ao inserir sorvete: %v", err)
		http.Error(w, "Erro ao criar o sorvete", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erro ao obter LastInsertId: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Sorvete criado com sucesso", "id": id})
}

func UpdateSorvete(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em UpdateSorvete.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do sorvete não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido para atualização '%s': %v", idStr, err)
		http.Error(w, "ID do sorvete inválido", http.StatusBadRequest)
		return
	}

	var sorvete models.Sorvete
	err = json.NewDecoder(r.Body).Decode(&sorvete)
	if err != nil {
		log.Printf("Erro ao decodificar JSON em UpdateSorvete: %v", err)
		http.Error(w, "Dados inválidos do sorvete", http.StatusBadRequest)
		return
	}

	if sorvete.Sabor == "" || sorvete.Preco <= 0 || sorvete.Tipo == "" {
		http.Error(w, "Sabor, Preço e Tipo são obrigatórios e Preço deve ser maior que zero", http.StatusBadRequest)
		return
	}

	query := "UPDATE sorvete SET sabor = ?, preco = ?, tipo = ?, disponivel = ?, descricao = ? WHERE id = ?"
	result, err := db.Exec(query, sorvete.Sabor, sorvete.Preco, sorvete.Tipo, sorvete.Disponivel, sorvete.Descricao, id)
	if err != nil {
		log.Printf("Erro ao atualizar sorvete ID %d: %v", id, err)
		http.Error(w, "Erro ao atualizar o sorvete", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter RowsAffected em UpdateSorvete: %v", err)
	} else if rowsAffected == 0 {
		http.Error(w, "Sorvete não encontrado ou nenhum dado alterado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sorvete atualizado com sucesso"})
}

func DeleteSorvete(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)

	db := config.GetDB()
	if db == nil {
		log.Println("Erro: Conexão com o banco de dados não inicializada em DeleteSorvete.")
		http.Error(w, "Erro interno do servidor: conexão DB não disponível.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	if idStr == "" {
		http.Error(w, "ID do sorvete não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ID inválido para deleção '%s': %v", idStr, err)
		http.Error(w, "ID do sorvete inválido", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM sorvete WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao excluir sorvete ID %d: %v", id, err)
		http.Error(w, "Erro ao excluir o sorvete", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter RowsAffected em DeleteSorvete: %v", err)
	} else if rowsAffected == 0 {
		http.Error(w, "Sorvete não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sorvete excluído com sucesso"})
}

func CorsPreflight(w http.ResponseWriter, r *http.Request) {
	SetCommonHeaders(w)
	w.WriteHeader(http.StatusOK)
}
