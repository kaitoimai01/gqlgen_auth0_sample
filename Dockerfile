FROM golang:1.17-buster

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

# TODO: gqlgen
# && go install github.com/99designs/gqlgen@v0.13.0
RUN go install github.com/cosmtrek/air@v1.27.3

COPY . ./
