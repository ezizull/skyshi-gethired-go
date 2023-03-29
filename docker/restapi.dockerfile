FROM golang:1.20-alpine
WORKDIR /usr/src/app

ENV GOPATH=/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy and build the app
COPY . .

RUN go build -o main .
CMD [ "./main" ]