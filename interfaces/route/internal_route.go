package route

func (r *routeSetup) InternalRoute() {

	route := r.Group.Group("/internal")

	auth := route.Group("/auth")
	authV1 := auth.Group("/v1")

	authV1.POST("/login", r.internalhandler.ValidateLoginUser)

	reg := route.Group("/reg")

	reg.POST("/add-user", r.internalhandler.AddUser)

}
