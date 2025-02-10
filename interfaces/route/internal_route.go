package route

func (r *routeSetup) InternalRoute() {

	route := r.echo.Group("/internal")

	auth := route.Group("/auth")
	authV1 := auth.Group("/v1")

	authV1.POST("/login", r.internalhandler.ValidateLoginUser)
}
