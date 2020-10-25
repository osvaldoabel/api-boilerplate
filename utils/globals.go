package utils

import "net/http"

type UserPayload struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Age      int    `json:"age"`
	Status   string `json:"status"`
}

type UserCollection struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	Status  string `json:"status"`
}

type DefaultQueryString struct {
	Limit  string `json:"per_page"`
	Offset string `json:"page"`
}

func JsonResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
