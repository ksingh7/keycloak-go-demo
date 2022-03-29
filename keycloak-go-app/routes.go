package main

func initializeRoutes() {

	// Unauthenticated routes
	router.GET("/login", login)
	router.GET("/health", health)
	// Test API from browser witout sending credentials payload
	router.GET("/loginWeb", login)

	// Authenticated routes
	authRoute := router.Group("/auth")
	{
		authRoute.GET("/quote", ValidateToken(), getQuote)
		authRoute.GET("/status", ValidateToken(), status)
		authRoute.GET("/logout", ValidateToken(), TokenRevoke(), logout)
	}

}
