package route

func (r *routeSetup) InternalRoute() {
	route := r.gin

	auth := route.Group("/auth")
	authV1 := auth.Group("/v1")

	authV1.POST("")
}
