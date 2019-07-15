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
	name           string
	input          string
	expectedData   string
	expectedCode   int
	path           string
	handler        routes.HandlerFunc
	query          string
	checkAccountID bool
}

type ResponseData struct {
	Response string  `json:"response"`
	Data     []Datas `json:"data"`
}
type Datas struct {
	RmID      string `json:"rm_id"`
	RmName    string `json:"rm_name"`
	RmPlace   string `json:"rm_place"`
	RmSumpart string `json:"rm_sumpart"`
	RmPrice   string `json:"rm_price"`
	RmStatus  string `json:"rm_status"`
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
	TestInit(t)
	tasks := []testCase{
		{
			name:           "Testing with user id",
			input:          cfg.RoomID,
			expectedData:   cfg.RoomID,
			expectedCode:   http.StatusOK,
			path:           "api/getroom/{id}",
			handler:        api.GetRoom,
			query:          "",
			checkAccountID: false,
		},
		{
			name:           "Testing with random id",
			input:          "9897",
			expectedData:   "",
			expectedCode:   http.StatusOK,
			path:           "api/getroom/{id}",
			handler:        api.GetRoom,
			query:          "",
			checkAccountID: false,
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("http://%s/%s?%s", cfg.URL, tc.input, tc.query)
			// url := fmt.Sprintf("http://%s/%s/%s", cfg.URL, tc.path, tc.input)
			resp := getRequest(url, "", tc.handler)
			log.Println(resp)
			if resp.Code != tc.expectedCode {
				t.Errorf("Expected:%d , But Got :%d - message:%v", tc.expectedCode, resp.Code, resp.Body)
			}

			buf := resp.Body.Bytes()
			var respData ResponseData
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}

			// for _, v := range respData.Data {
			// 	if tc.checkAccountID {
			// 		if tc.expectedData != v.Datas.RmID {
			// 			t.Errorf("Data account id is invalid: Expected: %s, but got %s", tc.expectedData, v.RmID)
			// 		}
			// 	} else {
			// 		if tc.expectedData != v.RmID {
			// 			t.Errorf("Data is invalid, Expected:%s, but got %s", tc.expectedData, v.RmID)
			// 		}
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
