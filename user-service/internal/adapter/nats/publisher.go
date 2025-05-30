package nats

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aftosmiros/moodu/user-service/internal/adapter/nats/dto"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(conn *nats.Conn) *Publisher {
	return &Publisher{conn: conn}
}

func (p *Publisher) Publish(subject string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return p.conn.Publish(subject, data)
}

func (p *Publisher) PublishUserCreated(ctx context.Context, event *dto.UserCreatedEvent) error {
	return p.Publish("user.created", event)
}

func (p *Publisher) PublishUserUpdated(ctx context.Context, event *dto.UserUpdatedEvent) error {
	return p.Publish("user.updated", event)
}

func (p *Publisher) PublishUserDeleted(ctx context.Context, event *dto.UserDeletedEvent) error {
	return p.Publish("user.deleted", event)
}

func (p *Publisher) PublishLoginAttempt(ctx context.Context, event *dto.LoginAttemptEvent) error {
	return p.Publish("user.login_attempt", event)
}
