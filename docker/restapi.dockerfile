FROM golang:1.20-alpine as builder
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .
EXPOSE 3030
CMD ["/main"]