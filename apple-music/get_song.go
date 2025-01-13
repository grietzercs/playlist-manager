package applemusic

import (
	"log"
	"fmt"
	"net/http"
	"io"
)

// should return something like string[]interface{}
// may need user token as well
// reference: https://developer.apple.com/documentation/applemusicapi/get_a_library_playlist
func GetUserPlaylists(user string, authToken string) {
	// should be enough for the api to give it to me?
	apiEndpoint := fmt.Sprintf("https://api.music.apple.com/v1/%s/library/playlists", user)
	request, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatal("request not valid")
	}
	authHeader := fmt.Sprintf("Bearer %s", authToken)
	request.Header.Add("Authorization", authHeader)
	//request.Header.Add("Cot")
	
	//client := &http.Client{}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal("Error on response", err)
	} else {
		defer response.Body.Close()
		data, _ := io.ReadAll(response.Body)
		fmt.Printf("response body: %s", response.Body)
		fmt.Printf("Response data: %s", data)
	}
}

func GetTopStorePlaylists(authToken string) {
	// should be enough for the api to give it to me?
	apiEndpoint := fmt.Sprintf("https://api.music.apple.com/v1/catalog/us/albums")
	request, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatal("request not valid")
	}
	authHeader := fmt.Sprintf("Bearer %s", authToken)
	request.Header.Add("Authorization", authHeader)
	//request.Header.Add("Cot")
	
	//client := &http.Client{}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal("Error on response", err)
	} else {
		defer response.Body.Close()
		data, _ := io.ReadAll(response.Body)
		fmt.Printf("response body: %s", response.Body)
		fmt.Printf("Response data: %s", data)
	}
}