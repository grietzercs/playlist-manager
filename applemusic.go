// this file will be appropriately placed in it's own folder later with better naming
package main

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
	"crypto/x509"
	"crypto/ecdsa"
	"encoding/pem"
	"errors"
	"log"
)

// still for local testing only
// a function for general use will come later
func generateToken() (string, error) {
	teamId := os.Getenv("appleTeamTag")
	keyId := os.Getenv("appleIdentifier")
	privKey, err := os.ReadFile("AuthKey.p8")
	if err != nil {
		log.Fatal(err)
	}
	
	token := jwt.Token{
		Method: jwt.SigningMethodES256,
		Header: map[string]interface{} {
			"alg":jwt.SigningMethodES256.Alg(),
			"kid":keyId,
		},
		Claims: jwt.MapClaims {
			"iss":teamId,
			"iat":time.Now().Unix(),
			"exp":time.Now().Add(time.Hour * 24).Unix(),
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