FROM golang:1.21.1-bullseye

ENV CGO_ENABLED=1

RUN apt-get update && apt-get install -y gcc build-essential

WORKDIR /home/echo

COPY . .

RUN go mod tidy

RUN go build -o echo-api cmd/main.go 