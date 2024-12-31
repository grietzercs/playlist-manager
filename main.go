package main

import (
	"log"
)

type Song struct {
	Name, Artist, Album string
}

func main() {
	secrets := getCredsFromFile()
	base64Creds := getb64Creds(secrets.ClientID, secrets.ClientSecret)
	getSpotifyUserToken(base64Creds)

	aMusicToken, err := generateToken()
	if err != nil {
		log.Fatal("Failed to get apple music token from function 'generateToken'\n", err)
	}
	print(aMusicToken)
}
