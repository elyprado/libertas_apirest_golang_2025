package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterMedicamento(router *mux.Router) {
	router.HandleFunc("/medicamentos", controllers.GetMedicamentos).Methods("GET")
	router.HandleFunc("/medicamentos/{id}", controllers.GetMedicamentoById).Methods("GET")
	router.HandleFunc("/medicamentos", controllers.CreateMedicamento).Methods("POST")
	router.HandleFunc("/medicamentos/{id}", controllers.UpdateMedicamento).Methods("PUT")
	router.HandleFunc("/medicamentos/{id}", controllers.DeleteMedicamento).Methods("DELETE")
}
