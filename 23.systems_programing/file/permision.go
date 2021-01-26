package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    err := os.Mkdir("new", 0755)
    if err != nil {
        log.Fatal(err)
    }
    stats, err := os.Stat("new")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Permission Folder Before: %s\n", stats.Mode())
    err = os.Chmod("new", 0700)
    if err != nil {
        log.Fatal(err)
    }
    stats, err = os.Stat("new")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Permission Folder After:  %s\n", stats.Mode())
}