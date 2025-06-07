package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	SetupRouterUsuario(router)
	SetupRouterMarca(router)
	SetupRouterTime(router)
	SetupRouterVendedor(router)
	SetupRouterMedicamento(router)
	SetupRouterListaCompras(router)
	SetupRouterImovel(router)
	SetupRouterCurso(router) 
	SetupRouterProdutos(router)
	SetupRouterVeiculo(router)
	SetupRouterFilmes(router) 
	router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(
			http.Dir("./static/"))))

	return router
}
