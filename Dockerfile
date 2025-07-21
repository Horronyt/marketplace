# Сборка приложения
FROM golang:1.24-alpine AS builder

# Установка goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app
COPY . .

# Сборка приложения
RUN go mod download
RUN go build -o marketplace ./cmd/main.go

# Финальный образ
FROM alpine:latest

RUN apk add --no-cache bash postgresql-client

WORKDIR /app

# Копируем бинарник приложения и goose
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/marketplace .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/wait-for-db.sh .
COPY --from=builder /app/entrypoint.sh .
COPY --from=builder /app/configs ./configs

# Делаем скрипты исполняемыми
RUN chmod +x wait-for-db.sh entrypoint.sh marketplace

# Переменные окружения для миграций
ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=qwerty
ENV DB_NAME=postgres

# Порт приложения
EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]