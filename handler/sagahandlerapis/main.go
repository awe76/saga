package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	handler "github.com/awe76/saga/handler/sagahandlerapis/v1"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC handler server endpoint
	grpcHandlerServerEndpoint = flag.String("grpc-handler-server-endpoint", ":50059", "gRPC handler server endpoint")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", *grpcHandlerServerEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", *grpcHandlerServerEndpoint, err)
	}

	server := grpc.NewServer()
	service := &handlerServiceServer{}

	handler.RegisterSagaHandlerServiceServer(server, service)
	log.Println("Listening on", *grpcHandlerServerEndpoint)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type handlerServiceServer struct {
	handler.UnimplementedSagaHandlerServiceServer
}

func (h *handlerServiceServer) HandleOperation(ctx context.Context, req *handler.HandleOperationRequest) (*handler.HandleOperationResponse, error) {
	resp := handler.HandleOperationResponse{
		IsRollback: req.IsRollback,
		WorkflowId: req.WorkflowId,
		Operation:  req.Operation,
	}

	if req.IsRollback {
		fmt.Printf("%s operation rollback is started\n", req.Operation.Name)
	} else {
		fmt.Printf("%s operation is started\n", req.Operation.Name)
	}

	rand.Seed(time.Now().UnixNano())
	pause := rand.Intn(100)

	// sleep for some random time
	time.Sleep(time.Duration(pause) * time.Millisecond)

	payload, err := json.Marshal(rand.Float32())
	if err != nil {
		return nil, err
	}

	resp.IsFailed = !req.IsRollback && rand.Float32() > 0.8
	resp.Payload = string(payload)

	return &resp, nil
}
