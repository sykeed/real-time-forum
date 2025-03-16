package handlers

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, message string ,err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
