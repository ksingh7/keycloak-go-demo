package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"` // Not used anywhere currently
}

// Login using the given username and password and prints access and refresh tokens
func login(c *gin.Context) {

	var requestBody loginRequest
	var responseBody loginResponse

	// Test API from browser witout sending credentials payload
	if c.FullPath() == "/loginWeb" {
		requestBody.Username = goDotEnvVariables("KEYCLOAK_USERNAME")
		requestBody.Password = goDotEnvVariables("KEYCLOAK_PASSWORD")
	} else {
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	AccessToken, RefreshToken, err := keycloakClientLogin(
		requestBody.Username,
		requestBody.Password,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"Status": "Invalid Credentials",
			"Error":  err,
		})
		return
	}

	responseBody.AccessToken = AccessToken
	responseBody.RefreshToken = RefreshToken

	c.JSON(200, gin.H{
		"Status":       "Login Successful",
		"AccessToken":  responseBody.AccessToken,
		"RefreshToken": responseBody.RefreshToken,
	})

}

// Prints logout success status in response body
// Most of the heavy lifting is done in Middleware TokenRevoke() method
func logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status": "User loggedout",
	})
}

// Unauthenticated simple Health Check API
func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"Health": "OK",
	})
}

// Authenticated simple Status Check API
func status(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status": "OK",
	})
}

// Authenticated API to print Quote message by calling getmeaquote external API
func getQuote(c *gin.Context) {

	response, err := http.Get("http://getmeaquote.designedbyaturtle.com/")

	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(responseData))
	c.JSON(200, gin.H{
		"Quote": string(responseData),
	})
}
