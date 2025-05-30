package mongodb

import (
	"context"
	"errors"
	"time"

	"cs/course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssignmentRepository interface {
	Create(*model.Assignment) error
	Update(*model.Assignment) error
	Delete(id string) error
	GetByID(id string) (*model.Assignment, error)
	List() ([]*model.Assignment, error)
	AddSubmission(assignmentID string, submissionID string) error
	AddSubmissions(assignmentID string, submissionIDs []string) error
}

type MongoAssignmentRepo struct {
	collection *mongo.Collection
}

func NewAS(db *mongo.Database) *MongoAssignmentRepo {
	return &MongoAssignmentRepo{
		collection: db.Collection("assignments"),
	}
}

func (r *MongoAssignmentRepo) Create(a *model.Assignment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(ctx, a)
	return err
}

func (r *MongoAssignmentRepo) GetByID(id string) (*model.Assignment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var a model.Assignment
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&a)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("assignment not found")
		}
		return nil, err
	}
	return &a, nil
}

func (r *MongoAssignmentRepo) List() ([]*model.Assignment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var assignments []*model.Assignment
	for cursor.Next(ctx) {
		var a model.Assignment
		if err := cursor.Decode(&a); err != nil {
			return nil, err
		}
		assignments = append(assignments, &a)
	}
	return assignments, nil
}

func (r *MongoAssignmentRepo) Update(a *model.Assignment) error {
	filter := bson.M{"_id": a.ID}
	update := bson.M{"$set": bson.M{
		"title":       a.Title,
		"description": a.Description,
		"course_id":   a.CourseID,
		"due_date":    a.DueDate,
	}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoAssignmentRepo) Delete(id string) error {
	// objID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return err
	// }

	res, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("assignment not found")
	}
	return nil
}

func (r *MongoAssignmentRepo) AddSubmission(assignmentID string, submissionID string) error {
	objID, err := primitive.ObjectIDFromHex(assignmentID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$addToSet": bson.M{
			"submissions": submissionID,
		},
		"$set": bson.M{
			"updated_at": time.Now().Format(time.RFC3339),
		},
	}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoAssignmentRepo) AddSubmissions(assignmentID string, submissionIDs []string) error {
	objID, err := primitive.ObjectIDFromHex(assignmentID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$addToSet": bson.M{"submissions": bson.M{"$each": submissionIDs}},
		"$set":      bson.M{"updated_at": time.Now().Format(time.RFC3339)},
	}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
