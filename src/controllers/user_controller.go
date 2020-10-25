package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"osvaldoabel/users-api/src/presenters"
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
	if params["Limit"] == "" {
		params["Limit"] = "10"
	}

	params["Offset"] = r.URL.Query().Get("page")
	if params["Offset"] == "" {
		params["Offset"] = "0"
	}

	params["OrderBy"] = r.URL.Query().Get("order_by")
	if params["OrderBy"] == "" {
		params["OrderBy"] = "ID"
	}

	uService := services.NewUserService()
	users := uService.All(params)
	results, err := json.Marshal(presenters.ToArray(users))

	if err != nil {
		http.Error(w, "Ops... Sorry, we have an Internal Server Error!", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, results, 200)
}

// func (u *UserController) Update(w http.ResponseWriter, r *http.Request) error {

// }

// func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) error {

// }
