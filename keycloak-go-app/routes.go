package main

func initializeRoutes() {

	router.GET("/login", login)

	// Test API from browser witout sending credentials payload
	router.GET("/loginWeb", login)

	router.GET("/status", status)
	router.GET("/health", health)

	authRoute := router.Group("/auth")
	{
		authRoute.GET("/quote", ValidateToken(), getQuote)
		authRoute.GET("/logout", ValidateToken(), TokenRevoke(), logout)
	}

}
