package app

import (
	"context"
	"log"

	"github.com/aftosmiros/moodu/user-service/config"
	"github.com/aftosmiros/moodu/user-service/internal/adapter/grpc/server"
	"github.com/aftosmiros/moodu/user-service/internal/adapter/mongo"
	nats "github.com/aftosmiros/moodu/user-service/internal/adapter/nats"
	redisadapter "github.com/aftosmiros/moodu/user-service/internal/adapter/redis"
	"github.com/aftosmiros/moodu/user-service/internal/usecase"
	mongopkg "github.com/aftosmiros/moodu/user-service/pkg/mongo"
	natshelper "github.com/aftosmiros/moodu/user-service/pkg/nats"
	"github.com/aftosmiros/moodu/user-service/pkg/security"
	"github.com/go-redis/redis/v8"
)

type App struct {
	api *server.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	// Mongo
	mongoDB, err := mongopkg.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, err
	}

	userRepo := mongo.NewUserRepository(mongoDB.Conn)
	sessionRepo := mongo.NewRefreshToken(mongoDB.Conn)
	if err := userRepo.EnsureIndexes(ctx); err != nil {
		return nil, err
	}

	// Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
		DB:   cfg.Redis.DB,
	})
	cache := redisadapter.NewRedisCache(redisClient)

	// NATS
	natsConn, err := natshelper.NewNatsConn(*cfg)
	if err != nil {
		return nil, err
	}
	producer := nats.NewPublisher(natsConn)

	// jwt
	jwtManager := security.NewJWTManager(cfg.JWT.Secret)
	passwordManager := security.NewPasswordManager()

	// Usecase
	userUC := usecase.NewUserService(userRepo, sessionRepo, cache, producer, jwtManager, passwordManager)

	// gRPC server
	api := server.New(
		cfg.GRPCServer,
		userUC,
		cfg.JWT.Secret,
	)

	return &App{api: api}, nil
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	a.api.Run(ctx, errCh)

	select {
	case <-ctx.Done():
		log.Println("Graceful shutdown...")
		return a.api.Stop(ctx)
	case err := <-errCh:
		return err
	}
}
