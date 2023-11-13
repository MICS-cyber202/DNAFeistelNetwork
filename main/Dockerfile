# syntax=docker/dockerfile:1.6
FROM golang:1.16-alpine
WORKDIR /dnaxor

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /dnaxor
CMD [./main.go]