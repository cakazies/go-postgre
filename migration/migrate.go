package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/local/go-postgre/cmd"
	cf "github.com/local/go-postgre/models"
)

func main() {
	cmd.InitViper()
	cf.Connect()
	stmt, err := cf.DB.Prepare(`
							CREATE TABLE public.rooms
							(
								rm_id integer NOT NULL,
								rm_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
								rm_place character varying(100) COLLATE pg_catalog."default" NOT NULL,
								rm_sumpart integer NOT NULL,
								rm_price integer NOT NULL,
								created_at timestamp without time zone NOT NULL,
								updated_at timestamp without time zone NOT NULL,
								deleted_at timestamp without time zone NOT NULL,
								rm_status integer NOT NULL,
								CONSTRAINT rooms_pk PRIMARY KEY (rm_id)
							);`)
	if err != nil {
		log.Println("something went wrong : ", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("something went wrong : ", err)
	}
	log.Println("Import Table Room Succesfull")

	for i := 1; i <= 10; i++ {
		rmID := strconv.Itoa(i)
		rmName := "name-" + strconv.Itoa(i)
		rmPlace := "place-" + strconv.Itoa(i)
		rmSumpart := strconv.Itoa(1000 + i)
		rmPrice := strconv.Itoa(100000 * i)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		updatedAt := time.Now().Format("2006-01-02 15:04:05")
		deletedAt := time.Now().Format("2006-01-02 15:04:05")
		rmStatus := "1"

		sql := fmt.Sprintf("INSERT INTO rooms (rm_id, rm_name, rm_place, rm_sumpart, rm_price, created_at, updated_at,deleted_at, rm_status) VALUES (%s, '%s', '%s', %s, %s, '%s', '%s', '%s', '%s'); ",
			rmID, rmName, rmPlace, rmSumpart, rmPrice, createdAt, updatedAt, deletedAt, rmStatus)
		_, errs := cf.DB.Query(sql)
		if errs != nil {
			log.Println("yang error adalah insert", errs)
		}
	}

	cf.DB.Close()
}
