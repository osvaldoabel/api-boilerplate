package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/services"
	"osvaldoabel/users-api/utils"
)

type UserController struct {
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {

	payload := &utils.UserPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Proplem with the payload"))
		return
	}

	uService := services.NewUserService()
	uService.Insert(payload)

	w.Header().Set("Content-Type", "application/json")
	// w.Write([])

}

func (u *UserController) All(w http.ResponseWriter, r *http.Request) {

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", 28, "123456")

	result, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	return
}

// func (u *UserController) Update(w http.ResponseWriter, r *http.Request) error {

// }

// func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) error {

// }
