package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(code int, msg string) APIError {
	return APIError{
		StatusCode: code,
		Message:    msg,
	}
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiError, ok := err.(APIError); ok {
				WriteJSON(w, apiError.StatusCode, apiError)
			} else {
				WriteJSON(w, http.StatusInternalServerError, "internal server error")
			}
		}
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}
