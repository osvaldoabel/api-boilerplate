package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"osvaldoabel/users-api/src/repositories"
	"osvaldoabel/users-api/src/services"
	"osvaldoabel/users-api/utils"
)

type UserController struct {
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {

	payload := &utils.UserPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	defer r.Body.Close()

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Proplem with the payload"))
		return
	}

	uService := services.NewUserService()
	uService.Insert(payload)

	w.Header().Set("Content-Type", "application/json")
}

func (u *UserController) All(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{}

	params["Limit"] = r.URL.Query().Get("per_page")
	params["Offset"] = r.URL.Query().Get("page")
	params["OrderBy"] = r.URL.Query().Get("order_by")

	uService := services.NewUserService()
	uService.UserRepository = repositories.NewUserRepository()

	users := uService.All(params)

	result, err := json.Marshal(users)

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
