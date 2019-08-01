package models

import (
	"net/http"
)

type Borrow struct {
	Id        int    `json:"id,omitempty"`
	RoomID    int    `json:"room_id,omitempty"`
	EventName string `json:"event_name,omitempty"`
	Borrower  string `json:"borrower,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

func GetBorrows(w http.ResponseWriter, r *http.Request) (ManyRooms, error) {

	return nil, nil
}
