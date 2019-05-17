package main

import (
	"log"

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
	cf.DB.Close()
}
