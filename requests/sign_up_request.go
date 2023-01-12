package requests

type SignUpRequest struct {
	//	required: true
	//	example: User
	Name string `json:"name"`
	SessionsRequest
	//	required: true
	C_Password string `json:"c_password"`
}
