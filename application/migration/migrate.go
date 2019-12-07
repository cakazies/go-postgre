package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	cf "github.com/cakazies/go-postgre/application/models"
	"github.com/cakazies/go-postgre/cmd"
	"github.com/cakazies/go-postgre/utils"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cmd.InitViper()
	var limit int
	limit = 10
	cf.Connect()
	// call function migrationrooms
	MigrationRooms(limit)
	MigrationUser(limit)
	MigrationBorrow(limit)
	cf.DB.Close()
}

// MigrationRooms function for migrations table rooms
func MigrationRooms(limit int) {
	tableName := "rooms"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	_, err := cf.DB.Query(drop)
	utils.FailError(err, fmt.Sprintf("Error Query Drop table %s", tableName))
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						rm_id SERIAL NOT NULL,
						rm_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
						rm_place character varying(100) COLLATE pg_catalog."default" NOT NULL,
						rm_sumpart integer NOT NULL,
						rm_price integer NOT NULL,
						created_at timestamp without time zone NOT NULL,
						updated_at timestamp without time zone ,
						deleted_at timestamp without time zone ,
						rm_status integer NOT NULL,
						CONSTRAINT %s_pk PRIMARY KEY (rm_id)
					);`, tableName, tableName)
	stmt, err := cf.DB.Prepare(queryCreate)
	utils.FailError(err, fmt.Sprintf("Error Create Table %s ", tableName))
	_, err = stmt.Exec()
	utils.FailError(err, fmt.Sprintf("Error Create Table %s ", tableName))
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		rmName := "name-" + strconv.Itoa(i)
		rmPlace := "place-" + strconv.Itoa(i)
		rmSumpart := strconv.Itoa(1000 + i)
		rmPrice := strconv.Itoa(100000 * i)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		rmStatus := "1"
		sql := fmt.Sprintf("INSERT INTO %s ( rm_name, rm_place, rm_sumpart, rm_price, created_at, rm_status) VALUES ('%s', '%s', %s, %s, '%s', '%s'); ",
			tableName, rmName, rmPlace, rmSumpart, rmPrice, createdAt, rmStatus)
		stmt, err := cf.DB.Query(sql)
		utils.FailError(err, fmt.Sprintf("Error Insert Data Table %s ", tableName))
		stmt.Close()
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}

func MigrationUser(limit int) {
	tableName := "users"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	_, err := cf.DB.Query(drop)
	utils.FailError(err, fmt.Sprintf("Error Query Drop table %s ", tableName))
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						id SERIAL NOT NULL,
						email character varying(200) COLLATE pg_catalog."default" NOT NULL,
						username character varying(100) COLLATE pg_catalog."default" NOT NULL,
						password character varying(250) COLLATE pg_catalog."default" NOT NULL,
						created_at timestamp without time zone NOT NULL,
						updated_at timestamp without time zone ,
						deleted_at timestamp without time zone,
						status integer NOT NULL,
						CONSTRAINT %s_pk PRIMARY KEY (id)
					);`, tableName, tableName)
	stmt, err := cf.DB.Prepare(queryCreate)
	utils.FailError(err, fmt.Sprintf("Error Create table %s ", tableName))
	_, err = stmt.Exec()
	utils.FailError(err, fmt.Sprintf("Error Exec Create table %s ", tableName))
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		email := "email_" + strconv.Itoa(i) + "@gmail.com"
		username := "username-" + strconv.Itoa(i)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		password := string(hashedPassword)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		status := "1"
		sql := fmt.Sprintf("INSERT INTO %s ( email, username, password, created_at, status) VALUES ( '%s', '%s', '%s', '%s', '%s'); ",
			tableName, email, username, password, createdAt, status)
		stmt, err := cf.DB.Query(sql)
		utils.FailError(err, fmt.Sprintf("Error Insert Data table %s ", tableName))
		stmt.Close()
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}

func MigrationBorrow(limit int) {
	tableName := "borrow"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	_, err := cf.DB.Query(drop)
	utils.FailError(err, fmt.Sprintf("Error Query Drop table %s ", tableName))
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						id SERIAL NOT NULL,
						room_id integer NOT NULL,
						event_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
						borrower character varying(100) COLLATE pg_catalog."default" NOT NULL,
						start_date timestamp without time zone NOT NULL,
						end_date timestamp without time zone NOT NULL,
						created_at timestamp without time zone NOT NULL,
						updated_at timestamp without time zone,
						deleted_at timestamp without time zone,
						CONSTRAINT %s_pk PRIMARY KEY (id)
					);`, tableName, tableName)
	stmt, err := cf.DB.Prepare(queryCreate)
	utils.FailError(err, fmt.Sprintf("Error Create table %s ", tableName))
	_, err = stmt.Exec()
	utils.FailError(err, fmt.Sprintf("Error Exec Create table %s ", tableName))
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		room_id := rand.Intn(10)
		event_name := "name-event-" + strconv.Itoa(i)
		borrower := "borrower-" + strconv.Itoa(i)
		start_date := time.Now().Format("2006-01-02 15:04:05")
		end_date := time.Now().Format("2006-01-02 15:04:05")
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		sql := fmt.Sprintf("INSERT INTO %s ( room_id, event_name, borrower,start_date,end_date,  created_at) VALUES ( %v, '%s', '%s', '%s', '%s', '%s'); ",
			tableName, room_id, event_name, borrower, start_date, end_date, createdAt)
		stmt, err := cf.DB.Query(sql)
		utils.FailError(err, fmt.Sprintf("Error Insert Data table %s ", tableName))
		stmt.Close()
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}
