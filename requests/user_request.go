package requests

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
