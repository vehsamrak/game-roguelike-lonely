# game-roguelike-lonely

### Локальное окружение
Приложение работает в docker контейнерах. Для запуска необходимо: 

1. Пробросить порты контейнеров можно перенеся файл `infrastructure/docker-compose.override.yml` в корень проекта.
1. Создать `.env.local` в корне проекта, создав его из `.env` и переопределив параметры доступа к БД.
1. Запуск контейнеров, из корня проекта: `docker-compose up -d`

### Обновление зависимостей
`docker-compose exec php composer update`.

### Миграции
`docker-compose exec php bin/console doc:mig:mig -n`

### Тесты
`docker-compose exec php bin/phpunit`
