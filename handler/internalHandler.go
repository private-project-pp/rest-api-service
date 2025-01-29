package handler

import "github.com/private-project-pp/pos-rest-api-service/usecase/authentication"

type InternalHandler struct {
	auth authentication.Authentication
}

func SetupInternalHandler(auth authentication.Authentication) *InternalHandler {
	return &InternalHandler{
		auth: auth,
	}
}
