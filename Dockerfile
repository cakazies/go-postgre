FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/go-postgres
COPY . .

# Using go get.
RUN go get

# building apps in go-postgres
RUN go build -o go-postgres

RUN go run configs/migrate/migration.go

# running go-postgres
ENTRYPOINT ./go-postgres

# running in port
EXPOSE 8002
