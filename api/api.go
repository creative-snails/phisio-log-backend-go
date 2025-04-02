package api

import (
	"encoding/json"
	"net/http"
)


type Error struct {
	Code int        `json:"code"`
	Message string  `json:"message"`
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := Error {
		Code: code,       
		Message: message,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}


