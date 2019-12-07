# GO-POSTGRESQL

Golang and Database Postgresql, migrate and Dummy Data

## How to run

first => you can configuration config in (`configs/config.dev.toml`) to `configs/config.toml`
and compare your local configs
second => run this `go run application/migration/migrate.go`

dont forget for `go get`

- `go get github.com/spf13/viper`
- `go get github.com/spf13/cobra`
- `go get github.com/gorilla/mux`
- `go get github.com/lib/pq`

## Run Local

`go run main.go`
