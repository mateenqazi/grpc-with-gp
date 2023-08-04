package main

import (
	"log"

	"github.com/mateenqazi/grpc-with-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	names := &proto.NamesList{
		Names: []string{"Mateen", "Qazi", "Alic", "Bob"},
	}

	//CallSayHello(client)

	// CallSayHelloServerStream(client, names)

	// CallSayHelloClientStream(client, names)

	SayHelloBidirectionStreaming(client, names)

}
