package main

func initializeRoutes() {

	// Unauthenticated routes
	router.GET("/login", login)
	router.GET("/health", health)

	// A slight modified version of /login API
	// Use to Test from browser witout sending credentials as payload
	router.GET("/loginWeb", login)

	// Authenticated routes
	authRoute := router.Group("/auth")
	{
		authRoute.GET("/quote", ValidateToken(), getQuote)
		authRoute.GET("/status", ValidateToken(), status)
		authRoute.GET("/logout", ValidateToken(), TokenRevoke(), logout)
	}

}
