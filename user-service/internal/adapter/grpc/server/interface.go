package server

import (
	"github.com/aftosmiros/moodu/user-service/internal/adapter/grpc/server/frontend"
)

type UserUsecase interface {
	frontend.UserUsecase
}
