package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterMarca(router *mux.Router) {
	router.HandleFunc("/marcas", controllers.GetMarcas).Methods("GET")
	router.HandleFunc("/marcas/{id}", controllers.GetMarcaById).Methods("GET")
	router.HandleFunc("/marcas", controllers.CreateMarca).Methods("POST")
	router.HandleFunc("/marcas/{id}", controllers.UpdateMarca).Methods("PUT")
	router.HandleFunc("/marcas/{id}", controllers.DeleteMarca).Methods("DELETE")
}
