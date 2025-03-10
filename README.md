# Web-calculus

Веб-сервис **распределенного вычисления арифметических выражений**. Операции *сложения и умножения, деление и вычитания* выполняются *"очень-очень"* долго, их вычисление в *"альтернативной реальности"* занимает *"гигантские"* вычислительные мощности. Соответственно, **каждое действие** нужно уметь **выполнять отдельно**, и масштабировать систему можем добавлением вычислительных мощностей в виде новых *"машин"*. Поэтому **пользователь** может с какой-то периодичностью **уточнять** у сервера *"не посчиталось ли выражение"*? Если **выражение** наконец будет **вычислено** - то он **получит результат**.

## **Backend часть** 
состоит из 2 элементов:
* Оркестратора - сервер, который принимает арифметическое выражение, переводит его в набор последовательных задач и обеспечивает порядок их выполнения.
* Агента - вычислитель, который может получить от оркестратора задачу, выполнить ее и вернуть серверу результат.

Также, весь сервис покрыт тестами.

```mermaid
flowchart LR
    A[Клиент] -->|http/отправляет выражение(я)| Б[Оркестратор]
    Б -->|http/отдает задачки| В[Агент(ы)]
    В -->|http/отправляет(ют) ответ(ы)| Б
    Б -->|http/отправляет статус выполнения и результат| A
```

## Запуск
### оркестратора:
```
go run ./cmd/orchestrator/main.go
```
default 
* `PORT=8080`
* `TIME_ADDITION_MS=1000`
* `TIME_SUBTRACTION_MS=1000`
* `TIME_MULTIPLICATIONS_MS=2000`
* `TIME_DIVISIONS_MS=2500`

или
```
export PORT=your_port \
TIME_ADDITION_MS=your_time_add \
TIME_SUBTRACTION_MS=your_time_sub \
TIME_MULTIPLICATIONS_MS=your_time_mul \
TIME_DIVISIONS_MS=your_time_div \
&& go run ./cmd/orchestrator/main.go
```
### агента:
```
go run ./cmd/agent/main.go
```
default 
* `PORT=8080`
* `COMPUTING_POWER=3`

или
```
export PORT=listen_port \
COMPUTING_POWER=your_power \
&& go run ./cmd/agent/main.go
```

## Использование

### POST /api/v1/calculate
добавление арифметического выражения
#### 201 - выражение принято, ответ: уникальный id выражения 
```
curl --location 'localhost:PORT/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1.1 + 2.2*(8.8 - 3.3)/5"
}'
```
#### 422 - невалидные данные, ответ: сообщение с ошибкой
```
curl --location 'localhost:PORT/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "9.03 + (1.23 -"
}'
```
#### 500 - неправильный запрос, ответ: invalid request
```
curl --location 'localhost:PORT/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data "{
  "pedro": "pe"
}"
```

### GET /api/v1/expressions
получение списка выражений
```
curl --location 'localhost:PORT/api/v1/expressions'
```
* 200 - список получен, ответ: список выражений с их id, статусами и результатами 
* 500 - список пустой, ответ: empty list expressions

### GET /api/v1/expressions/:id
получение выражения по его идентификатору
```
curl --location 'localhost:PORT/api/v1/expressions/:id'
```
* 200 - выражение получено, ответ: выражение с его id, статусом и результатом
* 404 - нет такого выражения, ответ: expression does not exist 
* 500 - некорректный id, ответ: invalid id

### GET /internal/task
получение задачи для выполнения
```
curl --location 'localhost:PORT/internal/task'
```
* 200 - задача получена, ответ: задача с операцией над числами
* 404 - нет такой задачи, ответ: no tasks available 
* 500 - что-то пошло не так

### POST /internal/task
прием результата выполнения задачи
#### 200 - результат записан, ответ: success 
```
curl --location 'localhost:PORT/internal/task' \
--header 'Content-Type: application/json' \
--data '{
  "id": 1,
  "result": 2.5
}'
```
#### 404 - задача не найдена, ответ: task not found
```
curl --location 'localhost:PORT/internal/task' \
--header 'Content-Type: application/json' \
--data '{
  "id": 123456,
  "result": 2.5
}'
```
#### 422 - невалидные данные, ответ: invalid request
```
curl --location 'localhost:PORT/internal/task' \
--header 'Content-Type: application/json' \
--data '{
  "id": "1",
  "result": "2.5"
}'
```
