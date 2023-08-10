package errors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field    string `json:"field"`
	Msg      string `json:"message"`
	Expected string `json:"expected"`
}

func formatValidatorError(err error, w http.ResponseWriter) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), fe.Tag(), fe.Param()}
		}
		json.NewEncoder(w).Encode(out)
	}
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Println("Error:", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func NotFoundError(w http.ResponseWriter) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

func BadRequest(w http.ResponseWriter, err error) {
	log.Println("Error:", err)
	formatValidatorError(err, w)
	w.WriteHeader(http.StatusBadRequest)
}

func UnauthorizedError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusUnauthorized)
}
