package controllers

import (
	"auth_service/internal/app/models"
	"auth_service/internal/app/usecases"
	pb "auth_service/pkg/api/auth"
)

func registerUserFromPbReristerUserRequest(req *pb.RegisterRequest) *usecases.RegisterUser {

	return &usecases.RegisterUser{
		Email:    models.Email(req.GetEmail()),
		Password: req.GetPassword(),
	}
}

func loginUserFromPbLoginUserRequest(req *pb.LoginRequest) *usecases.LoginUser {

	return &usecases.LoginUser{
		Email:    models.Email(req.GetEmail()),
		Password: req.GetPassword(),
	}

}
