package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/jsonpb"
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
	defer service.Close()

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

func getKey(op *api.Operation, isRollback bool) string {
	return fmt.Sprintf("%s:%s:%s:%v", op.Name, op.From, op.To, isRollback)
}

func addOp(m map[string]*api.Operation, op *api.Operation, isRollback bool) {
	key := getKey(op, isRollback)
	m[key] = op
}

func removeOp(m map[string]*api.Operation, op *api.Operation, isRollback bool) {
	key := getKey(op, isRollback)
	delete(m, key)
}

type sagaStateServiceServer struct {
	state.UnimplementedSagaStateServiceServer
	client store.StoreServiceClient
	conn   *grpc.ClientConn
}

func NewStateServiceServer() (*sagaStateServiceServer, error) {
	conn, err := grpc.Dial(*grpcStoreServerEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := store.NewStoreServiceClient(conn)
	return &sagaStateServiceServer{
		client: client,
		conn:   conn,
	}, nil
}

func (s *sagaStateServiceServer) getKey(id string) string {
	return fmt.Sprintf("workflow:state:%v", id)
}

func (s *sagaStateServiceServer) Close() {
	s.conn.Close()
}

func (s *sagaStateServiceServer) update(ctx context.Context, id string, update func(*api.State)) (*api.State, error) {
	key := s.getKey(id)
	data, err := s.client.Get(ctx, &store.GetRequest{Key: key})

	state := &api.State{}

	if err != nil {
		return state, err
	}

	err = jsonpb.UnmarshalString(data.Values[0], state)
	if err != nil {
		return state, err
	}

	update(state)

	marshaler := jsonpb.Marshaler{}
	value, err := marshaler.MarshalToString(state)

	_, err = s.client.Put(ctx, &store.PutRequest{Key: s.getKey(id), Value: value})
	return state, nil
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

	marshaler := jsonpb.Marshaler{}
	value, err := marshaler.MarshalToString(result)

	if err != nil {
		return nil, err
	}

	_, err = s.client.Put(ctx, &store.PutRequest{Key: s.getKey(id), Value: value})

	if err != nil {
		return nil, err
	}

	resp := &state.InitResponse{
		State: result,
	}

	return resp, nil
}

func (s *sagaStateServiceServer) CompleteOperation(ctx context.Context, req *state.CompleteOperationRequest) (*state.CompleteOperationResponse, error) {
	result, err := s.update(ctx, req.Id, func(t *api.State) {
		removeOp(t.InProgress, req.Operation, req.IsRollback)
		addOp(t.Done, req.Operation, req.IsRollback)

		ops, found := t.Data[req.Operation.To]
		if !found {
			values := make(map[string]string)
			ops = &api.Values{
				Values: values,
			}

			t.Data[req.Operation.To] = ops
			values[req.Operation.Name] = req.Payload
		}
	})

	if err != nil {
		return nil, err
	}

	return &state.CompleteOperationResponse{State: result}, nil
}

func (s *sagaStateServiceServer) FailOperation(ctx context.Context, req *state.FailOperationRequest) (*state.FailOperationResponse, error) {
	result, err := s.update(ctx, req.Id, func(t *api.State) {
		removeOp(t.InProgress, req.Operation, false)
		t.IsRollback = true
	})

	if err != nil {
		return nil, err
	}

	return &state.FailOperationResponse{State: result}, nil
}

func (s *sagaStateServiceServer) SpawnOperation(ctx context.Context, req *state.SpawnOperationRequest) (*state.SpawnOperationResponse, error) {
	result, err := s.update(ctx, req.Id, func(t *api.State) {
		addOp(t.InProgress, req.Operation, t.IsRollback)
	})

	if err != nil {
		return nil, err
	}

	return &state.SpawnOperationResponse{State: result}, nil
}

func (s *sagaStateServiceServer) EndWorkflow(ctx context.Context, req *state.EndWorkflowRequest) (*state.EndWorkflowResponse, error) {
	result, err := s.update(ctx, req.Id, func(t *api.State) {
		t.Completed = true
	})

	if err != nil {
		return nil, err
	}

	return &state.EndWorkflowResponse{State: result}, nil
}
