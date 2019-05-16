### Campaign Microservice

How to run
#### FYI
the apps have two configuration type, production and development.
default apps using development configuration (`configs/config.dev.toml`).
for production set `CAMPAIGN_ENV` as environment with value `PRODUCTION`.

#### Local
`go run main.go`

or

`env CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o campaign-microservice -v .`

`./campaign-microservice`

#### Docker
Build to image docker build -t campaign-microservice:v.01 .

Run Images :

`docker run -d --name campaign-microservice -p 8181:8181 campaign-microservice:v.01`
