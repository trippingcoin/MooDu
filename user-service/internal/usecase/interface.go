package usecase

import (
	"context"
	"time"

	"github.com/aftosmiros/moodu/user-service/internal/adapter/nats/dto"
	"github.com/aftosmiros/moodu/user-service/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByBarcode(ctx context.Context, barcode string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	SoftDelete(ctx context.Context, id string) error
	EnsureIndexes(ctx context.Context) error
}

type RefreshTokenRepo interface {
	Create(ctx context.Context, session domain.Session) error
	GetByToken(ctx context.Context, token string) (domain.Session, error)
	DeleteByToken(ctx context.Context, token string) error
}

type RedisCache interface {
	GetUser(ctx context.Context, barcode string) (*domain.User, error)
	SetUser(ctx context.Context, user *domain.User, ttl time.Duration) error
	DeleteUser(ctx context.Context, barcode string) error
}

type Publisher interface {
	Publish(subject string, v any) error
	PublishUserCreated(ctx context.Context, event *dto.UserCreatedEvent) error
	PublishUserUpdated(ctx context.Context, event *dto.UserUpdatedEvent) error
	PublishUserDeleted(ctx context.Context, event *dto.UserDeletedEvent) error
	PublishLoginAttempt(ctx context.Context, event *dto.LoginAttemptEvent) error
}
