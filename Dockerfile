FROM golang:latest AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o temperatura main.go

# FROM scratch
FROM golang:1.19-alpine
COPY --from=builder /app/temperatura /app/temperatura

CMD ["/app/temperatura"]
