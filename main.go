package main

import (
	"log"
	"playlist-manager/apple-music"
)

type Song struct {
	Name, Artist, Album string
}

func main() {
	secrets := getCredsFromFile()
	base64Creds := getb64Creds(secrets.ClientID, secrets.ClientSecret)
	getSpotifyUserToken(base64Creds)

	aMusicToken, err := applemusic.GenerateToken()
	if err != nil {
		log.Fatal("Failed to get apple music token from function 'generateToken'\n", err)
	}
	applemusic.TestAuthorization(aMusicToken)
	applemusic.GetTopStorePlaylists(aMusicToken)
}