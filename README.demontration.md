Зайти на: http://localhost:8080/swagger/index.html

Создать тестовых юзеров: 
```
{
  "email": "test333eee@mail.ru",
  "login": "Fasfsfsqwert444seseses",
  "password": "HHHfTAAAfg6333"
}
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNjM3NjM3OTQ1LCJ1c2VyX3V1aWQiOiI3MzBiMzQ2MC04ZGYyLTQwYzYtOTBiMy0xMGNmMDg4MDk1MTEifQ.kOB41dM0f9jos9ptPeZUcM0DKuwu3bgw5xkkS0Szn3w

```
{
  "email": "test4444eee@mail.ru",
  "login": "Fasfetwetwtwetwrt44seseses",
  "password": "HHHfTAAAfgwetwetwet6333"
}
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNjM3NjM3OTU4LCJ1c2VyX3V1aWQiOiI5NjFkYWE5OC03NDNhLTQ0N2EtOTdjNy0wYzA0NTgyZGQ0NWEifQ.7T4qjIyOSGDojt-zX7nScS57brOxqSM48zVJNXVs7Lg

Берем одного в качестве автора -- с его токеном авторизуемся.


Создать ивент х3:
```
{
  "end": "2021-12-03 12:00",
  "roomId": 1,
  "start": "2021-12-03 11:00",
  "title": "Confa!"
}
```
```
{
  "end": "2021-12-03 12:00",
  "roomId": 3,
  "start": "2021-12-03 11:00",
  "title": "Confa!"
}
```
```
{
  "end": "2022-12-03 12:00",
  "roomId": 2,
  "start": "2022-12-03 11:00",
  "title": "Confa!"
}
```


GetEvents даст: 
```
[
  {
    "Id": 1,
    "RoomId": 1,
    "Title": "Confa!",
    "Start": "2021-12-03T12:00:00Z",
    "End": "2021-12-03T12:00:00Z",
    "AuthorId": 0
  },
  {
    "Id": 2,
    "RoomId": 2,
    "Title": "Confa!!!!",
    "Start": "2021-12-03T08:00:00Z",
    "End": "2021-12-03T12:00:00Z",
    "AuthorId": 0
  },
  {
    "Id": 3,
    "RoomId": 3,
    "Title": "Confa!!!!",
    "Start": "2021-12-04T08:00:00Z",
    "End": "2021-12-04T12:00:00Z",
    "AuthorId": 0
  }
]
```

GetEvents для комнаты 1
```
[
  {
    "Id": 1,
    "RoomId": 1,
    "Title": "Confa!",
    "Start": "2021-12-03T12:00:00Z",
    "End": "2021-12-03T12:00:00Z",
    "AuthorId": 0
  }
]
```

Создать инвайт для второго нашего пользователя: 
```
{
  "eventId": 2,
  "receiverId": 1
}
```

Запатчить его статус на аксептед:
```
{
    
}
```

Показываем препуш: 
```
nghttp -ans https://localhost/index.html
```

Вывод должен быть: 
```
...

 requestStart: the time  just before  first byte  of request  was sent
               relative  to connectEnd.   If  '*' is  shown, this  was
               pushed by server.

...

id  responseEnd requestStart  process code size request path
 13      +379us        +80us    298us  200  231 /index.html
  2    +11.90ms *     +256us  11.64ms  200  73K /img/vino.jpg
```