package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jisuhan3201/multiplexer/multiplexerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) GLog(ctx context.Context, req *multiplexerpb.GLogRequest) (*multiplexerpb.GLogResponse, error) {
	fmt.Printf("received GLog rpc %v", req)
	result := req.GetId() + ":" + req.GetType()
	return &multiplexerpb.GLogResponse{
		Result: result,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("cannot listen : %v", err)
	}

	s := grpc.NewServer()
	multiplexerpb.RegisterGLogServierServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("cannot server : %v", err)
	}
}
