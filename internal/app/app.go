package app

import (
	"context"

	_ "log/slog"
	_ "net/http"
	"sync"

	grpcserv "github.com/sergeyseliverstovv/task_server/internal/app/grpc"
	httpserv "github.com/sergeyseliverstovv/task_server/internal/app/http"
)

func New() {
	ctx := context.Background()

	wg := &sync.WaitGroup{}

	// Запускаем gRPC сервер
	wg.Add(2)
	go func() {
		defer wg.Done()
		grpcserv.RunGRPC()
	}()

	// Запускаем REST сервер
	go func() {
		defer wg.Done()
		httpserv.RunRest(ctx)
	}()

	wg.Wait()

}

/*
type App struct {
	GRPCSrv *grpcserv.App
}
*/
// В New еще будет запущен сервер REST, еще нужно будет предать port RESTa
/*func New(log *slog.Logger, gRPCPort int, storage_url string) *App {

	// TODO: init storage

	// TODO: init auth server

	grpcApp := grpcserv.NewApp(log, gRPCPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}*/
