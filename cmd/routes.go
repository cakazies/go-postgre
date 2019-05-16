package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/local/go-postgre/routes"
	"github.com/spf13/viper"
)

func Route() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	// api.HandleFunc("/admin", routes.Index).Methods(http.MethodGet)
	api.HandleFunc("/admin", testHandler).Methods(http.MethodGet)

	// method borrow
	api.HandleFunc("/getborrow", testHandler).Methods(http.MethodGet)
	api.HandleFunc("/borrow", routes.BorIndex).Methods(http.MethodGet)

	// method rooms
	api.Handle("/getrooms", routes.HandlerFunc(routes.GetRooms)).Methods(http.MethodGet)
	api.Handle("/insertroom", routes.HandlerFunc(routes.InsertRooms)).Methods(http.MethodPost)
	api.Handle("/getroom/{rm_id}", routes.HandlerFunc(routes.GetRoom)).Methods(http.MethodGet)
	api.Handle("/updateroom/{rm_id}", routes.HandlerFunc(routes.UpdateRooms)).Methods(http.MethodPost)
	api.Handle("/deleteroom/{rm_id}", routes.HandlerFunc(routes.DeleteRoom)).Methods(http.MethodGet)
	api.Handle("/deleteallroom", routes.HandlerFunc(routes.DeleteAllRoom)).Methods(http.MethodGet)

	host := fmt.Sprintf(viper.GetString("app.host"))

	srv := &http.Server{
		Handler:      r,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("opi")
	log.Println("done")
	return
}
