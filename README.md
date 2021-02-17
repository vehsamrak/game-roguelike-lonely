# game-roguelike-lonely

### Локальное окружение
Приложение работает в docker контейнерах. Для запуска необходимо: 

1. Пробросить порты контейнеров можно перенеся файл `infrastructure/docker-compose.override.yml` в корень проекта.
1. Создать `.env` в корне проекта, создав его из `.env.local` и переопределив параметры доступа к БД.
1. Запуск контейнеров, из корня проекта: `docker-compose up -d`

### Обновление зависимостей
Обновление зависимостей осуществляется через composer внутри контейнера: `docker exec game_roguelike_lonely_php composer update`.

### Миграции
Запуск миграций: `docker exec game_roguelike_lonely_php bin/console doc:mig:mig -n`

### Тесты
Запуск юнит тестов с локальным интерпретатором: `bin/phpunit`

Запуск внутри контейнера: `docker exec game_roguelike_lonely_php bin/phpunit`
