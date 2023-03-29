FROM golang:1.20-alpine as build
WORKDIR /usr/src/app

RUN ls -la
COPY . .

RUN go build -o main

CMD ["./main"]
