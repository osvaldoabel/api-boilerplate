package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"osvaldoabel/users-api/src/presenters"
	"osvaldoabel/users-api/src/services"
	"osvaldoabel/users-api/utils"

	"github.com/gorilla/mux"
)

type UserController struct {
}

func getUserPayloads(r *http.Request) (*utils.UserPayload, error) {
	payload := &utils.UserPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	defer r.Body.Close()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return payload, nil
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {

	payload, err := getUserPayloads(r)
	if err != nil {
		utils.JsonResponse(w, nil, 400)
	}

	uService := services.NewUserService()
	user, err := uService.Insert(payload)
	if err != nil {
		utils.JsonResponse(w, nil, 500)
		return
	}

	result, err := json.Marshal(presenters.ToArray(user))
	if err != nil {
		utils.JsonResponse(w, nil, 500)
		return
	}

	utils.JsonResponse(w, result, 200)
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

	results, err := json.Marshal(presenters.ToCollection(users))
	if err != nil {
		http.Error(w, "Ops... Sorry, we have an Internal Server Error!", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, results, 200)
}

func (u *UserController) Show(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	result := []byte{}

	uService := services.NewUserService()

	user, err := uService.Find(params["id"])
	if err != nil {
		utils.JsonResponse(w, result, 400)
		return
	}
	if user.ID == "" {
		utils.JsonResponse(w, result, 404)
		return
	}

	result, err = json.Marshal(presenters.ToArray(user))
	if err != nil {
		utils.JsonResponse(w, result, 400)
		return
	}

	utils.JsonResponse(w, result, 200)
}

func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	uService := services.NewUserService()
	userPayload, err := getUserPayloads(r)

	if err != nil {
		utils.JsonResponse(w, nil, 400)
		return
	}

	user, err := uService.Update(params["id"], userPayload)
	result, err := json.Marshal(presenters.ToArray(user))

	utils.JsonResponse(w, result, 200)
}

func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uService := services.NewUserService()
	err := uService.Delete(params["id"])

	statusCode := 204

	if err != nil {
		statusCode = 404
	}

	utils.JsonResponse(w, nil, statusCode)
}
