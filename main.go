package main

import (
	"log"
	"net"

	"advanced/app"
	"advanced/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. Eksekusi Dependency Injection Global (Config, DB, Logger, dll)
	_, err := app.InjectAppConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi environment: %v", err)
	}

	// 2. Load Dependencies gRPC User
	grpcHandler, err := app.InjectUserGRPC()
	if err != nil {
		log.Fatalf("Gagal meregistrasikan service user ke gRPC: %v", err)
	}

	secMiddleware, err := app.InjectSecurityMiddleware()
	if err != nil {
		log.Fatalf("Gagal me-load middleware gRPC: %v", err)
	}

	// 3. Memulai gRPC Server (Memblokir Main Thread)
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("gRPC gagal bind ke tcp port: %v", err)
	}

	// Daftarkan interceptor security di server gRPC
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(secMiddleware.GRPCAuthInterceptor()),
	)
	proto.RegisterUserServiceServer(grpcServer, grpcHandler)

	// Mendaftarkan Server Reflection untuk mempermudah Tester (seperti Postman)
	reflection.Register(grpcServer)

	log.Println("Server gRPC murni berjalan dan merajai port grpc://localhost:9090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Gagal menjalankan gRPC server: %v", err)
	}
}
