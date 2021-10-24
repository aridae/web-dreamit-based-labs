# DreamIt API

Если тут ошибки где-то, тыкните меня, я не шарю, я просто позер :)

Спасибо Ивану Коваленко!

## Команды для докера 

Это шпора, а не пайплайн -- не надо их все запускать)))

Cобрать контейнеры: 
```bash
docker-compose up -d --force-recreate --remove-orphans
```

Посмотреть логи 
```bash
docker-compose logs -f
```

Зайти в бд, глянуть что там как
```bash 
psql -h localhost -p 5432 -d dreamit_api_db -U postgres
```

Вывод примерно такой должен быть:
```bash
dreamit_api_db=# \dt
             List of relations
 Schema |     Name      | Type  |  Owner   
--------+---------------+-------+----------
 public | auth_services | table | postgres
 public | auth_tokens   | table | postgres
 public | calendar      | table | postgres
 public | intervals     | table | postgres
 public | rooms         | table | postgres
 public | schedules     | table | postgres
 public | users         | table | postgres
```

Почистить вольюмы по имени: 
```bash
docker rm -fv <container-name>
```

В частности, может понадобиться почистить постгру: 
```bash
docker rm -fv postgresdreamit
```

Уничтожить и подчистить все контейнеры на машине:
```bash
make armageddon
```

Глянуть все контейнеры на машине: 
```bash
docker ps --all
```

Или вот так: 
```bash
docker container ls --all
```

## Как дергать апи курлом 

## Вики для разбора теории по лабе

Желаю всем создателям этого курса получасового сна и кофеиновой интоксикации :)


| Таска   |      Лаба      |  Статус | Теория к этой параше |
|----------|:-------------:|:-------------:|:-------------:|
| Спроектировать в формате Swagger (https://editor.swagger.io/) внешнее публичное API системы в идеологии REST. Предусмотреть как минимум один вызов на базе метода PATCH. |  1 | In progress | WIP | 
| По спроектированному swagger подготовить реализацию в программном коде. К реализации так же подключить swagger, уже для документирования | 1 | In progress | WIP | 
| Настроить Nginx для работы web-приложения в части маршрутизации | 1 | In progress | WIP | 
| Настроить Nginx в части балансировки | 1 | In progress | WIP | 
| Настроить Nginx таким образом, чтобы подменялось имя сервера в заголовках http-ответов | 1 | In progress | WIP | 
| Настроить кеширование (для всех GET-запросов, кроме /api) и gzip-сжатие в Nginx | 1 | In progress | WIP | 
