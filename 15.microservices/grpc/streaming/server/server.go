package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	api "github.com/arunpa0206/gotraining/15.microservices/grpc/streaming/api/generated"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) SayHelloUnary(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	fmt.Printf("SayHelloUnary function was invoked with %v\n", req)
	firstName := req.GetHello().GetFirstName()
	LastName := req.GetHello().GetLastName()
	result := "Hello " + firstName + LastName + "!"
	res := &api.HelloResponse{
		Result: result,
	}
	return res, nil
}

func (*server) SayHelloClientStreaming(stream api.HelloService_SayHelloClientStreamingServer) error {
	fmt.Printf("SayHelloClientStreaming function was invoked : \n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&api.HelloResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.GetHello().GetFirstName()
		result += "Hello " + firstName + "! "
	}

}

func (*server) SayHelloServerStreaming(req *api.HelloRequest, stream api.HelloService_SayHelloServerStreamingServer) error {
	fmt.Printf("SayHelloServerStreaming function was invoked with %v\n", req)
	firstName := req.GetHello().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &api.HelloResponseMultipleTimes{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			log.Fatalf("Error while sending to client stream: %v", err)
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) SayHelloBidirectionalStreaming(stream api.HelloService_SayHelloBidirectionalStreamingServer) error {
	fmt.Printf("SayHelloBidirectionalStreaming function was invoked \n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetHello().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&api.HelloResponseMultipleTimes{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}

}

func main() {

	fmt.Println("Hello World!!!")

	lis, err := net.Listen("tcp", "0.0.0.0:5051")

	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to server: %v", err)
	}
}
