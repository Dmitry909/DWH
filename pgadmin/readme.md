# PGADMIN для удобства

Для запуска `docker compose up`.

Логин `admin@example.com`, пароль `admin` (источник истины — docker-compose), затем Add new server, Name любое (например, `cloud`), вкладка Connection: host — возьмите [отсюда](https://console.yandex.cloud/folders/b1g4sjmc9tl8nf11vnbm/managed-postgresql/cluster/c9qeafr997bjppskeru4/hosts) хост **MASTER** не **REPLICA**, port `6432`, database `db1`, username `user1`, password — сами знаете, save password — можно установить в true, вкладка parameters: SSL mode — verify-full. В случае неудач обращаться к `@sermir2003`.
