package main

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v11"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type keycloakCreds struct {
	hostname     string
	clientId     string
	clientSecret string
	realm        string
	username     string
	password     string
}

var kCreds = &keycloakCreds{
	hostname:     goDotEnvVariables("KEYCLOAK_HOSTNAME"),
	clientId:     goDotEnvVariables("KEYCLOAK_CLIENT_ID"),
	clientSecret: goDotEnvVariables("KEYCLOAK_CLIENT_SECRET"),
	realm:        goDotEnvVariables("KEYCLOAK_REALM"),
}

func keycloakClientLogin(username string, password string) (string, string, error) {

	var keycloakClientLoginCreds = &keycloakCreds{
		username: username,
		password: password,
	}

	keycloakClient := gocloak.NewClient(kCreds.hostname)
	restyClient := keycloakClient.RestyClient()
	restyClient.SetDebug(false)

	kCTX := context.Background()
	jwt, err := keycloakClient.Login(
		kCTX,
		kCreds.clientId,
		kCreds.clientSecret, kCreds.realm,
		keycloakClientLoginCreds.username, keycloakClientLoginCreds.password,
	)

	if err != nil {
		log.Error().Msgf("%v", "keycloakClientLogin() Invalid Credentials", err)
		return "", "", err
	}

	return jwt.AccessToken, jwt.RefreshToken, err

}

func keycloakRetrospectToken(accessToken string) (bool, error) {

	keycloakClient := gocloak.NewClient(kCreds.hostname)
	restyClient := keycloakClient.RestyClient()
	restyClient.SetDebug(false)

	kCTX := context.Background()
	retrospectToken, err := keycloakClient.RetrospectToken(
		kCTX, accessToken,
		kCreds.clientId,
		kCreds.clientSecret, kCreds.realm,
	)

	if err != nil {
		log.Error().Msgf("%v", "keycloakRetrospectToken() Invalid or malformed token", err)
		return false, err
	}

	if *retrospectToken.Active {
		log.Info().Msgf("%v", "keycloakRetrospectToken() Token is active")
		return true, nil
	}

	return false, nil

}

func keycloakClientTokenRevoke(accessToken string) error {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	endpoint := kCreds.hostname + "auth/realms/" + kCreds.realm + "/protocol/openid-connect/revoke"

	data := url.Values{}
	data.Set("client_id", kCreds.clientId)
	data.Set("client_secret", kCreds.clientSecret)
	data.Set("token", accessToken)
	encodedData := data.Encode()
	fmt.Println(encodedData)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(encodedData))
	if err != nil {
		log.Error().Msgf("%v", "Error creating request", err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(req)

	if err != nil {
		log.Error().Msgf("%v", "Error sending request", err)
		return err
	}
	defer response.Body.Close()
	return nil
}
