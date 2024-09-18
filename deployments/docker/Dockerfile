FROM golang:1.22

LABEL maintaner="tatKOM"

WORKDIR app/

COPY . .

RUN go mod download

COPY go.mod go.sum ./

RUN go build -o server cmd/main.go
