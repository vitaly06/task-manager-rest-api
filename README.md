# Task Manager REST API
REST API для управления задачами с регистрацией пользователей, аутентификацией и разграничением доступа. Проект разработан на GO с использованием Fiber, GORM и PostgreSQL.
## Возможности
- Регистрация и авторизация пользователей
- Аутентификация с использованием JWT
- Хранение токена в cookie
- Создание, получение, обновление и удаление задач
- Разграничение доступа: каждый пользователь управляет только своими задачами
- Автоматические создание и обновление структур с помощью GORM AutoMigrate
- Запуск приложения через Docker Compose
## Технологии
- **GO** - язык программирования
- **Fiber v3** - веб-фреймворк
- **GORM** - ORM для работы с базой данных
- **PostgreSQL** - база данных
- **JWT** - механизм аутентификации
- **Docker** и **Docker Compose** - контейнеризация приложения
## Требования для запуска
Для запуска проекта необходимы:
- **Docker**
- **Docker Compose**
## Установка и запуск
### 1. Клонирование репозитория
```
git clone https://github.com/vitaly06/task-manager-rest-api
cd task-manager-rest-api
```
### 2. Настройка окружения
Создайте файл .env в корне проекта:
```
PORT=3000

DSN="host=localhost user=postgres password=your_password dbname=name_db port=5432"
JWT_SECRET="YOUR_SUPER_SECRET_JWT"

POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_password
POSTGRES_DB=name_db
```
Замените значения на собственные.
### 3. Запуск приложения
Соберите образ и запустите контейнеры:
```
docker compose up --build
```
После запуска API будет доступно по адресу:
```
http://localhost:3000
```
## API
Ниже представленые маршруты.
### Аутентификация
| Метод | Маршрут | Описание | Доступ | 
| --- | --- | --- | --- |
| POST | /auth/sign-up | Регистрация пользователя | Публичный |
| POST | /auth/sign-in | Авторизация пользователя | Публичный |
| POST | /auth/logout | Выход из аккаунта | Требуется авторизация |
### Управление задачами
| Метод | Маршрут | Описание | Доступ | 
| --- | --- | --- | --- |
| POST | /tasks | Создание задачи | Требуется авторизация |
| GET | /tasks | Получение задач пользователя | Требуется авторизация |
| GET | /tasks/:id | Получение задачи по ID | Требуется авторизация |
| PUT | /tasks/:id | Обновление задачи | Требуется авторизация |
| DELETE | /tasks/:id | Удаление задачи | Требуется авторизация |
## Структура проекта
```
.
├── cmd
│   └── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── consts
│   │   └── color.go
│   ├── database
│   │   └── postgres.go
│   ├── domain
│   │   ├── task.go
│   │   └── user.go
│   ├── handler
│   │   ├── auth_dto.go
│   │   ├── auth_handler.go
│   │   ├── task_dto.go
│   │   └── task_handler.go
│   ├── middleware
│   │   └── auth.go
│   ├── repository
│   │   ├── task_repository.go
│   │   └── user_repository.go
│   ├── routes
│   │   ├── auth_routes.go
│   │   └── task_routes.go
│   └── service
│       ├── auth_service.go
│       ├── jwt.go
│       └── task_service.go
└── README.md
```
## Безопасность
- Пароли хранятся в хэшированом виде
- Защищённые маршруты проверют JWT
- Пользователь не имеет доступ к задачам других пользователей
- Секреты хранятся в переменной окружения