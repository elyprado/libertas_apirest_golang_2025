package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterImovel(router *mux.Router) {
	router.HandleFunc("/imoveis", controllers.GetImoveis).Methods("GET")
	router.HandleFunc("/imoveis/{id}", controllers.GetImovelById).Methods("GET")
	router.HandleFunc("/imoveis", controllers.CreateImovel).Methods("POST")
	router.HandleFunc("/imoveis/{id}", controllers.UpdateImovel).Methods("PUT")
	router.HandleFunc("/imoveis/{id}", controllers.DeleteImovel).Methods("DELETE")
}

