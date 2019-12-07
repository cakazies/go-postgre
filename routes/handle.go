package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cakazies/go-postgre/application/api"
)

type (
	// HandlerFunc ...
	HandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)
)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var api api.Response
	var errs []string

	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	resp, err := fn(w, r)
	if err != nil {
		errs = append(errs, err.Error())
		api.Response.Code = strconv.Itoa(http.StatusBadRequest)
		api.Response.Message = string(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		api.Data = resp
		api.Response.Code = strconv.Itoa(http.StatusOK)
		api.Response.Message = "Success"
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(&api); err != nil {
		log.Println(err)
		return
	}
}
