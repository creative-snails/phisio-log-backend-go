package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success code, usuall 200
	Code int

	// Account balance
	Balance int64
}

type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

func RequestErrorHandler(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusBadRequest)
}

func InternalErrorHandler(w http.ResponseWriter) {
	writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
}

