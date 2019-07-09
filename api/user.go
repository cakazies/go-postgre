package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/local/go-postgre/models"
)

func Register(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, err
	}

	response := user.CreateAccount()
	return response, err
}

func Login(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	// data, err := models.GetRooms(w, r)
	log.Println("sapiiii login")
	// if err != nil {
	// 	return nil, err
	// }

	return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, nil
	///////////////////////////////////

	// account := &models.Account{}

	// err := json.NewDecoder(r.Body).Decode(account)
	// if err != nil {
	// 	fmt.Println(err)
	// 	util.Respond(w, util.MetaMsg(false, "Invalid request"))
	// 	return
	// }

	// response := account.CreateAccount()
	// util.Respond(w, response)
}
