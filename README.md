# Сервис укорачивания ссылок

gRPC-сервис на Go и PostgreSQL. Тестовое задание Ozon.

## Сборка и запуск

**Первый способ:**

```bash
docker pull n3kto/shorturl
docker run -p 127.0.0.1:8080:8080 -e HOST=:8080 -e DB_URL=postgresql://postgres:qwerty@127.0.0.1:5432/shorturl n3kto/shorturl
```

---

**Второй способ:**

```bash
docker-compose build server
docker-compose run -p 127.0.0.1:8080:8080 -d server
```

---

**Третий способ:**

```bash
# production
docker build -t shorturl .

# development
docker build -t shorturl -f Dockerfile.dev .

# test
docker build -t shorturl -f Dockerfile.test .


docker run -p 127.0.0.1:8080:8080 -e HOST=:8080 -e DB_URL=postgresql://postgres:qwerty@database:5432/shorturl shorturl
```

---

**Четвертый способ:**

```bash
# production
docker-compose -f docker-compose.dev.yml build server
docker-compose -f docker-compose.dev.yml -p 127.0.0.1:8080:8080 run server

# development
docker-compose -f docker-compose.dev.yml build server-development
docker-compose -f docker-compose.dev.yml -p 127.0.0.1:8080:8080 run server-development

# test
docker-compose -f docker-compose.dev.yml build server-test
docker-compose -f docker-compose.dev.yml -p 127.0.0.1:8080:8080 run server-test
```
