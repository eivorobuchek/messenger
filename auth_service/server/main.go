package main

import (
	pb "auth_service/pkg/api/auth"
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

type Server struct {
	pb.UnimplementedAuthServiceServer

	mx            sync.RWMutex
	registerUsers map[string]*string
}

func NewServer() *Server {
	return &Server{registerUsers: make(map[string]*string)}
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	implementation := NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, implementation)

	log.Printf("Auth service is running on port %v\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	if err := validateRegisterUserRequest(req); err != nil {
		return nil, err
	}

	email := req.GetEmail()
	password := req.GetPassword()
	if _, ok := s.registerUsers[email]; ok {

		errInfo := errdetails.BadRequest_FieldViolation{
			Field:       "email",
			Description: fmt.Sprintf("%s already exists", email),
		}

		return nil, status.Error(codes.AlreadyExists, errInfo.String())
	}

	s.mx.Lock()
	s.registerUsers[email] = &password
	s.mx.Unlock()

	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}

func validateRegisterUserRequest(req *pb.RegisterRequest) error {
	password := req.GetPassword()
	email := req.GetEmail()

	var violations []*errdetails.BadRequest_FieldViolation
	if len(email) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "email",
			Description: "empty",
		})
	}
	if len(password) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "password",
			Description: "empty",
		})
	}

	if len(violations) > 0 {
		rpcErr := status.New(codes.InvalidArgument, codes.InvalidArgument.String())

		detailedError, err := rpcErr.WithDetails(&errdetails.BadRequest{
			FieldViolations: violations,
		})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return detailedError.Err()
	}

	return nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := validateLoginUserRequest(req); err != nil {
		return nil, err
	}

	email := req.GetEmail()
	password := req.GetPassword()
	userPassword, ok := s.registerUsers[email]
	if !ok {
		errInfo := errdetails.BadRequest_FieldViolation{
			Field:       "email",
			Description: fmt.Sprintf("%s not found", email),
		}
		
		return nil, status.Error(codes.Unauthenticated, errInfo.String())
	}

	if *userPassword != password {
		errInfo := errdetails.BadRequest_FieldViolation{
			Field:       "password",
			Description: fmt.Sprintf("%s incorrect password", password),
		}

		return nil, status.Error(codes.Unauthenticated, errInfo.String())
	}

	return &pb.LoginResponse{Token: "dummy-token"}, nil
}

func validateLoginUserRequest(req *pb.LoginRequest) error {
	password := req.GetPassword()
	email := req.GetEmail()

	var violations []*errdetails.BadRequest_FieldViolation
	if len(email) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "email",
			Description: "empty",
		})
	}
	if len(password) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "password",
			Description: "empty",
		})
	}

	if len(violations) > 0 {
		rpcErr := status.New(codes.InvalidArgument, codes.InvalidArgument.String())

		detailedError, err := rpcErr.WithDetails(&errdetails.BadRequest{
			FieldViolations: violations,
		})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return detailedError.Err()
	}

	return nil
}
