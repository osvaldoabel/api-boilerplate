package utils

type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Age      int    `json:"age"`
	Status   string `json:"status"`
}
type DefaultQueryString struct {
	Limit  string `json:"per_page"`
	Offset string `json:"page"`
}