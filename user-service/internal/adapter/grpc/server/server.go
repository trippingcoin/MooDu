package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aftosmiros/moodu/user-service/config"
	"github.com/aftosmiros/moodu/user-service/internal/adapter/grpc/server/frontend"
	"github.com/aftosmiros/moodu/user-service/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type API struct {
	s           *grpc.Server
	cfg         config.GRPCServer
	addr        string
	userUsecase frontend.UserUsecase
	jwtSecret   string
}

func New(cfg config.GRPCServer, userUsecase frontend.UserUsecase, jwtSecret string) *API {
	return &API{
		cfg:         cfg,
		addr:        fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		userUsecase: userUsecase,
		jwtSecret:   jwtSecret,
	}
}

func (a *API) Run(ctx context.Context, errCh chan<- error) {
	go func() {
		log.Println("🚀 gRPC server listening on", a.addr)

		if err := a.run(ctx); err != nil {
			errCh <- fmt.Errorf("cannot start grpc server: %w", err)
			return
		}
	}()
}

func (a *API) run(ctx context.Context) error {
	a.s = grpc.NewServer(a.setOptions(ctx, a.jwtSecret)...)

	// Register your service
	userpb.RegisterUserServiceServer(a.s, frontend.NewUserHandler(a.userUsecase))

	reflection.Register(a.s)

	lis, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	return a.s.Serve(lis)
}

func (a *API) Stop(ctx context.Context) error {
	if a.s == nil {
		return nil
	}

	stopped := make(chan struct{})
	go func() {
		a.s.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		a.s.Stop()
	case <-stopped:
	}

	return nil
}
