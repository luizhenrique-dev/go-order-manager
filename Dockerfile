FROM golang:1.21.6-alpine as builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o order-api cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

FROM alpine:latest
WORKDIR /root/
COPY .env .env
COPY internal/infra/database/migrations ./internal/infra/database/migrations
COPY --from=builder /app/order-api .
EXPOSE 8080
CMD ["./order-api"]
