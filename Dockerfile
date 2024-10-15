# Указываем базовый образ
FROM golang:1.23.1 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./


# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в контейнер
COPY . .


# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp cmd/main.go

# Используем минималистичный образ для запуска приложения
FROM alpine:latest

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/myapp /myapp

# Указываем команду, которая будет выполнена при запуске контейнера
CMD ["/myapp"]
