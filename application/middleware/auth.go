package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/cakazies/go-postgre/application/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// JwtAuthentication for JWT
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuthPath := []string{"/api/user/register"}
		noAuthPath = append(noAuthPath, "/api/user/login")
		requestPath := r.URL.Path
		// looping for check pathnya
		for _, path := range noAuthPath {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			rsp := map[string]interface{}{"status": "invalid", "message": "Token is not Present ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		headerAuthorizationString := strings.Split(tokenHeader, " ")
		if len(headerAuthorizationString) != 2 {
			rsp := map[string]interface{}{"status": "invalid", "message": "Invalid/Format Auth Token ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		barier := headerAuthorizationString[0]
		if barier != "Bearer" {
			rsp := map[string]interface{}{"status": "invalid", "message": "Token is not Barier ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		tk := &models.Token{}
		tokenValue := headerAuthorizationString[1]
		token, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("api.secret_key")), nil
		})

		if err != nil {
			rsp := map[string]interface{}{"status": "invalid", "message": "Malformed Authentication Token Please Login Again;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		// check for time expired
		diff := tk.TimeExp.Sub(time.Now())
		if diff < 0 {
			rsp := map[string]interface{}{"status": "invalid", "message": "Time Expired, please login again;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		if !token.Valid {
			rsp := map[string]interface{}{"status": "invalid", "message": "Invalid/Format Auth Token ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		// fmt.Sprintf("User Id is %s", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
