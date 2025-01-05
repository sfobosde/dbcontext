# Этап 1: Устанавливаем зависимости и компилируем приложение
FROM golang:1.20.5 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go test -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]

# CMD ["sh", "-c", "if [ $? -eq 0 ]; then echo 'Успешно'; else echo 'Провалено'; fi"]
