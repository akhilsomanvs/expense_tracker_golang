package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func writeJSON[T any](w http.ResponseWriter, status int, payload APIResponse[T]) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func OK[T any](w http.ResponseWriter, data T) {
	writeJSON(w, http.StatusOK, APIResponse[T]{
		Success: true,
		Data:    data,
	})
}

func Created[T any](w http.ResponseWriter, data T) {
	writeJSON(w, http.StatusCreated, APIResponse[T]{
		Success: true,
		Data:    data,
	})
}

func BadRequest(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusBadRequest, APIResponse[any]{
		Success: false,
		Message: message,
	})
}

func Unauthorized[T any](w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusUnauthorized, APIResponse[T]{
		Success: false,
		Message: message,
	})
}

func Forbidden[T any](w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusForbidden, APIResponse[T]{
		Success: false,
		Message: message,
	})
}

func NotFound[T any](w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusNotFound, APIResponse[T]{
		Success: false,
		Message: message,
	})
}

func Conflict[T any](w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusConflict, APIResponse[T]{
		Success: false,
		Message: message,
	})
}

func InternalServerError[T any](w http.ResponseWriter, messages ...string) {

	message := "Internal server error"

	if len(messages) > 0 && messages[0] != "" {
		message = messages[0]
	}

	writeJSON(w, http.StatusInternalServerError, APIResponse[T]{
		Success: false,
		Message: message,
	})
}
