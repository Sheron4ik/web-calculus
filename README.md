# Web-calculus

Веб-сервис **подсчета арифметического выражения** с одной `POST` ручкой, расположенной по `url` - `/api/v1/calculate`. В тело `POST` запроса **необходимо** передать поле `expression` с веденным пользователем выражением. В случае **успешной обработки запроса** пользователь получит `http status 200` и результат посчитанного выражения. В **ином случае** - пользователь может получить `http status 422` с **описанием ошибки, которая произошла** *(например, пользователь забыл передать тело запроса с полем `expression` или ввел некорректное выражение)*, или `http status 500` с ошибкой, что что-то пошло не так.
Также, весь сервис покрыт тестами.

## Запуск

```
go run ./cmd/main.go
```
default `PORT=8080` или
```
export PORT=your_port && go run ./cmd/main.go
```

## Использование

```
curl --location 'localhost:PORT/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1+2*(8-3)/5"
}'
```
