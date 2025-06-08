package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterAluno(router *mux.Router) {
	router.HandleFunc("/alunos", controllers.GetAlunos).Methods("GET")
	router.HandleFunc("/alunos/{id}", controllers.GetAlunoById).Methods("GET")
	router.HandleFunc("/alunos", controllers.CreateAluno).Methods("POST")
	router.HandleFunc("/alunos/{id}", controllers.UpdateAluno).Methods("PUT")
	router.HandleFunc("/alunos/{id}", controllers.DeleteAluno).Methods("DELETE")
}
