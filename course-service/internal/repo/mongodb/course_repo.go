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

type MongoCourseRepo struct {
	collection *mongo.Collection
}

func NewCourseRepo(db *mongo.Database) *MongoCourseRepo {
	return &MongoCourseRepo{
		collection: db.Collection("courses"),
	}
}

func (r *MongoCourseRepo) Create(course *model.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	course.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(ctx, course)
	return err
}

func (r *MongoCourseRepo) Update(course *model.Course) error {
	objID, err := primitive.ObjectIDFromHex(course.ID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":       course.Title,
			"description": course.Description,
			"teacherid":   course.TeacherID,
		},
	}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoCourseRepo) GetByID(id string) (*model.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var course model.Course
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&course)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("course not found")
		}
		return nil, err
	}
	return &course, nil
}

func (r *MongoCourseRepo) List() ([]*model.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courses []*model.Course
	for cursor.Next(ctx) {
		var course model.Course
		if err := cursor.Decode(&course); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (r *MongoCourseRepo) Delete(id string) error {
	filter := bson.M{"id": id}
	res, err := r.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("course not found")
	}
	return nil
}
