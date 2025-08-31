# Используем официальный минимальный образ Go для сборки
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app main.go

# Используем минимальный образ для запуска
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
ENV PORT=8080
CMD ["./app"]
