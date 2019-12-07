package test

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cakazies/go-postgre/application/api"
	"github.com/gorilla/mux"

	mw "github.com/cakazies/go-postgre/application/middleware"
	conf "github.com/cakazies/go-postgre/application/models"
	"github.com/cakazies/go-postgre/routes"
	"github.com/cakazies/go-postgre/utils"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	DB      *sql.DB
	Cfg     testCaseParams
)

type testCaseParams struct {
	RoomID string
	URL    string
}

func initCOnfig() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("../configs")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
}

func TestInit(t *testing.T) {
	initCOnfig()
	conf.Connect()
	Cfg = testCaseParams{
		RoomID: viper.GetString("testing.room_id"),
		URL:    viper.GetString("app.host"),
	}
}

func getRequest(url, path string, handler routes.HandlerFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.Handle(path, routes.HandlerFunc(api.GetRoom))
	r.Handle(path, routes.HandlerFunc(handler))
	r.Use(mw.JwtAuthentication)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func postRequest(url, path string, handler routes.HandlerFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.Handle(path, routes.HandlerFunc(api.GetRoom))
	r.Handle(path, routes.HandlerFunc(handler))
	r.Use(mw.JwtAuthentication)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
