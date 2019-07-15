package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/local/go-postgre/application/models"
	"github.com/spf13/viper"
)

func Register(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, err
	}

	response := user.CreateAccount()
	return response, nil
}

func Login(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	log.Println(user)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, err
	}
	response := user.Login()
	return response, nil
}

func Me(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	tokenHeader := r.Header.Get("Authorization")

	headerAuthorizationString := strings.Split(tokenHeader, " ")
	token := headerAuthorizationString[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("api.secret_key")), nil
	})

	if err != nil {
		log.Fatalln("errornya is : ", err)
	}
	return claims, nil
}
