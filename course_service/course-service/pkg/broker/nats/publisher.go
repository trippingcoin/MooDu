package nats

import (
	"encoding/json"
	"log"

	"cs/course-service/internal/model"

	"github.com/nats-io/nats.go"
)

type NATSPublisher struct {
	conn *nats.Conn
}

func NewNATSPublisher(url string) (*NATSPublisher, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NATSPublisher{conn: nc}, nil
}

func (p *NATSPublisher) PublishCourseCreated(course *model.Course) error {
	msg, err := json.Marshal(course)
	if err != nil {
		return err
	}

	subject := "course.created"
	if err := p.conn.Publish(subject, msg); err != nil {
		return err
	}

	log.Println("Published course.created to NATS")
	return nil
}

func (p *NATSPublisher) PublishCourseUpdated(course *model.Course) error {
	msg, err := json.Marshal(course)
	if err != nil {
		return err
	}

	subject := "course.updated"
	if err := p.conn.Publish(subject, msg); err != nil {
		return err
	}

	log.Println("Published course.updated to NATS")
	return nil
}

func (p *NATSPublisher) PublishCourseDeleted(courseID string) error {
	msg, err := json.Marshal(map[string]string{"id": courseID})
	if err != nil {
		return err
	}

	subject := "course.deleted"
	if err := p.conn.Publish(subject, msg); err != nil {
		return err
	}

	log.Println("Published course.deleted to NATS")
	return nil
}

func (p *NATSPublisher) PublishAssignmentCreated(course *model.Assignment) error {
	msg, err := json.Marshal(course)
	if err != nil {
		return err
	}

	subject := "assignment.created"
	if err := p.conn.Publish(subject, msg); err != nil {
		return err
	}

	log.Println("Published assignment.created to NATS")
	return nil
}

func (p *NATSPublisher) PublishAssignmentUpdated(course *model.Assignment) error {
	msg, err := json.Marshal(course)
	if err != nil {
		return err
	}

	subject := "assignment.updated"
	if err := p.conn.Publish(subject, msg); err != nil {
		return err
	}

	log.Println("Published assignment.updated to NATS")
	return nil
}

func (p *NATSPublisher) PublishAssignmentDeleted(courseID string) error {
	msg, err := json.Marshal(map[string]string{"id": courseID})
	if err != nil {
		return err
	}

	subject := "assignment.deleted"
	if err := p.conn.Publish(subject, msg); err != nil {
		return err
	}

	log.Println("Published assignment.deleted to NATS")
	return nil
}
