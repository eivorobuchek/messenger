package handlers

import (
	pb "auth_service/pkg/api/auth"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}
