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
	SetupRouterCurso(router) // ✅ Aqui você adiciona o novo módulo

	router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(
			http.Dir("./static/"))))

	return router
}
