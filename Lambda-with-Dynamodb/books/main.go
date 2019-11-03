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
    // Fetch a specific book record from the DynamoDB database. We'll
    // make this more dynamic in the next section.
    bk, err := getItem("978-0486298238")
    if err != nil {
        return nil, err
    }

    return bk, nil
}

func main() {
    lambda.Start(show)
}