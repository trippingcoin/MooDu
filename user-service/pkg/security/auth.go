package security

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

func TokenFromCtx(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return "", false
	}

	tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")

	return tokenStr, true
}
