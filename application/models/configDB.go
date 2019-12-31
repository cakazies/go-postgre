package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cakazies/go-postgre/utils"
	// library for conenct postgresql
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	// DB variable for connection DB postgresql
	DB *sql.DB
)

// Connect function for checking connection to postgresql
func Connect() {
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	utils.FailError(err, "Check your config file, Database not connect")
	// defer result.Close()
	err = result.Ping()
	if err != nil {
		log.Println("Error DB Ping : ", err)
		return
	}

	DB = result
}
