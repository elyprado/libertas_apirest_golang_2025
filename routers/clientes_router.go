package routers

import (
	"clientes-go/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/clientes", controllers.CorsPreflight).Methods("OPTIONS")
	r.HandleFunc("/api/clientes/{id}", controllers.CorsPreflight).Methods("OPTIONS")

	r.HandleFunc("/api/clientes", controllers.GetClientes).Methods("GET")
	r.HandleFunc("/api/clientes/{id}", controllers.GetClienteByID).Methods("GET")
	r.HandleFunc("/api/clientes", controllers.CreateCliente).Methods("POST")
	r.HandleFunc("/api/clientes/{id}", controllers.UpdateCliente).Methods("PUT")
	r.HandleFunc("/api/clientes/{id}", controllers.DeleteCliente).Methods("DELETE")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	return r
}
