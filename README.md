# TODO List
## ЧТО СДЕЛАНО
- Из первого все, кроме 3 и 6
- Из второго все
## Развёртывание в Kubernetes

Для развёртывания приложения в Kubernetes выполните команду:
```sh
kubectl apply -f kuber/
```

## Логирование и Метрики

- Логи отправляются напрямую в **Loki**.
- Метрики собираются **Prometheus** каждые 5 секунд.
- Для отображения используется **Grafana**, развернутая в Kubernetes.

### Доступ к Grafana, Loki и Prometheus и приложению

- получение ссылки на Grafana:
  ```sh
  minikube service grafana --url
  ```
- Адрес **Prometheus**:
  ```
  http://prometheus.default.svc.cluster.local:9090
  ```
- Адрес **Loki**:
  ```
  http://loki.default.svc.cluster.local:3100
  ```
- Получение ссылки на **app**:
  ```sh
  minikube service app --url
  ```

## API

### Регистрация пользователя

#### Запрос:
```http
POST /register
```
```json
{
  "username": "joa",
  "password": "securepassword123",
  "first_name": "John",
  "last_name": "Doe"
}
```

### Работа с задачами (Tasks)

Для работы с API задач требуется JWT-токен.

#### Получение задачи по ID
```http
GET /task/:taskId
```

#### Создание задачи
```http
POST /task
```
##### Пример тела запроса:
```json
{
  "expiredAt": "2025-03-21T12:00:00Z",
  "name": "Название задачи",
  "description": "Описание задачи"
}
```

#### Удаление задачи
```http
DELETE /task/:taskId
```

