package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	SetupRouterUsuario(router)
	SetupRouterMarca(router)

	router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(
			http.Dir("./static/"))))

	return router
}

func SetupRouterMatheus() *mux.Router {
	router := mux.NewRouter()
	SetupRouterVendedor(router)

	router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(
			http.Dir("./static/"))))

	return router
}


