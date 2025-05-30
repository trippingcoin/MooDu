package repo

import (
	"context"
	"time"

	"admin/admin-service/internal/model"
	"admin/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdminRepo struct {
	queueColl      *mongo.Collection
	transcriptColl *mongo.Collection
	scheduleColl   *mongo.Collection
	retakeColl     *mongo.Collection
	certColl       *mongo.Collection
}

func NewMongoAdminRepo(db *mongo.Database) *MongoAdminRepo {
	return &MongoAdminRepo{
		queueColl:      db.Collection("queue"),
		transcriptColl: db.Collection("transcripts"),
		scheduleColl:   db.Collection("schedule"),
		retakeColl:     db.Collection("retakes"),
		certColl:       db.Collection("certificates"),
	}
}

func (r *MongoAdminRepo) CreateTranscriptRequest(studentID, purpose string) error {
	t := model.Transcript{
		StudentID: studentID,
		Purpose:   purpose,
		CreatedAt: time.Now(),
	}
	_, err := r.transcriptColl.InsertOne(context.TODO(), t)
	return err
}

func (r *MongoAdminRepo) ViewQueue() ([]*pb.QueueEntry, error) {
	cursor, err := r.queueColl.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var results []*pb.QueueEntry
	for cursor.Next(context.TODO()) {
		var entry model.QueueEntry
		if err := cursor.Decode(&entry); err != nil {
			return nil, err
		}
		results = append(results, &pb.QueueEntry{
			StudentId: entry.StudentID,
			Reason:    entry.Reason,
			JoinedAt:  entry.JoinedAt,
		})
	}
	return results, nil
}

func (r *MongoAdminRepo) JoinQueue(studentID, reason string) error {
	entry := model.QueueEntry{
		StudentID: studentID,
		Reason:    reason,
		JoinedAt:  time.Now().Format(time.RFC3339),
	}
	_, err := r.queueColl.InsertOne(context.TODO(), entry)
	return err
}

func (r *MongoAdminRepo) RegisterRetake(studentID, courseID, reason string) error {
	rt := model.Retake{
		StudentID: studentID,
		CourseID:  courseID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}
	_, err := r.retakeColl.InsertOne(context.TODO(), rt)
	return err
}

func (r *MongoAdminRepo) GetSchedule(studentID string) ([]*pb.ScheduleEntry, error) {
	cursor, err := r.scheduleColl.Find(context.TODO(), bson.M{"student_id": studentID})
	if err != nil {
		return nil, err
	}
	var results []*pb.ScheduleEntry
	for cursor.Next(context.TODO()) {
		var entry model.ScheduleEntry
		if err := cursor.Decode(&entry); err != nil {
			return nil, err
		}
		results = append(results, &pb.ScheduleEntry{
			CourseId: entry.CourseID,
			Day:      entry.Day,
			Time:     entry.Time,
			Room:     entry.Room,
		})
	}
	return results, nil
}

func (r *MongoAdminRepo) UpdateSchedule(courseID, day, timeStr, room string) error {
	update := bson.M{
		"$set": model.ScheduleEntry{
			CourseID: courseID,
			Day:      day,
			Time:     timeStr,
			Room:     room,
		},
	}
	_, err := r.scheduleColl.UpdateOne(context.TODO(), bson.M{"course_id": courseID}, update, options.Update().SetUpsert(true))
	return err
}

func (r *MongoAdminRepo) SubmitCertificateRequest(studentID, certType, details string) error {
	cr := model.CertificateRequest{
		StudentID:       studentID,
		CertificateType: certType,
		Details:         details,
		CreatedAt:       time.Now(),
	}
	_, err := r.certColl.InsertOne(context.TODO(), cr)
	return err
}
