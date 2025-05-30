package frontend

import (
	"context"

	dto "github.com/aftosmiros/moodu/user-service/internal/adapter/grpc/server/frontend/dto"
	"github.com/aftosmiros/moodu/user-service/proto/userpb"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	usecase UserUsecase
}

func NewUserHandler(uc UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := dto.FromRegisterRequest(req)

	id, err := h.usecase.Register(ctx, user, req.Password)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterResponse{UserId: id}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	token, err := h.usecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &userpb.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (h *UserHandler) GetProfile(ctx context.Context, req *userpb.GetProfileRequest) (*userpb.GetProfileResponse, error) {
	user, err := h.usecase.GetProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return dto.ToProfileResponse(user), nil
}
