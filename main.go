package main

import (
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

func main() {
	SetupDatabase()

	todoHandler := ServiceContainer().InjectTodoHandler()

	router := mux.NewRouter()

	router.HandleFunc("/todo", todoHandler.List).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
