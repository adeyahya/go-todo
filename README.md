# Go Todo
This repository contains a project that I use to explore the Go programming language within the scope of web application development.

## Setup Project
```bash
docker compose build
```
### npm install
```bash
docker compose run frontend npm install
```

### create database
```bash
docker compose up db -d
docker run go-todo-db-1 bash
su - postgres
psql
create database todo;
\q
exit
exit
```

### run
```bash
docker compose up -d
```

### seeing development logs
```bash
docker compose logs --follow
```

### Create migration file
```bash
docker compose run backend migrate create -ext sql -dir database/migration/ -seq migration_name
```