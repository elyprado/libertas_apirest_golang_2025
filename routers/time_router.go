package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterTime(router *mux.Router) {
	router.HandleFunc("/time", controllers.GetTime).Methods("GET")
	router.HandleFunc("/time/{idtime}", controllers.GetTimeById).Methods("GET")
	router.HandleFunc("/time", controllers.CreateTime).Methods("POST")
	router.HandleFunc("/time/{idtime}", controllers.UpdateTime).Methods("PUT")
	router.HandleFunc("/time/{idtime}", controllers.DeleteTime).Methods("DELETE")
}
