package authentication

import "github.com/private-project-pp/pos-rest-api-service/usecase/authentication/response"

func (s auth) UserList() (out response.UsersList, err error) {
	var result = make([]response.UserDataResponse, 0)
	out = response.UsersList{
		Users: result,
	}
	return out, nil
}
