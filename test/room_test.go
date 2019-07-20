package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/local/go-postgre/application/api"
	"github.com/local/go-postgre/routes"
)

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

func TestDeleteRoom(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Testing Insert with data valid",
			input:        "5",
			expectedData: "5",
			expectedCode: http.StatusOK,
			path:         "api/deleteroom",
			handler:      api.DeleteRoom,
			query:        "",
		},
		{
			name:         "Testing Delete with Data Invalid",
			input:        "9897",
			expectedData: "<nil>",
			expectedCode: http.StatusBadRequest,
			path:         "api/deleteroom",
			handler:      api.DeleteRoom,
			query:        "",
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			// url := fmt.Sprintf("http://%s/%s?%s", cfg.URL, tc.input, tc.query)
			url := fmt.Sprintf("http://%s/%s/%s", Cfg.URL, tc.path, tc.input)
			resp := postRequest(url, "", tc.handler)
			assert.Equal(t, resp.Code, tc.expectedCode, "Expedted Code is Wrong")
			// if resp.Code != tc.expectedCode {
			// 	t.Errorf("Expected:%d , But Got :%d - message:%v", tc.expectedCode, resp.Code, resp.Body)
			// }
			buf := resp.Body.Bytes()
			var respData Response
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}

		})
	}
}

func TestGetRoom(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Testing with user id",
			input:        Cfg.RoomID,
			expectedData: Cfg.RoomID,
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
			url := fmt.Sprintf("http://%s/%s/%s", Cfg.URL, tc.path, tc.input)
			resp := getRequest(url, "", tc.handler)
			assert.Equal(t, resp.Code, tc.expectedCode, "Expedted Code is Wrong")
			buf := resp.Body.Bytes()
			var respData Response
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}
			getData := fmt.Sprintf("%v", respData.Data["rm_id"])
			assert.Equal(t, getData, tc.expectedData, "Expedted Data is Wrong")
		})
	}
}
