package routers

import (
	"apigolang/controllers"
	//"net/http"
	"github.com/gorilla/mux"
)

func SetupRouterListaCompras(router *mux.Router) {
	router.HandleFunc("/itens", controllers.GetItens).Methods("GET")
	router.HandleFunc("/itens/{id}", controllers.GetItemById).Methods("GET")
	router.HandleFunc("/itens", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/itens/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/itens/{id}", controllers.DeleteItem).Methods("DELETE")
}
