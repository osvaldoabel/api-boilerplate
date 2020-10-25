package src

import (
	"net/http"
	"osvaldoabel/users-api/src/controllers"

	"github.com/gorilla/mux"
)

func Start() {

	userController := &controllers.UserController{}

	router := mux.NewRouter()

	router.HandleFunc("/v1/users", userController.All).Methods("GET")
	router.HandleFunc("/v1/users", userController.Create).Methods("POST")
	router.HandleFunc("/v1/users/{id}", userController.Show).Methods("GET")          //show
	router.HandleFunc("/v1/users/{id}/update", userController.Update).Methods("PUT") //update
	router.HandleFunc("/v1/users/{id}", userController.Delete).Methods("DELETE")     //delete

	http.Handle("/", router)
}
