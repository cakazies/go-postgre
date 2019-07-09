package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	ctr "github.com/local/go-postgre/api"
	"github.com/spf13/viper"
)

func Route() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	// modul rooms
	api.Handle("/getrooms", HandlerFunc(ctr.GetRooms)).Methods(http.MethodGet)
	api.Handle("/insertroom", HandlerFunc(ctr.InsertRooms)).Methods(http.MethodPost)
	api.Handle("/getroom/{rm_id}", HandlerFunc(ctr.GetRoom)).Methods(http.MethodGet)
	api.Handle("/updateroom/{rm_id}", HandlerFunc(ctr.UpdateRooms)).Methods(http.MethodPost)
	api.Handle("/deleteroom/{rm_id}", HandlerFunc(ctr.DeleteRoom)).Methods(http.MethodGet)
	api.Handle("/deleteallroom", HandlerFunc(ctr.DeleteAllRoom)).Methods(http.MethodGet)

	// modul users
	api.Handle("/user/register", HandlerFunc(ctr.DeleteRoom)).Methods(http.MethodPost)
	api.Handle("/user/login", HandlerFunc(ctr.DeleteAllRoom)).Methods(http.MethodPost)

	host := fmt.Sprintf(viper.GetString("app.host"))

	srv := &http.Server{
		Handler:      r,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
