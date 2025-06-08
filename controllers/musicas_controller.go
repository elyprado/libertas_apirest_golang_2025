package controllers

import (
	"apigolang/config"
	"apigolang/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMusicas(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, titulo, artista, album, ano, genero FROM musicas")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var musicas []models.Musica
	for rows.Next() {
		var musica models.Musica
		err := rows.Scan(&musica.ID, &musica.Titulo, &musica.Artista, &musica.Album, &musica.Ano, &musica.Genero)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		musicas = append(musicas, musica)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musicas)
}

func GetMusicaById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	row := db.QueryRow("SELECT id, titulo, artista, album, ano, genero FROM musicas WHERE id = ?", id)

	var musica models.Musica
	err = row.Scan(&musica.ID, &musica.Titulo, &musica.Artista, &musica.Album, &musica.Ano, &musica.Genero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musica)
}

func CreateMusica(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var musica models.Musica
	err = json.NewDecoder(r.Body).Decode(&musica)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO musicas (titulo, artista, album, ano, genero) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(query, musica.Titulo, musica.Artista, musica.Album, musica.Ano, musica.Genero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "música inserida com sucesso"})
}

func UpdateMusica(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var musica models.Musica
	err = json.NewDecoder(r.Body).Decode(&musica)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE musicas SET titulo = ?, artista = ?, album = ?, ano = ?, genero = ? WHERE id = ?"
	_, err = db.Exec(query, musica.Titulo, musica.Artista, musica.Album, musica.Ano, musica.Genero, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "música atualizada com sucesso"})
}

func DeleteMusica(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM musicas WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "música removida com sucesso"})
}
