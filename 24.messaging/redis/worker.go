package main

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

func handleTask(value string) {
    sum := 0
    values := strings.Split(value, ",")
    for i := range values {
        v, _ := strconv.Atoi(values[i])
        sum += v
    }
    fmt.Println("Sum of ", value, " is: ", strconv.Itoa(sum))
}

func main() {
    TaskKey := "task"
    broker := New()
    for {
        time.Sleep(time.Second)
        taskValue := broker.receiveTask(TaskKey)
        if len(taskValue) > 0 {
            go handleTask(taskValue)
        }
    }
}