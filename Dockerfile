FROM golang:1.17-buster

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

RUN go install github.com/cosmtrek/air@v1.27.3 \
  && go install github.com/99designs/gqlgen@v0.13.0

COPY . ./
