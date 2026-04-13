package middleware

import (
	"context"

	"advanced/constants"
	"advanced/core/environment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// SecurityMiddleware bertugas menangani injeksi dan membalas fungsi perantara
type SecurityMiddleware struct {
	env *environment.Config
}

// ProvideMiddleware adalah provider wire
func ProvideMiddleware(env *environment.Config) *SecurityMiddleware {
	return &SecurityMiddleware{
		env: env,
	}
}



// GRPCAuthInterceptor bertugas mencegat (intercept) setiap panggilan RPC dan memeriksa metadata
func (m *SecurityMiddleware) GRPCAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Logika bypass: Membiarkan method RegisterUser lewat tanpa otentikasi
		if info.FullMethod == "/user.UserService/RegisterUser" {
			return handler(ctx, req)
		}

		// Mengambil Metadata dari gRPC context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "%s: Metadata tidak valid/hilang", constants.RCAuthError)
		}

		// Memeriksa Authorization secara spesifik
		tokens := md.Get("authorization")
		if len(tokens) == 0 || tokens[0] != m.env.SECRET_KEY {
			return nil, status.Errorf(codes.Unauthenticated, "%s: %s", constants.RCAuthError, constants.MsgAksesDitolak)
		}

		// Jika valid, teruskan ke method utama
		return handler(ctx, req)
	}
}
