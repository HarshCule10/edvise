package services

import (
	"backend/config"
	"backend/internal/models"
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get collection
var assignmentCollection *mongo.Collection
var once sync.Once

func getAssignmentCollection() *mongo.Collection {
	once.Do(func() {
		if config.DB == nil {
			log.Fatal("Database connection is not initialized. Call ConnectDB first.")
		}
		assignmentCollection = config.DB.Collection("assignments")
	})
	return assignmentCollection
}

func CreateAssignment(assignment models.Assignment) (*mongo.InsertOneResult, error) {
	assignment.ID = primitive.NewObjectID()
	assignment.SubmittedAt = time.Now()
	return getAssignmentCollection().InsertOne(context.TODO(), assignment)
}

func GetAssignment(id string) (models.Assignment, error) {
	var assignment models.Assignment
	objID, _ := primitive.ObjectIDFromHex(id)
	err := getAssignmentCollection().FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&assignment)
	return assignment, err
}

func ListAssignments() ([]models.Assignment, error) {
	var assignments []models.Assignment
	cursor, err := getAssignmentCollection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var assignment models.Assignment
		cursor.Decode(&assignment)
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

func UpdateAssignment(id string, updatedAssignment models.Assignment) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": updatedAssignment}
	_, err := getAssignmentCollection().UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return err
}

func DeleteAssignment(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := getAssignmentCollection().DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
