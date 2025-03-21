package api

import (
	"net/http"
	"backend/internal/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/api/health", HealthCheck)

	http.HandleFunc("/api/assignment/upload", handlers.UploadAssignment)
	http.HandleFunc("/api/assignment/get", handlers.GetAssignment)
	http.HandleFunc("/api/assignment/list", handlers.ListAssignments)
	http.HandleFunc("/api/assignment/update", handlers.UpdateAssignment)
	http.HandleFunc("/api/assignment/delete", handlers.DeleteAssignment)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is running... 🚀"))
}
