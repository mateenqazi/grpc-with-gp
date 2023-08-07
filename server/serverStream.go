package main

import (
	"log"
	"time"

	"github.com/mateenqazi/grpc-with-go/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *proto.NamesList, stream proto.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, value := range req.Names {
		res := &proto.HelloResponse{
			Name: "Hello " + value,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}
