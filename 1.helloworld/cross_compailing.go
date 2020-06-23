package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}

