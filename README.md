# Маркетплейс
### Стек: Go, Gin, GORM, Postgres
### Для запуска нужно клонировать репозиторий и прописать команды:

docker build -t vk-test:v1 .

docker compose up

### Регистрация пользователя: POST /sign_up

В теле запроса нужно указать логин и пароль:

```json
{
    "login": "login",
    "password": "Abc123456"
}
```

### Авторизация пользователя: POST /sign_in
```json
{
    "login": "login",
    "password": "Abc123456"
}
```
Токен пользователя есть в хедере ответа. Его нужно добавлять в хедер запроса при размещении или отображении объявления.

### Размещение объявления: POST /post_ad
```json
{
    "title": "title",
    "description": "description",
    "image": "image.png",
    "price": 1
}
```

### Отображение ленты объявлений: GET /get_ads

- Для выбора страницы нужно указать параметр page: /get_ads?page=1

- Для выбора критерия сортировки(date или price) нужно указать параметр sort: /get_ads?sort=price или /get_ads?sort=date

- Для сортировки по убыванию нужно указать параметр order: get_ads?sort=price&order=desc.

Если объявление принадлежит пользователю, то в теле ответа есть параметр IsYours со значением true.

Пример ответа:
```json
[
    {
        "ID": 1,
        "CreatedAt": "2025-07-21T19:15:20.564471Z",
        "UpdatedAt": "2025-07-21T19:15:20.564471Z",
        "DeletedAt": null,
        "Title": "title",
        "Description": "description",
        "Image": "image.png",
        "Price": 1,
        "is_yours": true
    }
]
```
