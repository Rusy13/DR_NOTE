# Примеры запросов <a name="examples"></a>


* [Создание пользователя:POST http://localhost:8000/users c телом:]
```
{
    "name": "Bil Bil",
    "email": "bil.bil@example.com",
    "birthday": "1975-06-05"
}
```


* [Получение пользователя по ID: GET http://localhost:8000/users/3]


* [Получение всех пользователей: GET http://localhost:8000/users]


* [Получение подписок по ID: GET http://localhost:8000/users/3/subscribers]


* [Подписка на пользователя: GET http://localhost:8000/users/8/subscribe/2]


* [Отписка от пользователя: GET http://localhost:8000/users/5/unsubscribe/2]




```
Процесс запуска: 

docker-compose up --build
прогон миграций: make migration-up
go run main.go
```

Также создан телеграм бот