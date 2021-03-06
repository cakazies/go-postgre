package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/cakazies/go-postgre/application/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// Register function for register users
func Register(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		SentryInit(err)
		return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, err
	}

	response := user.CreateAccount()
	return response, nil
}

// Login function for login
func Login(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		SentryInit(err)
		return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, err
	}
	response := user.Login()
	return response, nil
}

// Me function for get detail in token
func Me(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	tokenHeader := r.Header.Get("Authorization")

	headerAuthorizationString := strings.Split(tokenHeader, " ")
	token := headerAuthorizationString[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("api.secret_key")), nil
	})

	if err != nil {
		SentryInit(err)
		log.Fatalln("errornya is : ", err)
	}
	return claims, nil
}
