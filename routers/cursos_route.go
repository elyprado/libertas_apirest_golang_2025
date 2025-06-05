// routers/cursos_route.go
package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterCurso(router *mux.Router) {
	router.HandleFunc("/cursos", controllers.GetCursos).Methods("GET")
	router.HandleFunc("/cursos/{id}", controllers.GetCursoById).Methods("GET")
	router.HandleFunc("/cursos", controllers.CreateCurso).Methods("POST")
	router.HandleFunc("/cursos/{id}", controllers.UpdateCurso).Methods("PUT")
	router.HandleFunc("/cursos/{id}", controllers.DeleteCurso).Methods("DELETE")
}
