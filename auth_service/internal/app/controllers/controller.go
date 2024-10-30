package controllers

import (
	"auth_service/internal/app/usecases"
	pb "auth_service/pkg/api/auth"
)

// Deps - server deps
type Deps struct {
	AuthUsecase usecases.AuthUsecase
}

// Controller - реализация обработчика gRPC/REST запросов
type Controller struct {
	pb.UnimplementedAuthServiceServer
	Deps
}

// New - returns *Controller
func New(d Deps) *Controller {
	return &Controller{
		Deps: d,
	}
}
