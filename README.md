# GO Exam — API Documentation

Бэкенд на Go (Gin + GORM + PostgreSQL). REST API + WebSocket.

## Содержание

- [Запуск](#запуск)
- [Переменные окружения](#переменные-окружения)
- [Авторизация](#авторизация)
- [Auth](#auth)
- [Users](#users)
- [Posts](#posts)
- [Comments](#comments)
- [Likes](#likes)
- [Communities](#communities)
- [Schedule](#schedule)
- [Chats](#chats)
- [Messages](#messages)
- [Direct Messages](#direct-messages)
- [WebSocket](#websocket)
- [Роли и доступ](#роли-и-доступ)

---

## Запуск

```bash
git clone <repo>
cd GO_Exam
cp .env.example .env   
go mod download
go run cmd/main.go
```

Сервер запускается на `http://localhost:8080`

---

## Переменные окружения

Файл `.env` в корне проекта:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=go_exam
JWT_SECRET=your_jwt_secret_key
```

---

## Авторизация

Все маршруты кроме `/auth/*` требуют заголовок:

```
Authorization: Bearer <token>
```

Токен получается при логине (`POST /auth/login`).

---

## Auth

### Регистрация

```
POST /auth/register
```

**Body:**
```json
{
  "email": "user@example.com",
  "password": "secret123",
  "first_name": "Иван",
  "last_name": "Иванов",
  "group": "IT-21"
}
```

**Response `201`:**
```json
{
  "id": 1,
  "email": "user@example.com",
  "first_name": "Иван",
  "last_name": "Иванов",
  "group": "IT-21",
  "role": "user",
  "avatar": ""
}
```

---

### Вход

```
POST /auth/login
```

**Body:**
```json
{
  "email": "user@example.com",
  "password": "secret123"
}
```

**Response `200`:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

## Users

### Текущий пользователь

```
GET /api/me
```

🔒 Требует токен

**Response `200`:**
```json
{
  "id": 1,
  "email": "user@example.com",
  "first_name": "Иван",
  "last_name": "Иванов",
  "group": "IT-21",
  "role": "user",
  "avatar": "/uploads/avatars/1_1778506904.jpg"
}
```

---

### Загрузить аватар

```
POST /api/avatar
```

🔒 Требует токен

**Body:** `multipart/form-data`

| Поле | Тип | Описание |
|------|-----|----------|
| `avatar` | file | JPG / JPEG / PNG, макс. 5 МБ |

**Response `200`:**
```json
{
  "avatar": "/uploads/avatars/1_1778506904.jpg"
}
```

---

### Изменить роль пользователя

```
POST /api/admin/users/:id/role
```

🔒 Только `admin`

**Body:**
```json
{
  "role": "teacher"
}
```

Допустимые роли: `user`, `teacher`, `admin`

**Response `200`:**
```json
{
  "message": "role updated"
}
```

---

## Posts

### Получить все посты

```
GET /api/posts?page=1&limit=20
```

🔒 Требует токен

| Query-параметр | По умолчанию | Описание |
|----------------|-------------|----------|
| `page` | `1` | Номер страницы |
| `limit` | `20` | Постов на странице (макс. 100) |

**Response `200`:**
```json
[
  {
    "id": 1,
    "content": "Привет, мир!",
    "image": "",
    "author_id": 1,
    "author": {
      "id": 1,
      "first_name": "Иван",
      "last_name": "Иванов"
    },
    "community_id": null,
    "likes_count": 5,
    "comments_count": 2,
    "created_at": "2026-05-20T10:00:00Z"
  }
]
```

---

### Получить пост по ID

```
GET /api/posts/:id
```

🔒 Требует токен

**Response `200`:** объект поста (см. выше)  
**Response `404`:** `{"error": "post not found"}`

---

### Создать пост

```
POST /api/posts
```

🔒 Требует токен

**Body:**
```json
{
  "content": "Текст поста",
  "image": ""
}
```

Ограничения: `content` от 1 до 500 символов.

**Response `201`:** объект созданного поста

---

### Удалить пост

```
DELETE /api/posts/:id
```

🔒 Только автор поста

**Response `200`:**
```json
{
  "message": "post deleted"
}
```

**Response `403`:** `{"error": "forbidden"}` — не автор  
**Response `404`:** `{"error": "post not found"}`

---

## Comments

### Получить комментарии к посту

```
GET /api/posts/:id/comments
```

🔒 Требует токен

**Response `200`:**
```json
[
  {
    "id": 1,
    "content": "Отличный пост!",
    "author_id": 2,
    "author": { "id": 2, "first_name": "Мария" },
    "post_id": 1,
    "created_at": "2026-05-20T11:00:00Z"
  }
]
```

---

### Создать комментарий

```
POST /api/posts/:id/comments
```

🔒 Требует токен

**Body:**
```json
{
  "content": "Текст комментария"
}
```

Ограничения: от 1 до 100 символов.

**Response `201`:** объект созданного комментария

---

## Likes

### Поставить / убрать лайк

```
POST /api/posts/:id/like
```

🔒 Требует токен

Повторный запрос убирает лайк (toggle).

**Response `200`:**
```json
{
  "message": "like added",
  "like_count": 6
}
```

или

```json
{
  "message": "like removed",
  "like_count": 5
}
```

---

## Communities

### Получить все сообщества

```
GET /api/teacher/communities/all
```

🔒 Только `teacher` или `admin`

**Response `200`:** массив сообществ

---

### Получить сообщество по ID

```
GET /api/communities/:id
```

🔒 Требует токен

**Response `200`:**
```json
{
  "id": 1,
  "name": "IT-21",
  "description": "Группа первого курса",
  "avatar": "/uploads/avatars/community_1.jpg",
  "owner_id": 3,
  "owner": { "id": 3, "first_name": "Преподаватель" },
  "members_count": 25,
  "created_at": "2026-05-01T00:00:00Z"
}
```

---

### Создать сообщество

```
POST /api/teacher/communities/create
```

🔒 Только `teacher` или `admin`

**Body:**
```json
{
  "name": "IT-21",
  "description": "Группа первого курса"
}
```

Ограничения: `name` минимум 3 символа.

**Response `201`:** объект сообщества

---

### Удалить сообщество

```
DELETE /api/admin/communities/:id
```

🔒 Только `admin` или владелец сообщества

**Response `200`:** `{"message": "community deleted"}`  
**Response `403`:** `{"error": "forbidden"}`

---

### Загрузить аватар сообщества

```
POST /api/communities/:id/avatar
```

🔒 Требует токен

**Body:** `multipart/form-data`

| Поле | Тип | Описание |
|------|-----|----------|
| `avatar` | file | JPG / JPEG / PNG, макс. 5 МБ |

**Response `200`:** `{"avatar": "/uploads/avatars/1_...jpg"}`

---

### Вступить в сообщество

```
POST /api/communities/:id/join
```

🔒 Требует токен

**Response `200`:** `{"message": "joined community"}`  
**Response `400`:** `{"error": "already joined"}`

---

### Выйти из сообщества

```
POST /api/communities/:id/leave
```

🔒 Требует токен

**Response `200`:** `{"message": "left community"}`  
**Response `400`:** `{"error": "not a member"}`

---

### Получить участников сообщества

```
GET /api/communities/:id/members
```

🔒 Требует токен

**Response `200`:** массив пользователей

---

### Получить посты сообщества

```
GET /api/communities/:id/posts
```

🔒 Требует токен

**Response `200`:** массив постов

---

### Создать пост в сообществе

```
POST /api/communities/:id/posts
```

🔒 Требует токен

**Body:**
```json
{
  "content": "Объявление для группы",
  "image": ""
}
```

**Response `201`:** объект поста

---

## Schedule

### Получить расписание сообщества

```
GET /api/communities/:id/schedule
```

🔒 Требует токен

**Response `200`:**
```json
[
  {
    "id": 1,
    "title": "Лекция по Go",
    "description": "Тема: горутины",
    "date": "2026-05-25",
    "time": "09:00",
    "teacher_id": 3,
    "teacher": { "id": 3, "first_name": "Преподаватель" },
    "community_id": 1,
    "created_at": "2026-05-20T00:00:00Z"
  }
]
```

---

### Создать занятие

```
POST /api/teacher/communities/:id/schedule
```

🔒 Только `teacher` или `admin`

**Body:**
```json
{
  "title": "Лекция по Go",
  "description": "Тема: горутины",
  "date": "2026-05-25",
  "time": "09:00"
}
```

Ограничения: `title` минимум 3 символа.

**Response `201`:** объект занятия

---

### Обновить занятие

```
PUT /api/teacher/schedule/:id/update
```

🔒 Автор занятия или `admin`

**Body:** те же поля что при создании

**Response `200`:** `{"message": "schedule updated"}`  
**Response `403`:** `{"error": "forbidden"}`

---

### Удалить занятие

```
DELETE /api/teacher/schedule/:id/delete
```

🔒 Автор занятия или `admin`

**Response `200`:** `{"message": "schedule deleted"}`  
**Response `403`:** `{"error": "forbidden"}`

---

## Chats

### Получить все чаты

```
GET /api/chats
```

🔒 Требует токен

**Response `200`:** массив чатов

---

### Создать групповой чат

```
POST /api/chats
```

🔒 Требует токен

**Body:**
```json
{
  "name": "Общий чат IT-21"
}
```

Ограничения: `name` минимум 2 символа.

**Response `201`:**
```json
{
  "id": 1,
  "name": "Общий чат IT-21",
  "is_dm": false,
  "created_at": "2026-05-20T10:00:00Z"
}
```

---

### Обновить чат

```
PUT /api/chats/:id
```

🔒 Требует токен

**Body:**
```json
{
  "name": "Новое название"
}
```

**Response `200`:** `{"message": "chat updated"}`

---

### Удалить чат

```
DELETE /api/chats/:id
```

🔒 Требует токен

**Response `200`:** `{"message": "chat deleted"}`

---

### Добавить участника в чат

```
POST /api/chats/:id/members
```

🔒 Требует токен

**Body:**
```json
{
  "user_id": 2
}
```

**Response `200`:** `{"message": "member added"}`

---

### Получить участников чата

```
GET /api/chats/:id/members
```

🔒 Требует токен

**Response `200`:** массив пользователей

---

## Messages

### Получить сообщения чата

```
GET /api/chats/:id/messages
```

🔒 Требует токен

**Response `200`:**
```json
[
  {
    "id": 1,
    "content": "Привет всем!",
    "chat_id": 1,
    "sender_id": 1,
    "sender": { "id": 1, "first_name": "Иван" },
    "created_at": "2026-05-20T10:05:00Z"
  }
]
```

---

### Отправить сообщение

```
POST /api/chats/:id/messages
```

🔒 Только участник чата

**Body:**
```json
{
  "content": "Текст сообщения"
}
```

**Response `201`:** объект сообщения  
**Response `403`:** `{"error": "not a chat member"}`

---

## Direct Messages

### Создать личный чат

```
POST /api/dm
```

🔒 Требует токен

Если чат между этими двумя пользователями уже существует — возвращает существующий.

**Body:**
```json
{
  "user_id": 5
}
```

**Response `200`:** объект чата с `"is_dm": true`

---

## WebSocket

### Подключение к чату в реальном времени

```
GET /api/ws?chat_id=1
```

🔒 Требует токен (заголовок `Authorization: Bearer <token>`)

Подключение устанавливает WebSocket-соединение. Сообщения отправляются и принимаются в JSON.

**Пример отправки:**
```json
"Привет!"
```

Все участники с тем же `chat_id` получат это сообщение.

---

## Роли и доступ

| Роль | Возможности |
|------|-------------|
| `user` | Посты, комментарии, лайки, чаты, вступление в сообщества |
| `teacher` | Всё что `user` + создание сообществ, управление расписанием |
| `admin` | Всё + удаление сообществ, изменение ролей пользователей |

---

## Коды ответов

| Код | Значение |
|-----|----------|
| `200` | Успех |
| `201` | Создано |
| `400` | Ошибка валидации / неверные данные |
| `401` | Нет токена или токен недействителен |
| `403` | Нет прав доступа |
| `404` | Ресурс не найден |
| `500` | Внутренняя ошибка сервера |
