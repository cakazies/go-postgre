package api

import (
	"net/http"

	"github.com/cakazies/go-postgre/application/models"
)

// GetRooms function for get all room table
func GetRooms(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	data, err := models.GetRooms(w, r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetRoom function for get single room
func GetRoom(r http.ResponseWriter, h *http.Request) (interface{}, error) {
	data, err := models.GetRoom(r, h)
	if err != nil {
		SentryInit(err)
		return nil, err
	}
	return data, nil
}

// InsertRooms function for insert in table room
func InsertRooms(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	_, err := models.InsertRooms(w, r)
	if err != nil {
		SentryInit(err)
		return nil, err
	}
	return nil, nil
}

// UpdateRooms function for update data in table room
func UpdateRooms(r http.ResponseWriter, h *http.Request) (interface{}, error) {
	_, err := models.UpdateRooms(r, h)
	if err != nil {
		SentryInit(err)
		return nil, err
	}
	return nil, nil
}

// DeleteRoom function for delete data in table room
func DeleteRoom(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	message, err := models.DeleteRoom(w, r)
	if err != nil {
		SentryInit(err)
		return nil, err
	}
	var convert = map[string]interface{}{"message": message}
	return convert, nil
}

// DeleteAllRoom function for delete all data in table room
func DeleteAllRoom(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	message, err := models.DeleteAllRoom(w, r)
	if err != nil {
		SentryInit(err)
		return nil, err
	}
	var convert = map[string]interface{}{"message": message}
	return convert, nil
}
