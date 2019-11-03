package main

import (
    "github.com/aws/aws-lambda-go/lambda"
)

type book struct {
    ISBN   string `json:"isbn"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

func show() (*book, error) {
    bk := &book{
        ISBN:   "978-1420931693",
        Title:  "The Republic",
        Author: "Plato",
    }

    return bk, nil
}

func main() {
    lambda.Start(show)
}