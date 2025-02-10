package authentication

import (
	"github.com/private-project-pp/pos-rest-api-service/infrastructures/user"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication/response"
)

type Authentication interface {
	UserList() (out response.UsersList, err error)
}

type auth struct {
	userInfra user.UserInfra
}

func SetupAuthentication(userInfra user.UserInfra) Authentication {
	return &auth{
		userInfra: userInfra,
	}
}
