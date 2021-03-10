package main

import (
	"context"
	"fmt"
	"io"
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

	checkBidirectionalStreaming(c)
}

func checkBidirectionalStreaming(c api.HelloServiceClient) {
	fmt.Println("Starting a Bidirectional Streaming RPC...")

	stream, err := c.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatal("Error while creating stream: %v", err)
	}

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

	waitc := make(chan struct{})

	go func() {
		//Send the above bunch of requests created
		for _, req := range requests {
			fmt.Printf("Sending requests: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Response Received from Bidirectional Server: %v\n", res.GetResult())
		}

		stream.CloseSend()
	}()
	<-waitc //Block until both are done.
}
