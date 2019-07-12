package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/local/go-postgre/application/api"
	"github.com/spf13/viper"
)

func Route() {
	r := mux.NewRouter()

	routers := r.PathPrefix("/api").Subrouter()

	// cek for middleware
	// routers.Use(middleware.JwtAuthentication)
	// modul rooms
	routers.Handle("/getrooms", HandlerFunc(api.GetRooms)).Methods(http.MethodGet)
	routers.Handle("/insertroom", HandlerFunc(api.InsertRooms)).Methods(http.MethodPost)
	routers.Handle("/getroom/{rm_id}", HandlerFunc(api.GetRoom)).Methods(http.MethodGet)
	routers.Handle("/updateroom/{rm_id}", HandlerFunc(api.UpdateRooms)).Methods(http.MethodPost)
	routers.Handle("/deleteroom/{rm_id}", HandlerFunc(api.DeleteRoom)).Methods(http.MethodGet)
	routers.Handle("/deleteallroom", HandlerFunc(api.DeleteAllRoom)).Methods(http.MethodGet)

	// modul users
	routers.Handle("/user/register", HandlerFunc(api.Register)).Methods(http.MethodPost)
	routers.Handle("/user/login", HandlerFunc(api.Login)).Methods(http.MethodPost)
	routers.Handle("/user", HandlerFunc(api.Me)).Methods(http.MethodGet)

	host := fmt.Sprintf(viper.GetString("app.host"))

	srv := &http.Server{
		Handler:      routers,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
