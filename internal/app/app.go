package app

import (
	grpcapp "github.com/AlexanderVasiliev23/go-grpc-auth/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
