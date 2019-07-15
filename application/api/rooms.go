package api

import (
	"net/http"

	"github.com/local/go-postgre/application/models"
)

func GetRooms(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	data, err := models.GetRooms(w, r)
	// var paramsGet = map[string]interface{}{"data": data}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetRoom(r http.ResponseWriter, h *http.Request) (interface{}, error) {
	data, err := models.GetRoom(r, h)
	// var getData = map[string]interface{}{"data": data}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertRooms(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	_, err := models.InsertRooms(w, r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdateRooms(r http.ResponseWriter, h *http.Request) (interface{}, error) {
	_, err := models.UpdateRooms(r, h)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	message, err := models.DeleteRoom(w, r)
	if err != nil {
		return nil, err
	}
	// var convert = map[string]interface{}{"resp": massage}
	return message, nil
}

func DeleteAllRoom(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	message, err := models.DeleteAllRoom(w, r)
	if err != nil {
		return nil, err
	}
	// var convert = map[string]interface{}{"resp": massage}
	return message, nil
}
