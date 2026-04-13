//go:build wireinject
// +build wireinject

package app

import (
	"advanced/core/database"
	"advanced/core/environment"
	"advanced/core/middleware"
	"advanced/src/user"

	"github.com/google/wire"
)

var (
	// CoreModule menyuntikan hal-hal foundation dari aplikasi
	CoreModule = wire.NewSet(
		environment.ProvideConfig,
		database.ProvidePostgreSQL,
		middleware.ProvideMiddleware,
	)

	// UserModule menyuntikan logika domain User
	UserModule = wire.NewSet(
		user.ProvideUserRepository,
		user.ProvideUserUsecase,
		user.ProvideDeliveryGRPC,
	)

	AppModule = wire.NewSet(CoreModule, UserModule)
)

// InjectAppConfig
func InjectAppConfig() (*environment.Config, error) {
	panic(wire.Build(AppModule))
}

// InjectGRPC digunakan untuk membangun Delivery GRPC User
func InjectUserGRPC() (*user.DeliveryGRPC, error) {
	panic(wire.Build(AppModule))
}

// InjectSecurityMiddleware digunakan untuk meng-extract middleware untuk gRPC interceptor
func InjectSecurityMiddleware() (*middleware.SecurityMiddleware, error) {
	panic(wire.Build(AppModule))
}
