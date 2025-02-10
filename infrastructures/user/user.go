package user

import (
	userService "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"google.golang.org/grpc"
)

type userInfra struct {
	infra userService.UserServiceClient
}

func SetupUser(clientConn *grpc.ClientConn) UserInfra {
	return &userInfra{
		infra: userService.NewUserServiceClient(clientConn),
	}
}
