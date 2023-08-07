package errors

import (
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, err error) {
	log.Println("Error:", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func NotFoundError(w http.ResponseWriter) {
	http.Error(w, "Not Found", http.StatusNotFound)
}
