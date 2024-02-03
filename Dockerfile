FROM golang:1.21.1-bullseye

RUN go install github.com/cosmtrek/air@latest

