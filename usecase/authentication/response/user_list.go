package response

type UsersList struct {
	Users []UserDataResponse `json:"users"`
}

type UserDataResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
