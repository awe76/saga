package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	api "github.com/awe76/saga/api/sagatransactionapis/v1"
	engine "github.com/awe76/saga/tracer/sagatracerapis/engine"
	tracerapi "github.com/awe76/saga/tracer/sagatracerapis/v1"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC tracer server endpoint
	grpcTracerServerEndpoint = flag.String("grpc-tracer-server-endpoint", ":50057", "gRPC tracer server endpoint")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", *grpcTracerServerEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", *grpcTracerServerEndpoint, err)
	}

	server := grpc.NewServer()

	service := &sagaTracerServiceServer{}

	tracerapi.RegisterSagaTracerServiceServer(server, service)
	log.Println("Listening on", *grpcTracerServerEndpoint)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type sagaTracerServiceServer struct {
	tracerapi.UnimplementedSagaTracerServiceServer
}

func (s *sagaTracerServiceServer) TraceWorkflow(req *tracerapi.TraceWorkflowRequest, stream tracerapi.SagaTracerService_TraceWorkflowServer) error {

	spawnOperation := func(op *api.Operation) error {
		response := tracerapi.TraceWorkflowResponse{
			WorkflowId: req.State.Id,
			IsRollback: req.State.IsRollback,
			Action: &tracerapi.TraceWorkflowResponse_SpawnOperation{
				SpawnOperation: &tracerapi.SpawnOperation{
					Operation: op,
				},
			},
		}
		return stream.Send(&response)
	}

	endWorkflow := func() error {
		response := tracerapi.TraceWorkflowResponse{
			WorkflowId: req.State.Id,
			IsRollback: req.State.IsRollback,
			Action: &tracerapi.TraceWorkflowResponse_CompleteWorkflow{
				CompleteWorkflow: &tracerapi.CompleteWorkflow{},
			},
		}
		return stream.Send(&response)
	}

	if req.State.IsRollback {
		t := engine.CreateReverseTracer(req.Workflow, req.State, endWorkflow, spawnOperation)
		t.ResolveWorkflow(req.Workflow.End)
	} else {
		t := engine.CreateDirectTracer(req.Workflow, req.State, endWorkflow, spawnOperation)
		t.ResolveWorkflow(req.Workflow.Start)
	}
	return nil
}
