package main

import (
	"log"
	"net/http"

	"github.com/adeyahya/go-todo/core/middleware"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

func main() {
	SetupDatabase()
	kernel := ServiceContainer()
	defer kernel.Db.Close()

	// handlers
	todoHandler := kernel.InjectTodoHandler()

	router := mux.NewRouter()
	router.Use(middleware.JsonMiddleware)

	router.HandleFunc("/todo", todoHandler.List).Methods(http.MethodGet)
	router.HandleFunc("/todo", todoHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", todoHandler.Get).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", muxHandlers.CompressHandler(router))
}
