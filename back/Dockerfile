FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o api ./cmd/api
RUN go build -o migrate ./database

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .

# For the migrations
COPY --from=builder /app/database ./database
COPY --from=builder /app/migrate .

EXPOSE 8080

CMD ["./api"]