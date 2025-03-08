package validator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	ssov1 "sso/protos/gen/go/sso"
)

const (
	emptyValue = 0
)

func ValidateLogin(req *ssov1.LoginRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}
	if req.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "missing app id")
	}
	return nil
}
