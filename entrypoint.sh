#!/bin/sh

set -e

# Создаем .env файл из переменных окружения
cat > .env <<EOL
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=qwerty
DB_NAME=postgres
APP_PORT=8080
EOL

# Формируем строку подключения к БД
CONN_STR="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Применяем миграции
./wait-for-db.sh db goose -dir ./migrations postgres "$CONN_STR" up

# Запускаем приложение
exec ./marketplace