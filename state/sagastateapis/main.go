package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	api "github.com/awe76/saga/api/sagatransactionapis/v1"
	state "github.com/awe76/saga/state/sagastateapis/v1"
	store "github.com/awe76/saga/store/storeapis/v1"
)

var (
	// command-line options:
	// gRPC store server endpoint
	grpcStoreServerEndpoint = flag.String("grpc-store-server-endpoint", "store:50055", "gRPC store server endpoint")
	// gRPC state server endpoint
	grpcStateServerEndpoint = flag.String("grpc-state-server-endpoint", ":50056", "gRPC state server endpoint")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", *grpcStateServerEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", *grpcStateServerEndpoint, err)
	}

	server := grpc.NewServer()

	service, err := NewStateServiceServer()

	if err != nil {
		return fmt.Errorf("failed to create gRPC server: %w", err)
	}

	state.RegisterSagaStateServiceServer(server, service)
	log.Println("Listening on", *grpcStateServerEndpoint)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type sagaStateServiceServer struct {
	state.UnimplementedSagaStateServiceServer
	client store.StoreServiceClient
}

func NewStateServiceServer() (*sagaStateServiceServer, error) {
	conn, err := grpc.Dial(*grpcStoreServerEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := store.NewStoreServiceClient(conn)
	return &sagaStateServiceServer{
		client: client,
	}, nil
}

func (s *sagaStateServiceServer) getKey(id string) string {
	return fmt.Sprintf("workflow:state:%v", id)
}

func (s *sagaStateServiceServer) Init(ctx context.Context, req *state.InitRequest) (*state.InitResponse, error) {
	id := uuid.New().String()

	result := &api.State{
		Id:         id,
		Completed:  false,
		IsRollback: false,
		Done:       make(map[string]*api.Operation),
		InProgress: make(map[string]*api.Operation),
		Data:       make(map[string]*api.Values),
	}

	values := &api.Values{
		Values: make(map[string]string),
	}

	values.Values["input"] = req.Payload
	result.Data[req.Start] = values

	putctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	value, err := proto.Marshal(result)

	if err != nil {
		return nil, err
	}

	_, err = s.client.Put(putctx, &store.PutRequest{Key: s.getKey(id), Value: string(value)})

	if err != nil {
		return nil, err
	}

	resp := &state.InitResponse{
		State: result,
	}

	return resp, nil
}
