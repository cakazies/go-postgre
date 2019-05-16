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
	Rm_id      int    `json:"rm_id"`
	Rm_name    string `json:"rm_name"`
	Rm_place   string `json:"rm_place"`
	Rm_sumpart int    `json:"rm_sumpart"`
	Rm_price   int    `json:"rm_price"`
	Rm_status  int    `json:"rm_status"`
}

type ManyRooms []Rooms

func GetRooms(w http.ResponseWriter, r *http.Request) (ManyRooms, error) {
	sql := "select rm_id,rm_name,rm_place,rm_sumpart,rm_price, rm_status FROM rooms"
	data, err := db.Query(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error Query, and ", err)
		return nil, errors.New(saveError)
	}
	var manyRooms ManyRooms
	// defer db.Close()

	for data.Next() {
		var perRoom Rooms
		err = data.Scan(&perRoom.Rm_id, &perRoom.Rm_name, &perRoom.Rm_place, &perRoom.Rm_sumpart, &perRoom.Rm_price, &perRoom.Rm_status)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and ", err)
			return nil, errors.New(saveError)
		}
		manyRooms = append(manyRooms, perRoom)
	}
	return manyRooms, nil
}

func GetRoom(r http.ResponseWriter, h *http.Request) (*Rooms, error) {
	params := mux.Vars(h)
	rm_id := params["rm_id"]
	var room Rooms
	sql := "SELECT rm_id,rm_name,rm_place,rm_sumpart,rm_price, rm_status FROM rooms WHERE rm_id = $1"
	statement, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	err = statement.QueryRow(rm_id).Scan(&room.Rm_id, &room.Rm_name, &room.Rm_place, &room.Rm_sumpart, &room.Rm_price, &room.Rm_status)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &room, nil
}

func UpdateRooms(r http.ResponseWriter, h *http.Request) (map[string]interface{}, error) {
	params := mux.Vars(h)
	rm_id := params["rm_id"]
	rm_name := h.FormValue("rm_name")
	rm_place := h.FormValue("rm_place")
	rm_sumpart := h.FormValue("rm_sumpart")
	rm_price := h.FormValue("rm_price")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE rooms SET rm_name = $1, rm_place = $2, rm_sumpart = $3, rm_price = $4, updated_at = $5 WHERE rm_id = $6 "
	statement, err := db.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	statement.Exec(rm_name, rm_place, rm_sumpart, rm_price, updated_at, rm_id)

	// defer db.Close()
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
	fmt.Println("ini adalah cek")
	fmt.Println(sql)
	_, errs := db.Query(sql)
	if errs != nil {
		log.Println("yang error adalah insert", errs)
		panic(errs)
		return nil, errs
	}
	// defer db.Close()
	return nil, nil
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) (string, error) {
	params := mux.Vars(r)
	rm_id := params["rm_id"]

	sql := "DELETE FROM rooms WHERE rm_id = $1"

	statement, err := db.Prepare(sql)

	if err != nil {
		saveError := fmt.Sprintf("Error Query Deleted, and ", err)
		return "", errors.New(saveError)
	}
	statement.Exec(rm_id)
	// defer db.Close()
	return "Berhasil dihapus", nil
}

func DeleteAllRoom(w http.ResponseWriter, r *http.Request) (string, error) {
	sql := "DELETE FROM rooms"
	_, err := db.Query(sql)

	if err != nil {
		saveError := fmt.Sprintf("Error Query Deleted, and ", err)
		return "", errors.New(saveError)
	}
	// defer db.Close()
	return "Semua data Rooms Berhasil dihapus", nil
}
