package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ParseJWT(tokenString string, APISecret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("An error has occurred.")
		}
		return []byte(APISecret), nil
	})
}

func CreateToken(client string) (string, error) {
	log.Println("Creating token for", client)
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["client"] = client
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if client == os.Getenv("API_SECRET_MACAPA") {
		return token.SignedString([]byte(os.Getenv("API_SECRET_MACAPA")))
	}
	if client == os.Getenv("API_SECRET_VAREJAO") {
		return token.SignedString([]byte(os.Getenv("API_SECRET_VAREJAO")))
	}
	return "", fmt.Errorf("Unexpected client: %s", client)
}

func TokenValid(r *http.Request) error {
	log.Println("Validating token")
	tokenString := ExtractToken(r)
	token, err := ParseJWT(tokenString, os.Getenv("API_SECRET_MACAPA"))
	if err != nil {
		token, err = ParseJWT(tokenString, os.Getenv("API_SECRET_VAREJAO"))
	}
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractClient(r *http.Request) (string, error) {
	tokenString := ExtractToken(r)
	token, err := ParseJWT(tokenString, os.Getenv("API_SECRET_MACAPA"))
	if err != nil {
		token, err = ParseJWT(tokenString, os.Getenv("API_SECRET_VAREJAO"))
	}
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		client := fmt.Sprintf("%v", claims["client"])
		return client, nil
	}
	return "", nil
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
