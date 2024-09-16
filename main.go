package main

import (
	"log/slog"
	_ "net"
	_ "net/http"
	"os"
	_ "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/rs/zerolog/log"
	"github.com/sergeyseliverstovv/task_server/internal/app"
	"github.com/sergeyseliverstovv/task_server/internal/config"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/reflection"
)

func main() {
	//ctx := context.Background()
	// TODO: инициалищировать конфиг
	cfg := config.MustLoad()

	// TODO: инициализация логгера
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	log.Info("start serves", slog.Any("env", cfg)) // Потом убрать

	// TODO: инициализация приложения
	app.New()

	//application := app.New(log, cfg.GRPC.Port, cfg.StorageUrl)
	//runRest(ctx)
	//go application.GRPCSrv.MustRun()

	//log.Info("Starting wwwww")

	//wg := &sync.WaitGroup{}
	//wg.Add(1)

	//TODO: запустить gRPC сервер
	/*go func() {
		defer wg.Done()
		application.GRPCSrv.MustRun()
	}()*/

	// TODO: запустить REST сервер
	/*go func() {
		defer wg.Done()
		runRest()
	}()

	wg.Wait()*/
	//application.GRPCSrv.MustRun()
	//TODO: запусить REST сервер

}
