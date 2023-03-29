FROM golang:1.20-alpine AS build

ENV MYSQL_HOST=localhost \
	MYSQL_PORT=3306 \
	MYSQL_USER=root \
	MYSQL_PASSWORD=root \
	MYSQL_DBNAME=skyshi_gethired
	
WORKDIR /go/src/skyshi/gethired.go

COPY . .

RUN go build -o main

CMD ["./main"]