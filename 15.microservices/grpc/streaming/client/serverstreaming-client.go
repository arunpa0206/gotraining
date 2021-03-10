package main

import (
	"context"
	"fmt"
	"io"
	"log"

	api "github.com/arunpa0206/gotraining/15.microservices/grpc/streaming/api/generated"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}
	defer cc.Close()

	c := api.NewHelloServiceClient(cc)

	checkServerStreaming(c)

}

func checkServerStreaming(c api.HelloServiceClient) {
	fmt.Println("Starting a Server Streaming RPC...")
	req := &api.HelloRequest{
		Hello: &api.Hello{FirstName: "Jobet",
			LastName: " Samuel",
		},
	}
	res, err := c.SayHelloServerStreaming(context.Background(), req)

	if err != nil {
		log.Fatal("error while calling Hello RPC: %v", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			// Reached end of stream
			break
		}
		if err != nil {
			log.Fatal("Error while reading froms stream: %v", err)
		}
		log.Printf(" Response from Server Streaming RPC: %v", msg)
	}
}
