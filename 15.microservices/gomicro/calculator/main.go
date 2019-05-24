package main

import (
	"fmt"
	"os"

	"context"

	proto "github.com/arunpa0206/gotrainingv1/15.microservices/gomicro/calculator/proto"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

/*
Example usage of top level service initialisation
*/

type Calculate struct{}

func (g *Calculate) Multiply(ctx context.Context, req *proto.CalculateRequest, rsp *proto.CalculateResponse) error {

	rsp.Result = req.A * req.B
	return nil
}

// Setup and the client
func runClient(service micro.Service) {

	// Create new greeter client
	calculate := proto.NewCalculateService("calculate", service.Client())

	// Call the greeter
	rsp, err := calculate.Multiply(context.TODO(), &proto.CalculateRequest{A: 8, B: 2})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Result)
}

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("calculate"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// Setup some flags. Specify --run_client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	proto.RegisterCalculateHandler(service.Server(), new(Calculate))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
