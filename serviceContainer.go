package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/adeyahya/go-todo/handlers"
	"github.com/adeyahya/go-todo/repositories"
	"github.com/golang-migrate/migrate/v4"
)

type IServiceContainer interface {
	InjectTodoHandler() handlers.TodoHandler
}

type kernel struct{}

var DB *sql.DB

func SetupDatabase() {
	// database migration
	m, initMigrationError := migrate.New("file://database/migration", "sqlite://database/data.db")
	if initMigrationError != nil {
		log.Fatal(initMigrationError)
	}
	m.Up()

	// init database
	db, dbErr := sql.Open("sqlite", "./database/data.db")
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	DB = db
	defer db.Close()
}

func (k *kernel) InjectTodoHandler() handlers.TodoHandler {
	todoRepository := repositories.TodoRepository{DB: DB}
	todoHandler := handlers.TodoHandler{TodoRepository: &todoRepository}

	return todoHandler
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}

	return k
}
