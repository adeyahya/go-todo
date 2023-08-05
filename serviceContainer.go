package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/adeyahya/go-todo/handlers"
	"github.com/adeyahya/go-todo/repositories"
	"github.com/golang-migrate/migrate/v4"
)

type kernel struct {
	Db *sql.DB
}

func SetupDatabase() {
	// database migration
	m, initMigrationError := migrate.New("file://database/migration", "sqlite://database/data.db")
	if initMigrationError != nil {
		log.Fatal(initMigrationError)
	}
	m.Up()
}

func (k *kernel) InjectTodoHandler() handlers.TodoHandler {
	todoRepository := repositories.TodoRepository{DB: k.Db}
	todoHandler := handlers.TodoHandler{TodoRepository: &todoRepository}

	return todoHandler
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() *kernel {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
			db, dbErr := sql.Open("sqlite", "./database/data.db")
			if dbErr != nil {
				log.Fatal(dbErr)
			}
			k.Db = db
		})
	}

	return k
}
