# Build stage
FROM golang:1.23.1 AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./api/cmd/main.go

# Final Stage
FROM alpine:3.13

RUN apk add --no-cache bash

WORKDIR /app

COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/api/cmd/docs /app/api/cmd/docs
COPY --from=builder /go/src/app/.env /app/.env

RUN chmod +x /app/app

EXPOSE 8081

CMD ["/app/app"]
