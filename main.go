package main

import (
	"log"
	"net/http"

	"database/sql"

	"github.com/adeyahya/go-todo/repository"
	"github.com/adeyahya/go-todo/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

func main() {
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
	defer db.Close()

	// repositories
	todoRepository := repository.Todo{
		Db: db,
	}

	// services
	todoService := service.Todo{
		Repository: &todoRepository,
	}

	router := mux.NewRouter()

	router.HandleFunc("/todo", todoService.GetAllTodo).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
