package authentication

type Authentication interface {
}

type auth struct {
}

func SetupAuthentication() Authentication {
	return &auth{}
}
