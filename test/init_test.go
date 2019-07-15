package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/local/go-postgre/application/api"

	mw "github.com/local/go-postgre/application/middleware"
	conf "github.com/local/go-postgre/application/models"
	"github.com/local/go-postgre/routes"
	"github.com/local/go-postgre/utils"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	DB      *sql.DB
	cfg     testCaseParams
)

type testCaseParams struct {
	RoomID string
	URL    string
}

type testCase struct {
	name         string
	input        string
	expectedData string
	expectedCode int
	path         string
	handler      routes.HandlerFunc
	query        string
}
type Response struct {
	Response Rest                   `json:"response"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

type Rest struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

type DataResponse struct {
	RmID      string
	RmName    string
	RmPlace   string
	RmSumpart string
	RmPrice   string
	RmStatus  string
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
	cfg = testCaseParams{
		RoomID: viper.GetString("testing.room_id"),
		URL:    viper.GetString("app.host"),
	}
}
func TestGetRoom(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Testing with user id",
			input:        cfg.RoomID,
			expectedData: cfg.RoomID,
			expectedCode: http.StatusOK,
			path:         "api/getroom",
			handler:      api.GetRoom,
			query:        "",
		},
		{
			name:         "Testing with random id",
			input:        "9897",
			expectedData: "<nil>", // because not value
			expectedCode: http.StatusBadRequest,
			path:         "api/getroom",
			handler:      api.GetRoom,
			query:        "",
		},
		{
			name:         "Testing with variable",
			input:        "abc",
			expectedData: "<nil>", // because not value
			expectedCode: http.StatusBadRequest,
			path:         "api/getroom",
			handler:      api.GetRoom,
			query:        "",
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			// url := fmt.Sprintf("http://%s/%s?%s", cfg.URL, tc.input, tc.query)
			url := fmt.Sprintf("http://%s/%s/%s", cfg.URL, tc.path, tc.input)
			resp := getRequest(url, "", tc.handler)
			if resp.Code != tc.expectedCode {
				// if resp.Code != 99 {
				t.Errorf("Expected:%d , But Got :%d - message:%v", tc.expectedCode, resp.Code, resp.Body)
			}

			buf := resp.Body.Bytes()
			var respData Response
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}

			getData := fmt.Sprintf("%v", respData.Data["rm_id"])
			if tc.expectedData != getData {
				t.Errorf("Data account id is invalid: Expected: %s, but got %s", tc.expectedData, getData)
			}
			// for k, v := range respData.Data {
			// 	if k == "rm_id" {
			// 		log.Println(k, v)
			//
			// 	}
			// }

		})
	}
}

func getRequest(url, path string, handler routes.HandlerFunc) *httptest.ResponseRecorder {
	// req, err := http.NewRequest("GET", "http://127.0.0.1:8000/api/getroom/1", nil)
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
