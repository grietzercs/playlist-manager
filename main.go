package main

import ()

type Song struct {
	Name, Artist, Album string
}

func main() {
	secrets := getCredsFromFile()
	base64Creds := getb64Creds(secrets.ClientID, secrets.ClientSecret)
	sendRequest(base64Creds)
}
