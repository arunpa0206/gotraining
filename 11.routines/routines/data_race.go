package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	// Initialize a waitgroup variable
	var wg sync.WaitGroup
	// `Add(1) signifies that there is 1 task that we need to wait for
	wg.Add(1)
	go func() {
		i = 5
		// Calling `wg.Done` indicates that we are done with the task we are waiting fo
		wg.Done()
	}()
	// `wg.Wait` blocks until `wg.Done` is called the same number of times
	// as the amount of tasks we have (in this case, 1 time)
	wg.Wait()
	return i
}

