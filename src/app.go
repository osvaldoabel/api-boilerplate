package src

import (
	"osvaldoabel/users-api/src/controllers"

	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	userController := &controllers.UserController{}

	router := mux.NewRouter()

	router.HandleFunc("/users", userController.All).Methods("GET")
	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userController.Show).Methods("GET")          //show
	router.HandleFunc("/users/{id}/update", userController.Update).Methods("PUT") //update
	router.HandleFunc("/users/{id}", userController.Delete).Methods("DELETE")     //delete

	http.Handle("/", router)
}
