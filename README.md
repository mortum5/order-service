<p align="center">
  <img src="https://socialify.git.ci/mortum5/order-service/image?description=1&descriptionEditable=&font=Inter&issues=1&language=1&name=1&owner=1&pattern=Signal&pulls=1&stargazers=1&theme=Light"     alt="order-service" width="640" height="320" />
</p>

![Repository Top Language](https://img.shields.io/github/languages/top/mortum5/order-service)
![Github Open Issues](https://img.shields.io/github/issues/mortum5/order-service)
![GitHub contributors](https://img.shields.io/github/contributors/mortum5/order-service)

# Order Service

## About 

Тестовое задание которое включае в себя следующие задачи:
- Написание сервиса по получению заказов с помощью **Nats Streaming** и сохранению их в **Postgres БД**
- Добавление кеш системы для ускорения доступа к данным
- Создание **HTTP** сервера способного выдавать заказы по его уникальному идентификатору

## Solution notes

- Кеш использует библиотеку `go-cache`, сохраняет каждый новый заказ и восстанавливает после падения сервиса последние 100 заказов основывая на обратном лексиграфическом порядке их id
- В качестве интерфейса реализован swagger с помощью библиотеки `swag`

## Getting started

#### Before Run

Install packages:
```sh
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/swaggo/swag/cmd/swag@latest  
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Copy config:
```sh
$> cp config/env.example config/.env
```

А также 
Доступные команды:
```sh
$> make generate     # сгененрировать swagger и sqlc код
$> make init         # запуск проекта включая проведение миграций
$> make start        # поднятие контейнеров
$> COUNT=5 make send # генерация и отправка тестовых заказов
$> make stop         # остановка контейнеров
```
