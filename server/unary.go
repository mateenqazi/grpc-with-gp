package main

import (
	"context"

	"github.com/mateenqazi/grpc-with-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Name: "Hello",
	}, nil
}
