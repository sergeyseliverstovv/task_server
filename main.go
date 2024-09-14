package main

import (
	"log/slog"
	"os"

	"github.com/sergeyseliverstovv/task_server/internal/config"
)

func main() {
	// TODO: инициалищировать конфиг
	cfg := config.MustLoad()

	// TODO: инициализация логгера
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	log.Info("start serves", slog.Any("env", cfg)) // Потом убрать

	// TODO: инициализация приложения

	//TODO: запустить gRPC сервер

	//TODO: запусить REST сервер
}
