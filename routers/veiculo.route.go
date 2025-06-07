package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterVeiculo(router *mux.Router) {
	router.HandleFunc("/veiculos", controllers.GetVeiculos).Methods("GET")
	router.HandleFunc("/veiculos/{id}", controllers.GetVeiculoById).Methods("GET")
	router.HandleFunc("/veiculos", controllers.CreateVeiculo).Methods("POST")
	router.HandleFunc("/veiculos/{id}", controllers.UpdateVeiculo).Methods("PUT")
	router.HandleFunc("/veiculos/{id}", controllers.DeleteVeiculo).Methods("DELETE")
}
