package main

import (
	"auth_service/internal/app/controllers"
	"auth_service/internal/app/usecases"
	auth_repository "auth_service/internal/repositories/inmemory"
	"auth_service/server"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// repository

	authRepo := auth_repository.NewRepository(1000)

	// usecases
	authUsecase := usecases.NewAuthUsecase(usecases.AuthDeps{
		AuthRepository: authRepo,
	})

	authController := controllers.New(controllers.Deps{
		AuthUsecase: authUsecase,
	})

	// middlewares

	// infrastructure server
	config := server.Config{
		GRPCPort:        ":8082",
		GRPCGatewayPort: ":8080",
	}

	srv, err := server.New(ctx, config, server.Contollers{
		AuthServiceServer: authController,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}

}
