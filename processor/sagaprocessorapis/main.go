package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	handler "github.com/awe76/saga/handler/sagahandlerapis/v1"
	processor "github.com/awe76/saga/processor/sagaprocessorapis/v1"
	state "github.com/awe76/saga/state/sagastateapis/v1"
	tracer "github.com/awe76/saga/tracer/sagatracerapis/v1"
)

var (
	// command-line options:
	// gRPC state server endpoint
	grpcStateServerEndpoint = flag.String("grpc-state-server-endpoint", "state:50056", "gRPC state server endpoint")
	// gRPC tracer server endpoint
	grpcTracerServerEndpoint = flag.String("grpc-tracer-server-endpoint", "tracer:50057", "gRPC tracer server endpoint")
	// gRPC processor server endpoint
	grpcProcessorServerEndpoint = flag.String("grpc-processor-server-endpoint", ":50058", "gRPC processor server endpoint")
	// gRPC handler server endpoint
	grpcHandlerServerEndpoint = flag.String("grpc-handler-server-endpoint", "handler:50059", "gRPC handler server endpoint")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", *grpcProcessorServerEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", *grpcProcessorServerEndpoint, err)
	}

	server := grpc.NewServer()
	service, err := NewProcessorServiceServer()

	if err != nil {
		return fmt.Errorf("failed to create gRPC server: %w", err)
	}

	processor.RegisterSagaProcessorServiceServer(server, service)
	log.Println("Listening on", *grpcProcessorServerEndpoint)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type sagaProcessorServiceServer struct {
	processor.UnimplementedSagaProcessorServiceServer
	stateClient   state.SagaStateServiceClient
	tracerClient  tracer.SagaTracerServiceClient
	handlerClient handler.SagaHandlerServiceClient
}

func NewProcessorServiceServer() (*sagaProcessorServiceServer, error) {
	stateConn, err := grpc.Dial(*grpcStateServerEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect to state: %v", err)
	}

	stateClient := state.NewSagaStateServiceClient(stateConn)

	tracerConn, err := grpc.Dial(*grpcTracerServerEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect to tracer: %v", err)
	}

	tracerClient := tracer.NewSagaTracerServiceClient(tracerConn)

	handlerConn, err := grpc.Dial(*grpcHandlerServerEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect to handler: %v", err)
	}

	handlerClient := handler.NewSagaHandlerServiceClient(handlerConn)

	return &sagaProcessorServiceServer{
		stateClient:   stateClient,
		tracerClient:  tracerClient,
		handlerClient: handlerClient,
	}, nil
}

type executionResult struct {
	value *processor.ExecuteWorkflowResponse
	err   error
}

func (s *sagaProcessorServiceServer) ExecuteWorkflow(ctx context.Context, req *processor.ExecuteWorkflowRequest) (*processor.ExecuteWorkflowResponse, error) {
	initResp, err := s.stateClient.Init(ctx, &state.InitRequest{
		Start:   req.Workflow.Start,
		Payload: req.Workflow.Payload,
	})

	if err != nil {
		return nil, fmt.Errorf("did not init workflow state: %v", err)
	}

	currentState := initResp.State
	workflowId := currentState.Id
	workflow := req.Workflow

	fmt.Printf("%s (%s) is started\n", workflow.Name, workflowId)

	for {
		tracerStream, err := s.tracerClient.TraceWorkflow(ctx, &tracer.TraceWorkflowRequest{
			Workflow: workflow,
			State:    currentState,
		})

		if err != nil {
			return nil, fmt.Errorf("did not trace workflow: %v", err)
		}

		for {
			cmd, err := tracerStream.Recv()
			if err == io.EOF {
				break
			}

			spawnOperation := cmd.GetSpawnOperation()
			if spawnOperation != nil {
				_, err := s.stateClient.SpawnOperation(ctx, &state.SpawnOperationRequest{
					Id:         workflowId,
					IsRollback: cmd.IsRollback,
					Operation:  spawnOperation.Operation,
				})

				if err != nil {
					return nil, fmt.Errorf("did not update state with spawned operation: %v", err)
				}

				resp, err := s.handlerClient.HandleOperation(ctx, &handler.HandleOperationRequest{
					WorkflowId: workflowId,
					IsRollback: cmd.IsRollback,
					Payload:    "",
					Operation:  spawnOperation.Operation,
				})

				if err != nil {
					return nil, fmt.Errorf("did not handle operation: %v", err)
				}

				if resp.IsFailed {
					failedState, err := s.stateClient.FailOperation(ctx, &state.FailOperationRequest{
						Id:        resp.WorkflowId,
						Operation: resp.Operation,
					})

					if err != nil {
						return nil, fmt.Errorf("did not handle failed operation result %v", err)
					}

					currentState = failedState.State
				} else {
					completedState, err := s.stateClient.CompleteOperation(ctx, &state.CompleteOperationRequest{
						IsRollback: resp.IsRollback,
						Id:         resp.WorkflowId,
						Payload:    resp.Payload,
						Operation:  resp.Operation,
					})

					if err != nil {
						return nil, fmt.Errorf("did not handle operation result %v", err)
					}

					currentState = completedState.State
				}
			}

			completeWorkflow := cmd.GetCompleteWorkflow()
			if completeWorkflow != nil {
				finalState, err := s.stateClient.EndWorkflow(ctx, &state.EndWorkflowRequest{Id: workflowId})

				if err != nil {
					return nil, fmt.Errorf("did not update state with completed workflow: %v", err)
				}

				return &processor.ExecuteWorkflowResponse{
					State: finalState.State,
				}, nil
			}
		}
	}
}
