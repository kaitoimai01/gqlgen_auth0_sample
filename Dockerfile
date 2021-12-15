FROM golang:1.17-buster

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

RUN go install github.com/cosmtrek/air@v1.27.3

COPY . ./
