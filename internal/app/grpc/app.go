package grpcserv

import (
	_ "fmt"
	"log"
	_ "log/slog"
	"net"

	api "github.com/sergeyseliverstovv/api/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	api.UnimplementedTaskManagerServer
}

func RunGRPC() error {
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(grpcServer)

	serverApi := &Server{}
	api.RegisterTaskManagerServer(grpcServer, serverApi)

	log.Printf("gRPC server is running on port 8080")

	lis, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		return err
	}

	return grpcServer.Serve(lis)
}

/*type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	handlers.Register(gRPCServer)

	return &App{log: log, gRPCServer: gRPCServer, port: port}
}*/

/*func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc run"

	log := a.log.With(slog.String("op", op))

	log.Info("start gRPC server on port")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server running on", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "stop grpc"

	a.log.With(slog.String("op", op)).Info("stopping grpc")

	a.gRPCServer.GracefulStop() // cервер grpc не остановится пока не будут обработаны соединения
}*/
