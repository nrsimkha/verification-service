# Сервис цензуры комментариев

## Описание
Сервис цензуры комментариев предназначен для проверки текста комментариев на соответствие установленным правилам. Сервис принимает текст комментария и возвращает результат валидации.

### Функционал
- HTTP-сервер с обработчиком POST-запросов.
- Валидация текста комментария.
- Ответ **200 OK**, если комментарий прошел проверку.
- Ответ **400 Bad Request**, если комментарий не прошел проверку.
- Поддержка сквозного идентификатора запроса.
- Журналирование всех HTTP-запросов (время, IP-адрес, HTTP-код ответа, уникальный ID запроса).

### Недопустимые слова
Список запрещённых слов и фраз хранится в файле `pkg/api/api.go` в переменной `spamWords`.

Пример запрещённых слов:
- `qwerty`
- `йцукен`
- `zxvbnm`

## Установка и запуск
1. Клонируйте репозиторий:
   ```bash
   git clone <URL_репозитория>
   ```
2. Перейдите в директорию проекта:
   ```bash
   cd cmd/server
   ```
3. Запустите сервер:
   ```bash
   go run .
   ```

## Эндпоинты API
- **POST** `/verification` — проверка комментария на соответствие правилам
  - **Тело запроса:**
    ```json
    {
      	"NewsID":   9,
		"TextBody": "dsd! The simplicity йцукен of Go is definitely one of its biggest strengths. Its concurrency model, powered by goroutines and channels, makes handling parallel tasks so much easier. It's great to see how Go prioritizes efficiency without overcomplicating things, which is why it's become so popular for building high-performance systems.",
		"ParentID": {
			"Int64": 2,
			"Valid": true
		},
		"PubTime": 1257894002
    }
    ```
  - **Ответ:**
    - **200 OK** — комментарий прошел проверку
    - **400 Bad Request** — комментарий не прошел проверку

## Технологии
- **Язык программирования:** Go
- **Роутер:** [github.com/gorilla/mux](https://github.com/gorilla/mux)
- **API:** RESTful API

