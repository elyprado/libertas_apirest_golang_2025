// routers/Filmes_route.go
package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterFilmes(router *mux.Router) {
	router.HandleFunc("/filmes", controllers.GetFilmes).Methods("GET")
	router.HandleFunc("/filmes/{id}", controllers.GetfilmeById).Methods("GET")
	router.HandleFunc("/filmes", controllers.Createfilme).Methods("POST")
	router.HandleFunc("/filmes/{id}", controllers.Updatefilme).Methods("PUT")
	router.HandleFunc("/filmes/{id}", controllers.Deletefilme).Methods("DELETE")
}