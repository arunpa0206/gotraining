package main

import (
	"context"
	"fmt"
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
	checkUnary(c)
	//checkServerStreaming(c)
	//checkClientStreaming(c)
	//checkBidirectionalStreaming(c)
}

func checkUnary(c api.HelloServiceClient) {
	fmt.Println("Starting a Unary RPC...")
	req := &api.HelloRequest{
		Hello: &api.Hello{FirstName: "Jobet",
			LastName: " Samuel",
		},
	}
	res, err := c.SayHelloUnary(context.Background(), req)

	if err != nil {
		log.Fatal("error while calling Hello RPC: %v", err)
	}

	log.Printf("Respone from Unary Server: %v", res.Result)
}
