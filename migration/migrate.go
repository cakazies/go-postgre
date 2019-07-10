package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/local/go-postgre/cmd"
	cf "github.com/local/go-postgre/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cmd.InitViper()
	var limit int
	limit = 10
	cf.Connect()
	// call function migrationrooms
	migrationRooms(limit)
	migrationUser(limit)
	cf.DB.Close()
}

func migrationRooms(limit int) {
	tableName := "rooms"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	_, err := cf.DB.Query(drop)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error Query Drop table %s, error is : %s", tableName, err))
	}
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						rm_id SERIAL NOT NULL,
						rm_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
						rm_place character varying(100) COLLATE pg_catalog."default" NOT NULL,
						rm_sumpart integer NOT NULL,
						rm_price integer NOT NULL,
						created_at timestamp without time zone NOT NULL,
						updated_at timestamp without time zone NOT NULL,
						deleted_at timestamp without time zone ,
						rm_status integer NOT NULL,
						CONSTRAINT %s_pk PRIMARY KEY (rm_id)
					);`, tableName, tableName)
	stmt, err := cf.DB.Prepare(queryCreate)
	if err != nil {
		log.Fatal("something went wrong : ", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(fmt.Sprintf("something went wrong in query create table %s error is : %s", tableName, err))
	}
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		rmName := "name-" + strconv.Itoa(i)
		rmPlace := "place-" + strconv.Itoa(i)
		rmSumpart := strconv.Itoa(1000 + i)
		rmPrice := strconv.Itoa(100000 * i)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		updatedAt := time.Now().Format("2006-01-02 15:04:05")
		deletedAt := time.Now().Format("2006-01-02 15:04:05")
		rmStatus := "1"
		sql := fmt.Sprintf("INSERT INTO %s ( rm_name, rm_place, rm_sumpart, rm_price, created_at, updated_at,deleted_at, rm_status) VALUES ('%s', '%s', %s, %s, '%s', '%s', '%s', '%s'); ",
			tableName, rmName, rmPlace, rmSumpart, rmPrice, createdAt, updatedAt, deletedAt, rmStatus)
		stmt, errs := cf.DB.Query(sql)
		if errs != nil {
			log.Fatal("yang error adalah insert errors id : ", errs)
		}
		stmt.Close()
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}

func migrationUser(limit int) {
	tableName := "users"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	_, err := cf.DB.Query(drop)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error Query Drop table %s, error is : %s", tableName, err))
	}
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						id SERIAL NOT NULL,
						email character varying(200) COLLATE pg_catalog."default" NOT NULL,
						username character varying(100) COLLATE pg_catalog."default" NOT NULL,
						password character varying(250) COLLATE pg_catalog."default" NOT NULL,
						created_at timestamp without time zone NOT NULL,
						updated_at timestamp without time zone NOT NULL,
						deleted_at timestamp without time zone,
						status integer NOT NULL,
						CONSTRAINT %s_pk PRIMARY KEY (id)
					);`, tableName, tableName)
	stmt, err := cf.DB.Prepare(queryCreate)
	if err != nil {
		log.Fatal("something went wrong : ", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(fmt.Sprintf("something went wrong in query create table %s error is : %s", tableName, err))
	}
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		email := "email_" + strconv.Itoa(i) + "@gmail.com"
		username := "username-" + strconv.Itoa(i)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		password := string(hashedPassword)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		updatedAt := time.Now().Format("2006-01-02 15:04:05")
		deletedAt := time.Now().Format("2006-01-02 15:04:05")
		status := "1"
		sql := fmt.Sprintf("INSERT INTO %s ( email, username, password, created_at, updated_at,deleted_at, status) VALUES ( '%s', '%s', '%s', '%s', '%s', '%s', '%s'); ",
			tableName, email, username, password, createdAt, updatedAt, deletedAt, status)
		stmt, errs := cf.DB.Query(sql)
		if errs != nil {
			log.Fatal("yang error adalah insert", errs)
		}
		stmt.Close()
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}
