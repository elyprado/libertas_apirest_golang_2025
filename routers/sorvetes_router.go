package routers

import (
	"net/http"
	"sorvetes-go/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/sorvetes", controllers.CorsPreflight).Methods("OPTIONS")
	r.HandleFunc("/api/sorvetes/{id}", controllers.CorsPreflight).Methods("OPTIONS")

	r.HandleFunc("/api/sorvetes", controllers.GetSorvetes).Methods("GET")
	r.HandleFunc("/api/sorvetes/{id}", controllers.GetSorveteByID).Methods("GET")
	r.HandleFunc("/api/sorvetes", controllers.CreateSorvete).Methods("POST")
	r.HandleFunc("/api/sorvetes/{id}", controllers.UpdateSorvete).Methods("PUT")
	r.HandleFunc("/api/sorvetes/{id}", controllers.DeleteSorvete).Methods("DELETE")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	return r
}
