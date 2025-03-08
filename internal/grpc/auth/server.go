package auth

import ssov1 "protos/gen/go/sso"

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
