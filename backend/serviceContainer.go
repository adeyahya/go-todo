package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	utils "github.com/adeyahya/go-todo/core"
	"github.com/adeyahya/go-todo/handlers"
	"github.com/adeyahya/go-todo/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
)

type kernel struct {
	Db       *sql.DB
	Validate *validator.Validate
}

func SetupDatabase() {
	dbConfig := utils.GetDatabaseConfig()
	// database migration
	m, initMigrationError := migrate.New(
		"file://database/migration",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Name,
		),
	)
	if initMigrationError != nil {
		log.Fatal(initMigrationError)
	}
	m.Up()
}

func (k *kernel) InjectTodoHandler() handlers.TodoHandler {
	todoRepository := repositories.TodoRepository{DB: k.Db}
	todoHandler := handlers.TodoHandler{
		TodoRepository: &todoRepository,
		Validate:       k.Validate,
	}

	return todoHandler
}

func (k *kernel) InjectAuthHandler() handlers.AuthHandler {
	userRepository := repositories.UserRepository{DB: k.Db}
	authHandler := handlers.AuthHandler{
		UserRepository: &userRepository,
		Validate:       k.Validate,
	}

	return authHandler
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() *kernel {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
			dbConfig := utils.GetDatabaseConfig()
			db, dbErr := sql.Open("postgres",
				fmt.Sprintf(
					"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
					dbConfig.Host,
					dbConfig.Port,
					dbConfig.User,
					dbConfig.Password,
					dbConfig.Name,
				),
			)
			if dbErr != nil {
				log.Fatal(dbErr)
			}
			k.Validate = validator.New()
			k.Db = db
		})
	}

	return k
}
