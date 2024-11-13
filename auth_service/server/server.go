package server

import (
	pb "auth_service/pkg/api/auth"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

// Config - server config
type Config struct {
	GRPCPort        string
	GRPCGatewayPort string

	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}

type Contollers struct {
	pb.AuthServiceServer
}

type Server struct {
	Contollers

	grpc struct {
		lis    net.Listener
		server *grpc.Server
	}

	grpcGateway struct {
		lis    net.Listener
		server *http.Server
	}
}

// New - returns *Server
func New(ctx context.Context, cfg Config, svcs Contollers) (*Server, error) {
	srv := &Server{Contollers: svcs}

	// grpc
	{
		// middlewares
		grpcServerOptions := unaryInterceptorsToGrpcServerOptions(cfg.UnaryInterceptors...)
		grpcServerOptions = append(grpcServerOptions,
			grpc.ChainUnaryInterceptor(cfg.ChainUnaryInterceptors...),
		)

		grpcServer := grpc.NewServer(grpcServerOptions...)
		// router
		pb.RegisterAuthServiceServer(grpcServer, srv)

		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", cfg.GRPCPort)
		if err != nil {
			return nil, fmt.Errorf("server: failed to listen: %v", err)
		}

		srv.grpc.server = grpcServer
		srv.grpc.lis = lis
	}

	// grpc gateway
	//{
	//	// router
	//	mux := runtime.NewServeMux()
	//	if err := pb.RegisterAuthServiceServer(mux, srv); err != nil {
	//		return nil, fmt.Errorf("server: failed to register handler: %v", err)
	//	}
	//
	//	// middlewares
	//	// ...
	//
	//	httpServer := &http.Server{Handler: mux}
	//
	//	lis, err := net.Listen("tcp", cfg.GRPCGatewayPort)
	//	if err != nil {
	//		return nil, fmt.Errorf("server: failed to listen: %v", err)
	//	}
	//
	//	srv.grpcGateway.server = httpServer
	//	srv.grpcGateway.lis = lis
	//}

	return srv, nil
}

// Run - serve
func (s *Server) Run(ctx context.Context) error {
	group := errgroup.Group{}

	group.Go(func() error {
		log.Println("start serve grpc", s.grpc.lis.Addr())
		if err := s.grpc.server.Serve(s.grpc.lis); err != nil {
			return fmt.Errorf("server: serve grpc: %v", err)
		}
		return nil
	})

	//group.Go(func() error {
	//	log.Println("start serve grpc gateway", s.grpcGateway.lis.Addr())
	//	if err := s.grpcGateway.server.Serve(s.grpcGateway.lis); err != nil {
	//		return fmt.Errorf("server: serve grpc gateway: %v", err)
	//	}
	//	return nil
	//})

	return group.Wait()
}

func unaryInterceptorsToGrpcServerOptions(interceptors ...grpc.UnaryServerInterceptor) []grpc.ServerOption {
	opts := make([]grpc.ServerOption, 0, len(interceptors))
	for _, interceptor := range interceptors {
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}
	return opts
}
