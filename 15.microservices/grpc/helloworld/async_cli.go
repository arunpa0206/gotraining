package main

import (
	"log"
	// "os"

	// "strconv"
	pb "github.com/arunpa0206/gotraining/15.microservices/grpc/helloworld/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	numOfChan   = 6
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	/* name := defaultName
	   if len(os.Args) > 1 {
	       name = os.Args[1]
	   } */

	// callMultipleAsyncRPC(c)
	callMultipleAsyncRPCWithRet(c)
	for {
	}
}

var cnt int = 0

func callMultipleAsyncRPC(c pb.GreeterClient) {
	log.Print("Calling RPC without getting return value but with goroutine")
	for i := 0; i < 5; i++ {
		go func() {
			log.Print("Inside goroutine: Start c.SyaHello", cnt)
			cnt++
			c.SayHello(context.Background(), &pb.HelloRequest{Name: "may"})
			cnt++
			log.Print("Inside goroutine: Exit c.SayHello", cnt)
		}()
	}
	log.Print("Bless with goroutine")
}

func callMultipleAsyncRPCWithRet(c pb.GreeterClient) {
	log.Print("Calling RPC with getting return value but and goroutine")
	var channels [numOfChan]chan string
	hello_strings := [numOfChan]string{"async request", "hej!", "hello", "golang", "你好！"}

	for i := range channels {
		channels[i] = make(chan string)
		log.Print(hello_strings[i])
	}

	for i := 0; i < numOfChan; i++ {
		go func(hello string, ch chan string) {
			log.Print("Inside goroutine: Start c.SyaHello", hello)
			r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: hello})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}

			ch <- hello
			log.Print("Inside goroutine: Exit c.SayHello", r.Message)
		}(hello_strings[i], channels[i])
	}

	var res_hellos [numOfChan]string

	for i := 0; i < numOfChan; i++ {
		log.Print("Let me wait for goroutines")
		res_hellos[i] = <-channels[i]
		log.Print("Got: ", res_hellos[i])
	}

	log.Print("Bless with goroutine")
}
