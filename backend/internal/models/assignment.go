package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Assignment represents a student assignment stored in MongoDB
type Assignment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StudentID   string             `bson:"student_id" json:"student_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	FileURL     string             `bson:"file_url" json:"file_url"`
	Status      string             `bson:"status" json:"status"` // pending, graded, reviewed
	SubmittedAt time.Time          `bson:"submitted_at" json:"submitted_at"`
}
