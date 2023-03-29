FROM golang:1.20-alpine 
WORKDIR /app/src/skyshi_gethired
ENV GOPATH=/app

COPY . /app/src/skyshi_gethired

RUN go build -o main

CMD ["./main"]
