package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/services"
	"osvaldoabel/users-api/utils"

	"github.com/fatih/structs"
)

type UserController struct {
}

type UserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {

	payload := &UserPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Proplem with the payload"))
		return
	}

	utils.Dd(payload, true)
	uService := services.NewUserService()
	uService.Insert(structs.Map(payload))

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

}

func (u *UserController) All(w http.ResponseWriter, r *http.Request) {

	// user := domain.User{"Alex", []string{"snowboarding", "programming"}}

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", "123456")

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
