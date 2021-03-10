package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	checkClientStreaming(c)
	//checkBidirectionalStreaming(c)
}

func checkClientStreaming(c api.HelloServiceClient) {
	fmt.Println("Starting a Client Streaming RPC...")

	requests := []*api.HelloRequestMultipleTimes{
		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Jobet",
			},
		},
		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Mark",
			},
		},

		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "John",
			},
		},

		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Matthew",
			},
		},
	}

	stream, err := c.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatal("Error while calling  HelloClient: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while receiving response from Server: %v", err)
	}
	fmt.Printf("Response received from Client Streaming Server: %v\n", res)

}
