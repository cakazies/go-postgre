package models

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Rooms struct for table ROoms
type Rooms struct {
	RmID      int    `json:"rm_id,omitemp"`
	RmName    string `json:"rm_name,omitemp"`
	RmPlace   string `json:"rm_place,omitemp"`
	RmSumpart int    `json:"rm_sumpart,omitemp"`
	RmPrice   int    `json:"rm_price,omitemp"`
	RmStatus  int    `json:"rm_status,omitemp"`
}

// ManyRooms strtuck for many room array
type ManyRooms []Rooms

// GetRooms function for get all data from table
func GetRooms(w http.ResponseWriter, r *http.Request) (ManyRooms, error) {
	qulimit := ""
	quShort := ""
	qulimit = LimitOffset(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))
	quShort = ShortBy(r.URL.Query().Get("sort_by"))

	sql := "select rm_id,rm_name,rm_place,rm_sumpart,rm_price, rm_status FROM rooms " + quShort + qulimit
	data, err := DB.Query(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error Query : %s", err)
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

// GetRoom functionfor get perRooms
func GetRoom(w http.ResponseWriter, r *http.Request) (*Rooms, error) {
	var room Rooms

	params := mux.Vars(r)
	rmID, _ := strconv.Atoi(params["rm_id"])

	url := strings.Split((r.URL.String()), "/")
	if rmID == 0 {
		rmID, _ = strconv.Atoi(url[5])
	}
	if rmID < 1 {
		return nil, errors.New("ID Only Positive and Integer")
	}

	qulimit := ""
	quShort := ""
	qulimit = LimitOffset(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))
	quShort = ShortBy(r.URL.Query().Get("sort_by"))
	DB.Ping()
	sql := "select rm_id,rm_name,rm_place,rm_sumpart,rm_price, rm_status FROM rooms WHERE rm_id = $1 " + quShort + qulimit
	statement, err := DB.Prepare(sql)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	err = statement.QueryRow(rmID).Scan(&room.RmID, &room.RmName, &room.RmPlace, &room.RmSumpart, &room.RmPrice, &room.RmStatus)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

// UpdateRooms function for update data rooms
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

// InsertRooms function for insert data in table Rooms
func InsertRooms(data *Rooms) (map[string]interface{}, error) {
	rmID := data.RmID
	rmName := data.RmName
	rmPlace := data.RmPlace
	rmSumpart := data.RmSumpart
	rmPrice := data.RmPrice
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	deletedAt := time.Now().Format("2006-01-02 15:04:05")
	rmStatus := "1"

	sql := fmt.Sprintf("INSERT INTO rooms (rm_id, rm_name, rm_place, rm_sumpart, rm_price, created_at, updated_at,deleted_at, rm_status) VALUES (%v, '%s', '%s', %v, %v, '%s', '%s', '%s', '%s'); ",
		rmID, rmName, rmPlace, rmSumpart, rmPrice, createdAt, updatedAt, deletedAt, rmStatus)
	_, errs := DB.Query(sql)
	if errs != nil {
		log.Println("yang error adalah insert", errs)
		return nil, errs
	}
	// defer DB.Close()
	return nil, nil
}

// DeleteRoom function for delete one data in table rooms
func DeleteRoom(w http.ResponseWriter, r *http.Request) (string, error) {
	params := mux.Vars(r)
	rmID, _ := strconv.Atoi(params["rm_id"])

	url := strings.Split((r.URL.String()), "/")
	if rmID == 0 {
		rmID, _ = strconv.Atoi(url[5])
	}
	if rmID < 1 {
		return "", errors.New("ID Only Positive and Integer")
	}
	err := CekExistData(rmID)
	if err != nil {
		return "", err
	}
	sql := "DELETE FROM rooms WHERE rm_id = $1"
	statement, err := DB.Prepare(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error Query Deleted, and %s", err)
		return "", errors.New(saveError)
	}
	statement.Exec(rmID)
	defer statement.Close()
	return "Berhasil dihapus", nil
}

// DeleteAllRoom function for Delete all data in table rooms
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

// CekExistData function for validate data exist
func CekExistData(id int) error {
	rmID := ""
	sql := "SELECT rm_id FROM rooms WHERE rm_id = $1"
	statement, err := DB.Prepare(sql)
	defer statement.Close()
	if err != nil {
		return err
	}
	err = statement.QueryRow(id).Scan(&rmID)
	if err != nil {
		return errors.New("ID Doesn't Exist")
	}
	return nil
}
