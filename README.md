# GO-POSTGRESQL

Golang and Database Postgresql, migrate and Dummy Data

## Collection Postman

you can import collection in
> `utils/postman/go-postgre.postman_collection.json`

## Run Migration with

> go run application/migration/migrate.go

## How to run

- you can configuration config (**Database**) in (`env.dev`) to `.env` and compare your local configs

- install dependencies go get
  - `go get github.com/spf13/viper`
  - `go get github.com/spf13/cobra`
  - `go get github.com/gorilla/mux`
  - `go get github.com/lib/pq`

## Run Local

`go run main.go`
