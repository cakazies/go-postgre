# GO-POSTGRESQL

This repo about implementation golang connection with Database, this database using Postgresql like named, and this repo using docker, and migration table and feeder data for trying, for logging using sentry and setting configure with in `.toml`

## Collection Postman

you can import collection in
> `utils/postman/go-postgre.postman_collection.json`

## Documentation API 

You can read documentation for this repo using this [link](https://github.com/cakazies/go-postgre/wiki)

## Run Migration with

> go run application/migration/migrate.go

## How to run

- you can configuration config (**Database**) in folder `configs/config.dev.toml` to `configs/config.toml` and compare your local configs, database, host and depends

you can install dependencies using `go mod` or one by one,

- install dependencies go get
  - `go get github.com/spf13/viper`
  - `go get github.com/spf13/cobra`
  - `go get github.com/gorilla/mux`
  - `go get github.com/lib/pq`

## Run Local

`go run main.go`

## Run Docker

- build docker `docker build -t go-postgresql .`
- Run `docker run -it --rm --name cont-go-postgresql go-postgresql`