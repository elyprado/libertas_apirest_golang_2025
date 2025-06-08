package router

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterMusica(router *mux.Router) {
	router.HandleFunc("/musicas", controllers.GetMusicas).Methods("GET")
	router.HandleFunc("/musicas/{id}", controllers.GetMusicaById).Methods("GET")
	router.HandleFunc("/musicas", controllers.CreateMusica).Methods("POST")
	router.HandleFunc("/musicas/{id}", controllers.UpdateMusica).Methods("PUT")
	router.HandleFunc("/musicas/{id}", controllers.DeleteMusica).Methods("DELETE")
}
