# 🎮 Game API

Простой REST API на **Gin** для управления играми и персонажами — пет-проект для изучения Go, GORM/PostgreSQL и JWT-авторизации.

## 🛠️ Стек

- **Go** + **Gin** — HTTP-сервер и роутинг
- **PostgreSQL** + **GORM** — база данных и ORM
- **golang-jwt/jwt/v5** — авторизация по токенам
- **bcrypt** — хеширование паролей
- **Docker Compose** — локальный Postgres для разработки

## ✨ Возможности

- 🕹️ CRUD для игр (`Game`)
- 🧙 CRUD для персонажей (`Character`), привязанных к игре
- 🔐 Регистрация и логин пользователей
- 🛡️ JWT-защита на создание/изменение/удаление данных

## 📁 Структура проекта

```
game-api/
├── main.go              # точка входа: конфиг, БД, миграции, роутер
├── models/               # структуры Game, Character, User
├── db/                    # подключение к PostgreSQL
├── handlers/              # HTTP-хендлеры
├── routes/                # регистрация роутов
├── middleware/             # JWT-мидлвар
├── docker-compose.yml      # Postgres для локальной разработки
└── start-db.sh              # скрипт запуска Postgres
```

## 🚀 Быстрый старт

1. Создай `.env` в корне проекта:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password_here
   DB_NAME=game_api
   JWT_SECRET=your_jwt_secret_here
   ```

2. Подними Postgres:
   ```bash
   ./start-db.sh
   ```

3. Запусти сервер:
   ```bash
   go run main.go
   ```

4. Проверь, что всё живо:
   ```bash
   curl http://localhost:8080/health
   ```

## 📡 API эндпоинты

| Метод | Путь | Доступ | Описание |
|---|---|---|---|
| POST | `/register` | 🌐 публичный | Регистрация пользователя |
| POST | `/login` | 🌐 публичный | Логин, выдаёт JWT-токен |
| GET | `/health` | 🌐 публичный | Health-check |
| GET | `/games` | 🌐 публичный | Список игр |
| GET | `/games/:id` | 🌐 публичный | Игра по ID (с персонажами) |
| GET | `/games/:id/characters` | 🌐 публичный | Персонажи игры |
| GET | `/characters/:id` | 🌐 публичный | Персонаж по ID |
| POST | `/games` | 🔒 требует токен | Создать игру |
| PUT | `/games/:id` | 🔒 требует токен | Обновить игру |
| DELETE | `/games/:id` | 🔒 требует токен | Удалить игру |
| POST | `/games/:id/characters` | 🔒 требует токен | Создать персонажа в игре |
| PUT | `/characters/:id` | 🔒 требует токен | Обновить персонажа |
| DELETE | `/characters/:id` | 🔒 требует токен | Удалить персонажа |

Для защищённых эндпоинтов нужен заголовок:
```
Authorization: Bearer <token>
```

## 🧪 Пример использования

```bash
# Регистрация
curl -X POST localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"nurmuhammed","password":"supersecret123"}'

# Логин — получаем токен
TOKEN=$(curl -s -X POST localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"nurmuhammed","password":"supersecret123"}' | jq -r .token)

# Создание игры с токеном
curl -X POST localhost:8080/games \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title":"Elden Ring"}'
```

## 🗃️ Модели данных

- **Game** — `Title`, список `Characters` (один-ко-многим)
- **Character** — `Name`, `Class`, `GameID` (внешний ключ на игру)
- **User** — `Username`, `PasswordHash` (никогда не отдаётся в JSON-ответах)
