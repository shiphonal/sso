package validator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	ssov1 "sso/protos/gen/go/sso"
)

func ValidateRegister(req *ssov1.RegisterRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "UserId is empty")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "UserId is empty")
	}

	return nil
}
