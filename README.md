## Домашнее задание

Требуется разработать создание и просмотр анкет в социальной сети.

**Функциональные требования**:
1) Простейшая авторизация пользователя.
2) Возможность создания пользователя, где указывается следующая информация:
    - Имя
    - Фамилия
    - Возраст
    - Пол
    - Интересы
    - Город
3) Страницы с анкетой

Есть возможность авторизации, регистрации, получение анкет по ID.
<br>Отсутствуют SQL-инъекции.
<br>Пароль хранится безопасно.

### Инструкция по локальному запуску

1) Склонировать репозиторий
2) Внутри репозитория выполнить команду
   <br>``docker-compose up``
   
3) [Swagger](http://127.0.0.1:5050/swagger/index.html)
4) [Postman-коллекция](https://github.com/AntonOcean/highload-architect/blob/970ec4692831e9e9d9abfa2e08683dea5b06925f/backend/docs/Backend%20swagger.postman_collection.json)

### Нагрузочное тестирование

[Отчет](https://github.com/AntonOcean/highload-architect/blob/e1f1b57b659ac771825fc25b2a18e71310dd0964/backend/docs/highload-report-v2.pdf)

### Добавление репликации

[Отчет](https://github.com/AntonOcean/highload-architect/blob/e1f1b57b659ac771825fc25b2a18e71310dd0964/backend/docs/replica-report.pdf)

### Добавление кеша для ленты новостей

[Отчет](https://github.com/AntonOcean/highload-architect/blob/cfe9084b72d4c5f1045a94ef62b12bf2ebfa24e7/backend/docs/cache-report.pdf)

### Создание чатов с шардированием
[Отчет](https://github.com/AntonOcean/highload-architect/blob/269ca8cb67519c38fb63b175fc0b9347dd85adb3/chat/docs/sharding-report.pdf)
<br>``docker-compose -f docker-compose.chat.yaml up --scale pg-worker=5``