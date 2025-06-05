package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCursos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, "Erro ao conectar no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS curso (
	  idcurso INT AUTO_INCREMENT PRIMARY KEY,
	  nome VARCHAR(255) NOT NULL,
	  descricao TEXT,
	  cargaHoraria VARCHAR(50),
	  valor DECIMAL(10,2)
	);`
	if _, err := db.Exec(createTable); err != nil {
		http.Error(w, "Erro ao criar/verificar tabela 'curso': "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT idcurso, nome, descricao, cargaHoraria, valor FROM curso")
	if err != nil {
		http.Error(w, "Erro ao executar SELECT: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cursos []models.Curso
	for rows.Next() {
		var c models.Curso
		if err := rows.Scan(
			&c.Idcurso,
			&c.Nome,
			&c.Descricao,
			&c.CargaHoraria,
			&c.Valor,
		); err != nil {
			http.Error(w, "Erro ao escanear linha: "+err.Error(), http.StatusInternalServerError)
			return
		}
		cursos = append(cursos, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cursos)
}

func GetCursoById(w http.ResponseWriter, r *http.Request) {
	// ======== CORS ========
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, "Erro ao conectar no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT idcurso, nome, descricao, cargaHoraria, valor FROM curso WHERE idcurso = ?", id)
	var c models.Curso
	err = row.Scan(&c.Idcurso, &c.Nome, &c.Descricao, &c.CargaHoraria, &c.Valor)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Curso não encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao buscar curso: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func CreateCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, "Erro ao conectar no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var c models.Curso
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"INSERT INTO curso (nome, descricao, cargaHoraria, valor) VALUES (?, ?, ?, ?)",
		c.Nome, c.Descricao, c.CargaHoraria, c.Valor,
	)
	if err != nil {
		http.Error(w, "Erro ao inserir curso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Erro ao obter ID inserido: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Curso criado com sucesso",
		"id":      lastID,
	})
}

func UpdateCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, "Erro ao conectar no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var c models.Curso
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.Exec(
		"UPDATE curso SET nome = ?, descricao = ?, cargaHoraria = ?, valor = ? WHERE idcurso = ?",
		c.Nome, c.Descricao, c.CargaHoraria, c.Valor, id,
	)
	if err != nil {
		http.Error(w, "Erro ao atualizar curso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Erro ao verificar linhas afetadas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Curso não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Curso atualizado com sucesso",
	})
}

func DeleteCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, "Erro ao conectar no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("DELETE FROM curso WHERE idcurso = ?", id)
	if err != nil {
		http.Error(w, "Erro ao excluir curso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Erro ao verificar linhas afetadas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Curso não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Curso excluído com sucesso",
	})
}
