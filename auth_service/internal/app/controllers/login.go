package controllers

import (
	pb "auth_service/pkg/api/auth"
	"context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := validateLoginUserRequest(req); err != nil {
		return nil, err
	}

	loginInfo := loginUserFromPbLoginUserRequest(req)

	newLoginInfo, err := c.AuthUsecase.Login(ctx, loginInfo)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: newLoginInfo.Token.String()}, nil
}

// TODO: Не получилось подрубить buf.build/bufbuild/protovalidate
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
