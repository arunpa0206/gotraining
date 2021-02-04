// Golang program to read and write the files
package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	// fmt package implements formatted
	// I/O and has functions like Printf
	// and Scanf
	fmt.Printf("Writing to a file in Go lang\n")

	// in case an error is thrown it is received
	// by the err variable and Fatalf method of
	// log prints the error message and stops
	// program execution
	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Defer is used for purposes of cleanup like
	// closing a running file after the file has
	// been written and main //function has
	// completed execution
	defer file.Close()

	// len variable captures the length
	// of the string written to the file.
	len, err := file.WriteString("Welcome to Golang." +
		" This program demonstrates reading and writing" +
		" operations to a file in Go lang.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	// Name() method returns the name of the
	// file as presented to Create() method.
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {

	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "test.txt"

	// The ioutil package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	CreateFile()
	ReadFile()
}
