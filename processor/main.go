package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	processor "github.com/awe76/saga/processor/workflowapis/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:50051"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	processor.RegisterSagaWorkflowServiceServer(server, &sagaWorkflowServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type sagaWorkflowServiceServer struct {
	processor.UnimplementedSagaWorkflowServiceServer
}

// Echo sends input value backward.
func (s *sagaWorkflowServiceServer) Echo(ctx context.Context, req *processor.EchoRequest) (*processor.EchoResponse, error) {

	value := req.GetValue()
	log.Println("Echo: ", value)

	return &processor.EchoResponse{Value: value}, nil
}
