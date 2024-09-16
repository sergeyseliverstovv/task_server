package httpserv

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	api "github.com/sergeyseliverstovv/api/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunRest(ctx context.Context) {

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := api.RegisterTaskManagerHandlerFromEndpoint(ctx, mux, "localhost:1414", opts)
	if err != nil {
		panic(err)
	}
	//log := log.With(slog.String("op", op))

	//log.Info("Starting REST service")
	log.Printf("Starting REST service")

	err = http.ListenAndServe("localhost:1414", mux)
	if err != nil {
		panic(err)
	}
	log.Printf("DKFKKFKFKFKFK")
}
