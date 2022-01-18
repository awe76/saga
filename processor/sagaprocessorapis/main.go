package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	broker "github.com/awe76/saga/api/sagabrokerapis/v1"
	api "github.com/awe76/saga/api/sagatransactionapis/v1"
	processor "github.com/awe76/saga/processor/sagaprocessorapis/v1"
	state "github.com/awe76/saga/state/sagastateapis/v1"
	tracer "github.com/awe76/saga/tracer/sagatracerapis/v1"
	nats "github.com/nats-io/nats.go"
)

var (
	// command-line options:
	// gRPC processor server endpoint
	grpcProcessorServerEndpoint = flag.String("grpc-processor-server-endpoint", ":50058", "gRPC processor server endpoint")
	// gRPC state server endpoint
	grpcStateServerEndpoint = flag.String("grpc-state-server-endpoint", "state:50056", "gRPC state server endpoint")
	// gRPC tracer server endpoint
	grpcTracerServerEndpoint = flag.String("grpc-tracer-server-endpoint", "tracer:50056", "gRPC tracer server endpoint")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func closeBrokerConnection(service *sagaProcessorServiceServer) {
	if service != nil {
		service.CloseBrokerConnection()
	}
}

func run() error {
	listener, err := net.Listen("tcp", *grpcProcessorServerEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", *grpcProcessorServerEndpoint, err)
	}

	server := grpc.NewServer()
	service, err := NewProcessorServiceServer()

	defer closeBrokerConnection(service)

	if err != nil {
		return fmt.Errorf("failed to create gRPC server: %w", err)
	}

	go service.HandleOperationResults()

	processor.RegisterSagaProcessorServiceServer(server, service)
	log.Println("Listening on", *grpcProcessorServerEndpoint)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type sagaProcessorServiceServer struct {
	processor.UnimplementedSagaProcessorServiceServer
	stateClient  state.SagaStateServiceClient
	tracerClient tracer.SagaTracerServiceClient
	ec           *nats.EncodedConn
	sender       chan *broker.OperationPayload
	receiver     chan *broker.OperationPayload
	handlers     map[string]chan *broker.OperationPayload
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

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, fmt.Errorf("did not connect to nats: %v", err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		return nil, fmt.Errorf("did not connect to nats json channel: %v", err)
	}

	receiver := make(chan *broker.OperationPayload)
	ec.BindRecvChan("rop", receiver)

	sender := make(chan *broker.OperationPayload)
	ec.BindSendChan("sop", sender)

	handlers := make(map[string]chan *broker.OperationPayload)

	return &sagaProcessorServiceServer{
		stateClient:  stateClient,
		tracerClient: tracerClient,
		ec:           ec,
		sender:       sender,
		receiver:     receiver,
		handlers:     handlers,
	}, nil
}

func (s *sagaProcessorServiceServer) RegisterHandler(id string, handler chan *broker.OperationPayload) {
	s.handlers[id] = handler
}

func (s *sagaProcessorServiceServer) UnregisterHandler(id string) {
	delete(s.handlers, id)
}

func (s *sagaProcessorServiceServer) HandleOperationResults() {
	for {
		select {
		case req := <-s.receiver:
			if handler, found := s.handlers[req.WorkflowId]; found {
				handler <- req
			}
		}
	}
}

func (s *sagaProcessorServiceServer) CloseBrokerConnection() {
	if s.ec != nil {
		s.ec.Close()
	}
}

func (s *sagaProcessorServiceServer) ExecuteWorkflow(ctx context.Context, req *processor.ExecuteWorkflowRequest) (*processor.ExecuteWorkflowResponse, error) {

	initialState, err := s.stateClient.Init(ctx, &state.InitRequest{
		Start:   req.Workflow.Start,
		Payload: req.Payload,
	})

	if err != nil {
		return nil, fmt.Errorf("did not init workflow state: %v", err)
	}

	tracerCh := make(chan *tracer.TraceWorkflowRequest)
	tracerCh <- &tracer.TraceWorkflowRequest{
		Workflow: req.Workflow,
		State:    initialState.State,
	}

	workflowId := initialState.State.Id

	handlerCh := make(chan *broker.OperationPayload)
	s.RegisterHandler(workflowId, handlerCh)

	defer s.UnregisterHandler(workflowId)

	for {
		select {
		case tracerReq := <-tracerCh:
			tracerStream, err := s.tracerClient.TraceWorkflow(ctx, tracerReq)

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

					s.sender <- &broker.OperationPayload{
						WorkflowName: req.Workflow.Name,
						WorkflowId:   workflowId,
						IsRollback:   cmd.IsRollback,
						Name:         spawnOperation.Operation.Name,
						From:         spawnOperation.Operation.From,
						To:           spawnOperation.Operation.To,
						Payload:      "",
					}
				}

				completeWorkflow := cmd.GetCompleteWorkflow()
				if completeWorkflow != nil {
					state, err := s.stateClient.EndWorkflow(ctx, &state.EndWorkflowRequest{Id: workflowId})

					if err != nil {
						return nil, fmt.Errorf("did not update state with completed workflow: %v", err)
					}

					return &processor.ExecuteWorkflowResponse{
						State: state.State,
					}, nil
				}
			}

		case resp := <-handlerCh:
			var currentState *api.State

			op := &api.Operation{
				Name: resp.Name,
				From: resp.From,
				To:   resp.To,
			}

			if resp.IsFailed {
				failedResp, err := s.stateClient.FailOperation(ctx, &state.FailOperationRequest{
					Id:        resp.WorkflowId,
					Operation: op,
				})

				if err != nil {
					return nil, fmt.Errorf("did not handle failed operation result %v", err)
				}

				currentState = failedResp.State

			} else {
				completeResp, err := s.stateClient.CompleteOperation(ctx, &state.CompleteOperationRequest{
					IsRollback: resp.IsRollback,
					Id:         resp.WorkflowId,
					Payload:    resp.Payload,
					Operation:  op,
				})

				if err != nil {
					return nil, fmt.Errorf("did not handle operation result %v", err)
				}

				currentState = completeResp.State
			}

			tracerCh <- &tracer.TraceWorkflowRequest{
				Workflow: req.Workflow,
				State:    currentState,
			}
		}
	}
}
