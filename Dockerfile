FROM golang:1.17.0-alpine3.14

ENV ROOT=/go/src
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY ./src/main.go ${ROOT}

COPY ./src/go.mod ${ROOT}

RUN go mod tidy