package users_administration

import "github.com/private-project-pp/pos-rest-api-service/infrastructures/user"

type userAdm struct {
	userInfra user.UserInfra
}

func SetupUsersAdministration(userInfra user.UserInfra) UsersAdministration {
	return &userAdm{
		userInfra: userInfra,
	}
}
