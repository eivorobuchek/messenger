package controllers

import (
	pb "auth_service/pkg/api/auth"
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	if err := validateRegisterUserRequest(req); err != nil {
		return nil, err
	}

	registerInfo := registerUserFromPbReristerUserRequest(req)

	newRegister, err := c.AuthUsecase.RegisterUser(ctx, registerInfo)
	if err != nil {
		return nil, err
	}

	response := &pb.RegisterResponse{
		Message: fmt.Sprintf("User registered successfully with email: %s", newRegister.Email),
	}

	return response, nil
}

func rpcDetailedInfo(field, description string) errdetails.BadRequest_FieldViolation {
	return errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: fmt.Sprintf("%s: %s", field, description),
	}
}

// TODO: Не получилось подрубить buf.build/bufbuild/protovalidate
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
