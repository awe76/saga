package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/awe76/saga/gateway/sagagatewayapis/v1" // Update
)

var (
	// command-line options:
	// gRPC state server endpoint
	grpcStateServerEndpoint = flag.String("grpc-state-server-endpoint", "state:50056", "gRPC state server endpoint")
	// gRPC gateway endpoint
	grpcGatewayEndpoint = flag.String("grpc-gateway-endpoint", ":8081", "gRPC gateway endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterSagaStateServiceHandlerFromEndpoint(ctx, mux, *grpcStateServerEndpoint, opts)
	if err != nil {
		return err
	}

	log.Println("Listening on", *grpcGatewayEndpoint)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(*grpcGatewayEndpoint, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
