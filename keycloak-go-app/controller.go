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
	ExpiresIn    int    `json:"expiresIn"`
}

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

func logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status": "User loggedout",
	})
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"Health": "OK",
	})
}

func status(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status": "OK",
	})
}

// Calling external API and pringitn its output
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
