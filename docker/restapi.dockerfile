FROM golang:1.20-alpine 
WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest
