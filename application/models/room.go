package models

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Rooms struct {
	RmID      int    `json:"rm_id,omitemp"`
	RmName    string `json:"rm_name,omitemp"`
	RmPlace   string `json:"rm_place,omitemp"`
	RmSumpart int    `json:"rm_sumpart,omitemp"`
	RmPrice   int    `json:"rm_price,omitemp"`
	RmStatus  int    `json:"rm_status,omitemp"`
}

type ManyRooms []Rooms

func GetRooms(w http.ResponseWriter, r *http.Request) (ManyRooms, error) {
	sql := "select rm_id,rm_name,rm_place,rm_sumpart,rm_price, rm_status FROM rooms"
	data, err := DB.Query(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error Query, and %s", err)
		return nil, errors.New(saveError)
	}
	var manyRooms ManyRooms
	for data.Next() {
		var perRoom Rooms
		err = data.Scan(&perRoom.RmID, &perRoom.RmName, &perRoom.RmPlace, &perRoom.RmSumpart, &perRoom.RmPrice, &perRoom.RmStatus)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and %s", err)
			return nil, errors.New(saveError)
		}
		manyRooms = append(manyRooms, perRoom)
	}
	return manyRooms, nil
}

func GetRoom(r http.ResponseWriter, h *http.Request) (*Rooms, error) {
	params := mux.Vars(h)
	rmID := params["rm_id"]
	var room Rooms
	sql := "SELECT rm_id,rm_name,rm_place,rm_sumpart,rm_price, rm_status FROM rooms WHERE rm_id = $1"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	err = statement.QueryRow(rmID).Scan(&room.RmID, &room.RmName, &room.RmPlace, &room.RmSumpart, &room.RmPrice, &room.RmStatus)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func UpdateRooms(r http.ResponseWriter, h *http.Request) (map[string]interface{}, error) {
	params := mux.Vars(h)
	rmID := params["rm_id"]
	rmName := h.FormValue("rm_name")
	rmPlace := h.FormValue("rm_place")
	rmSumpart := h.FormValue("rm_sumpart")
	rmPrice := h.FormValue("rm_price")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE rooms SET rm_name = $1, rm_place = $2, rm_sumpart = $3, rm_price = $4, updated_at = $5 WHERE rm_id = $6 "
	statement, err := DB.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	statement.Exec(rmName, rmPlace, rmSumpart, rmPrice, updatedAt, rmID)
	return nil, nil
}

func InsertRooms(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	rmID := r.FormValue("rm_id")
	rmName := r.FormValue("rm_name")
	rmPlace := r.FormValue("rm_place")
	rmSumpart := r.FormValue("rm_sumpart")
	rmPrice := r.FormValue("rm_price")
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	deletedAt := time.Now().Format("2006-01-02 15:04:05")
	rmStatus := "1"

	sql := fmt.Sprintf("INSERT INTO rooms (rm_id, rm_name, rm_place, rm_sumpart, rm_price, created_at, updated_at,deleted_at, rm_status) VALUES (%s, '%s', '%s', %s, %s, '%s', '%s', '%s', '%s'); ",
		rmID, rmName, rmPlace, rmSumpart, rmPrice, createdAt, updatedAt, deletedAt, rmStatus)
	_, errs := DB.Query(sql)
	if errs != nil {
		log.Println("yang error adalah insert", errs)
		return nil, errs
	}
	// defer DB.Close()
	return nil, nil
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) (string, error) {
	params := mux.Vars(r)
	rm_id := params["rm_id"]

	sql := "DELETE FROM rooms WHERE rm_id = $1"
	statement, err := DB.Prepare(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error Query Deleted, and %s", err)
		return "", errors.New(saveError)
	}
	statement.Exec(rm_id)
	// defer DB.Close()
	return "Berhasil dihapus", nil
}

func DeleteAllRoom(w http.ResponseWriter, r *http.Request) (string, error) {
	sql := "DELETE FROM rooms"
	_, err := DB.Query(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error Query Deleted, and %s", err)
		return "", errors.New(saveError)
	}
	// defer DB.Close()
	return "Semua data Rooms Berhasil dihapus", nil
}
