package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Secrets struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Username     string `json:"username"`
}

const apiEndpoint = "https://api.spotify.com/v1"
const apiAuthEndpoint = "https://accounts.spotify.com/api/token"
const callbackURI = "http://localhost:5173/callback"

func getCredsFromFile() Secrets {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secrets := Secrets{
		os.Getenv("clientID"),
		os.Getenv("clientSecret"),
		os.Getenv("Username"),
	}
	return secrets
}

func getb64Creds(id string, pass string) string {
	authFormat := id + ":" + pass
	b64Creds := base64.URLEncoding.EncodeToString([]byte(authFormat))
	return b64Creds
}

func getSpotifyUserToken(creds string) string {
	authHeader := "Basic " + creds
	params := url.Values{}
	params.Add("grant_type", "client_credentials")
	var responseData map[string]interface{}

	request, err := http.NewRequest("POST", apiAuthEndpoint, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Fatal("Error creating new http request with url: ", apiEndpoint)
		return ""
	}
	request.Header.Add("Authorization", authHeader)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error on response", err)
		return ""
	} else {
		defer response.Body.Close()
		data, _ := io.ReadAll(response.Body)
		json.Unmarshal(data, &responseData)
		fmt.Println("Status: ", response.Status)
		fmt.Println("Data: ", responseData["access_token"])
	}
	// maybe can use gob package for more simpler return
	return responseData["access_token"].(string)
}
