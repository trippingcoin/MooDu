package usecase

import (
	"context"
	"time"

	"github.com/aftosmiros/moodu/user-service/internal/adapter/nats/dto"
	"github.com/aftosmiros/moodu/user-service/internal/domain"
	"github.com/aftosmiros/moodu/user-service/pkg/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo        UserRepository
	sessionRepo     RefreshTokenRepo
	cache           RedisCache
	producer        Publisher
	jwtManager      *security.JWTManager
	passwordManager *security.PasswordManager
}

func NewUserService(
	userRepo UserRepository,
	sessionRepo RefreshTokenRepo,
	cache RedisCache,
	producer Publisher,
	jwtManager *security.JWTManager,
	passwordManager *security.PasswordManager,
) *UserService {
	return &UserService{
		userRepo:        userRepo,
		sessionRepo:     sessionRepo,
		cache:           cache,
		producer:        producer,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
	}
}

func (uc *UserService) Register(ctx context.Context, user *domain.User, plainPassword string) (string, error) {
	hashedPwd, err := uc.passwordManager.HashPassword(plainPassword)
	if err != nil {
		return "", err
	}
	user.PasswordHash = hashedPwd
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().UTC()

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return "", err
	}

	if err := uc.cache.SetUser(ctx, user, 24*time.Hour); err != nil {
		return "", err // не фейлим, но логируем в будущем
	}

	event := &dto.UserCreatedEvent{
		ID:       user.ID.Hex(),
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	}

	if err := uc.producer.PublishUserCreated(ctx, event); err != nil {
		return "", err
	}

	return user.ID.Hex(), nil
}

func (uc *UserService) Login(ctx context.Context, email, password string) (*domain.Token, error) {
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := uc.passwordManager.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, err
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(user.ID.Hex(), user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.jwtManager.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		return nil, err
	}

	session := domain.Session{
		UserID:       user.ID.Hex(),
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	return &domain.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UserService) GetProfile(ctx context.Context, userID string) (*domain.User, error) {
	user, err := uc.cache.GetUser(ctx, userID)
	if err == nil {
		return user, nil
	}

	user, err = uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	_ = uc.cache.SetUser(ctx, user, 24*time.Hour) // можно логировать ошибку

	return user, nil
}

func (uc *UserService) RefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	session, err := uc.sessionRepo.GetByToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if session.ExpiresAt.Before(time.Now().UTC()) {
		return nil, domain.ErrSessionExpired
	}

	user, err := uc.userRepo.GetByID(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(user.ID.Hex(), user.Role)
	if err != nil {
		return nil, err
	}
	newRefreshToken, err := uc.jwtManager.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		return nil, err
	}

	if err := uc.sessionRepo.DeleteByToken(ctx, refreshToken); err != nil {
		return nil, err
	}

	newSession := domain.Session{
		UserID:       user.ID.Hex(),
		RefreshToken: newRefreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}
	if err := uc.sessionRepo.Create(ctx, newSession); err != nil {
		return nil, err
	}

	return &domain.Token{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
