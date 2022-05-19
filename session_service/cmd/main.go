package main

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/vgeorgo/grpc-envoy-experiment/session_service_go/internal/svc"
	sessionV1 "github.com/vgeorgo/grpc-envoy-experiment/session_service_go/pkg/pb/session/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
)

func main() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	initListener()
	initServer()
	initServices()

	// Executes a goroutine apart from the application
	go signalsListener(server)

	logger.Info("Starting gRPC server...")
	if err := server.Serve(listener); err != nil {
		logger.Panic("Failed to start gRPC server", zap.Error(err))
	}
}

func initListener() {
	var err error
	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("SESSION_SERVICE_PORT"))

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Panic("Failed to listen", zap.String("address", addr), zap.Error(err))
	}

	logger.Info("Started listening...", zap.String("address", addr))
	return
}

func initServer() {
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpcZap.Option{
		grpcZap.WithLevels(grpcZap.DefaultCodeToLevel),
		grpcZap.WithDecider(func(fullMethodName string, err error) bool {
			// will not log gRPC calls if it was a call to healthcheck and no error was raised
			//if err == nil && fullMethodName == "foo.bar.healthcheck" {
			//	return false
			//}

			// by default everything will be logged
			return true
		}),
	}

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpcZap.ReplaceGrpcLoggerV2(logger)

	server = grpc.NewServer(
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcZap.StreamServerInterceptor(logger, opts...),
			// grpc_auth.StreamServerInterceptor(myAuthFunction),
		)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcZap.UnaryServerInterceptor(logger, opts...),
			// grpc_auth.UnaryServerInterceptor(myAuthFunction),
		)))
}

func initServices() {
	sessionV1.RegisterJWTServiceServer(server, &svc.JWTServer{})

	logger.Info("Handlers registered")
}

func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	_ = <-sigs

	logger.Info("Gracefully stopping server...")
	server.GracefulStop()
}
