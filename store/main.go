package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	store "github.com/awe76/saga/store/storeapis/v1"
	clientv3 "go.etcd.io/etcd/client/v3"

	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := ":50055"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()

	service, err := NewStoreServiceServer()

	if err != nil {
		return fmt.Errorf("failed to create gRPC server: %w", err)
	}

	defer service.Close()

	store.RegisterStoreServiceServer(server, service)
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type storeServiceServer struct {
	store.UnimplementedStoreServiceServer
	cli *clientv3.Client
}

func NewStoreServiceServer() (*storeServiceServer, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{":2379", ":22379", ":32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &storeServiceServer{
		cli: cli,
	}, nil
}

func (s *storeServiceServer) Close() {
	s.cli.Close()
}

func (s *storeServiceServer) Put(ctx context.Context, req *store.PutRequest) (*store.PutResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	_, err := s.cli.Put(ctx, req.Key, req.Value)
	cancel()

	if err != nil {
		return nil, err
	}

	return &store.PutResponse{}, nil
}

func (s *storeServiceServer) Get(ctx context.Context, req *store.GetRequest) (*store.GetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	resp, err := s.cli.Get(ctx, req.Key)
	cancel()

	if err != nil {
		return nil, err
	}

	var values []string

	for _, kv := range resp.Kvs {
		values = append(values, string(kv.Value))
	}

	return &store.GetResponse{
		Values: values,
	}, nil
}
