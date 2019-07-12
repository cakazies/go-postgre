package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	DB *sql.DB
)

func Connect() {
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("Check your config file, Database not connect, Error is :", err)
		panic(err)
	}
	log.Println("koneksi sukses lek")
	// defer result.Close()
	DB = result

}
