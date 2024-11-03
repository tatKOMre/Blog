# Используем официальный образ Go в качестве базового
FROM golang:1.22

LABEL maintaner="tatKOM"

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Собираем приложение
RUN go build -o main.go

# Открываем порт для доступа к серверу
EXPOSE 8080

# Указываем команду для запуска приложения
CMD ["cmd/main"]
