# Build stage
FROM golang:1.22.5-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/app ./api/cmd/main.go

# Final stage
FROM alpine:3.13

RUN apk add --no-cache bash

WORKDIR /app

COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/.env /app/.env 
COPY wait-for-it.sh /usr/local/bin/  
RUN chmod +x /usr/local/bin/wait-for-it.sh  # Asegúrate de que sea ejecutable

RUN echo "probando"
RUN cat /app/.env
RUN echo "termine"

RUN chmod 777 /app/.env
RUN chmod +x /app/app

EXPOSE 8080

CMD ["./app"]
