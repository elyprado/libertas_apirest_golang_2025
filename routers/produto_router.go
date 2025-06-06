package routers

import (
	"apigolang/controllers"
	"github.com/gorilla/mux"
)

func SetupRouterProdutos(router *mux.Router) {
	router.HandleFunc("/produtos", controllers.GetProdutos).Methods("GET")
	router.HandleFunc("/produtos/{id}", controllers.GetProdutoById).Methods("GET")
	router.HandleFunc("/produtos", controllers.CreateProduto).Methods("POST")
	router.HandleFunc("/produtos/{id}", controllers.UpdateProduto).Methods("PUT")
	router.HandleFunc("/produtos/{id}", controllers.DeleteProduto).Methods("DELETE")
}
