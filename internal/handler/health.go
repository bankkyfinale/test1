package handler

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthCheckHandler handles the health check route
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
