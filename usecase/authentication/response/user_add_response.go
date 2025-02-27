package response

type UserAddResponse struct {
	UserId  string `json:"user_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
