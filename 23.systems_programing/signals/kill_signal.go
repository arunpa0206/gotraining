package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	go func() {
		for {
			fmt.Println("Doing Work")
			time.Sleep(1 * time.Second)
		}
	}()
	<-killSignal
	fmt.Println("Thanks for using Golang!")
}