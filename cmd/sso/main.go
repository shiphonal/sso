package main

import (
	"log/slog"
	"sso/internal/config"
	"sso/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting logger", slog.Any("config", cfg))

	// TODO: инициализация приложения

	// TODO: инициализация приложения

	// TODO: инициализация grpc

}
