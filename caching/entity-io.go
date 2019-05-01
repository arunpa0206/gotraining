package main

import (
	"fmt"
	// Import the Radix.v2 redis package.
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	// Establish a connection to the Redis server listening on port 6379 of the
	// local machine. 6379 is the default port, so unless you've already
	// changed the Redis configuration file this should work.
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	// Importantly, use defer to ensure the connection is always properly
	// closed before exiting the main() function.
	defer conn.Close()
	//add(conn,"entities:person:celebrety:name:Amitabh","233")
	// Send our command across the connection. The first parameter to Cmd()
	// is always the name of the Redis command (in this example HMSET),
	// optionally followed by any necessary arguments (in this example the
	// key, followed by the various hash fields and values).
	resp := conn.Cmd("SADD", "entities:person:celebrety:name:Amitabh", "222")
	// Check the Err field of the *Resp object for any errors.
	if resp.Err != nil {
		log.Fatal(err)
	}

	fmt.Println("Entity added!")

	r := conn.Cmd("SMEMBERS", "entities:person:celebrety:name:Amitabh")
	if r.Err != nil {
		log.Fatal(err)
	}
	// This:
	l, _ := r.List()
	fmt.Println(l)
	for _, elemStr := range l {
		fmt.Println(elemStr)
	}
}
