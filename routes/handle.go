package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/local/go-postgre/api"
)

type (
	HandlerFunc func(http.ResponseWriter, *http.Request) (map[string]interface{}, error)
)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var api api.Response
	var errs []string

	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	resp, err := fn(w, r)

	if err != nil {
		errs = append(errs, err.Error())
		api.Response = errs
	} else {
		api.Data = resp
		api.Response = "ok"
	}
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(&api); err != nil {
		log.Println(err)
		return
	}
}
