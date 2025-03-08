package validator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	ssov1 "sso/protos/gen/go/sso"
)

func ValidateIsAdmin(req *ssov1.IsAdminRequest) error {
	if req.GetUserId() == emptyValue {
		return status.Error(codes.InvalidArgument, "missing user_id")
	}

	return nil
}
