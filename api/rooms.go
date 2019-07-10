package api

import (
	"log"
	"net/http"

	"github.com/local/go-postgre/models"
)

func GetRooms(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	data, err := models.GetRooms(w, r)
	var paramsGet = map[string]interface{}{"data": data}
	if err != nil {
		return nil, err
	}
	return paramsGet, nil
}

func GetRoom(r http.ResponseWriter, h *http.Request) (map[string]interface{}, error) {
	data, err := models.GetRoom(r, h)
	var getData = map[string]interface{}{"data": data}
	if err != nil {
		return nil, err
	}
	return getData, nil
}

func InsertRooms(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	_, err := models.InsertRooms(w, r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdateRooms(r http.ResponseWriter, h *http.Request) (map[string]interface{}, error) {
	_, err := models.UpdateRooms(r, h)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	massage, err := models.DeleteRoom(w, r)
	if err != nil {
		log.Println("error fungsi router ", err)
		return nil, err
	}
	var convert = map[string]interface{}{"resp": massage}
	return convert, nil
}

func DeleteAllRoom(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	massage, err := models.DeleteAllRoom(w, r)
	if err != nil {
		log.Println("error fungsi router ", err)
		return nil, err
	}
	var convert = map[string]interface{}{"resp": massage}
	return convert, nil
}
