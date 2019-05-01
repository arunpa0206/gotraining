// main declares the CLI that spins up the server of
// our API.
// It takes some arguments, validates if they're valid
// and match the expected type and then intiialize the
// server.
package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

// cliArgs defines the configuration that the CLI
// expects. By using a struct we can very easily
// aggregate them into an object and check what are
// the expected types.
// If we need to mock this later it's just a matter
// of reusing the struct.
type cliArgs struct {
	Port int `arg:"-p,help:port to listen to"`
}

var (
	// args is a reference to an instantiation of
	// the configuration that the CLI expects but
	// with some values set.
	// By setting some values in advance we provide
	// default values that the user might provide
	// or not.
	args = &cliArgs{
		Port: 8080,
	}
)

// main parses the arguments from the CLI as specified
// by our configuration described in `cliArgs` and then
// populates the `args` reference we defined in the `vars`
// section above.
func main() {
	arg.MustParse(args)
	fmt.Printf("port=%d\n", args.Port)
}
