package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadAssignment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing form")
		return
	}

	// Retrieve file and metadata
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error retrieving file")
		return
	}
	defer file.Close()

	// Upload file to GCS
	fileURL, err := services.UploadFileToGCS(file, fileHeader)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("File upload failed: %v", err))
		return
	}

	assignment := createAssignmentFromRequest(r, fileURL)

	// Save assignment metadata in MongoDB
	result, err := services.CreateAssignment(assignment)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to save assignment")
		return
	}

	// Convert ObjectID to string
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve inserted ID")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Assignment uploaded successfully!",
		"id":      id.Hex(),
	})
}

// RUD operations to be implemented
func GetAssignment(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "GetAssignment not implemented yet")
}

func ListAssignments(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "ListAssignments not implemented yet")
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "UpdateAssignment not implemented yet")
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "DeleteAssignment not implemented yet")
}

func createAssignmentFromRequest(r *http.Request, fileURL string) models.Assignment {
	return models.Assignment{
		StudentID:   r.FormValue("student_id"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		FileURL:     fileURL,
		Status:      "pending",
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
