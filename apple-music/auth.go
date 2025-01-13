// this file will be appropriately placed in it's own folder later with better naming
package applemusic

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
	"crypto/x509"
	"crypto/ecdsa"
	"encoding/pem"
	"errors"
	"log"
	"net/http"
	"fmt"
	"io"
)

// still for local testing only
// a function for general use will come later
func GenerateToken() (string, error) {
	teamId := os.Getenv("appleTeamTag")
	keyId := os.Getenv("appleIdentifier")
	privKey, err := os.ReadFile("./apple-music/AuthKey.p8")
	if err != nil {
		log.Fatal(err)
	}

	nowTime := time.Now().Unix()
	expTime := time.Now().Add(time.Hour).Unix()
	fmt.Printf("nowTime: %s\texpTime: %s", nowTime, expTime)
	
	token := jwt.Token{
		Method: jwt.SigningMethodES256,
		Header: map[string]interface{} {
			"alg":jwt.SigningMethodES256.Alg(),
			"kid":keyId,
		},
		Claims: jwt.MapClaims {
			"iss":teamId,
			"iat":time.Now().Unix(),
			"exp":time.Now().Add(time.Minute * 10).Unix(),
		},
		Signature: []byte(privKey),
	}

	key, err := parsePKCS8PrivateKeyFromPEM(privKey)
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

func parsePKCS8PrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
	var err error
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errors.New("Invalid Key: Key must be PEM encoded PKCS8 private key")
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var privateKey *ecdsa.PrivateKey
	var validKey bool 
	if privateKey, validKey = parsedKey.(*ecdsa.PrivateKey); !validKey {
		return nil, errors.New("Key is not a valid PKCS8 private key")
	}

	return privateKey, nil
}

func TestAuthorization(token string) bool {
	authURL := "https://api.music.apple.com/v1/test"
	request, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		log.Fatal("Failed test auth", err)
		return false
	}
	authHeader := fmt.Sprintf("Bearer %s", token)
	request.Header.Add("Authorization", authHeader)
	
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error on response", err)
		return false
	} else {
		defer response.Body.Close()
		data, _ := io.ReadAll(response.Body)
		fmt.Printf("response body: %s", response.Body)
		fmt.Printf("Response data: %s", data)
		return true
	}
}