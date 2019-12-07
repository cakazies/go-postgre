package api

import (
	"net/http"

	"github.com/cakazies/go-postgre/application/models"
)

// GetBorrows function for Get borrows
func GetBorrows(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	data, err := models.GetBorrows(w, r)
	// var paramsGet = map[string]interface{}{"data": data}
	if err != nil {
		return nil, err
	}
	return data, nil
}
