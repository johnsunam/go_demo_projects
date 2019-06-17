package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest a simple function
func HandleRequest() (string, error) {
	return fmt.Sprintf("Hello world"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
