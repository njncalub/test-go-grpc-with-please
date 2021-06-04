// Package main runs the test server.
package main

import (
	"fmt"

	"github.com/njncalub/test-go-grpc-with-please/src/greetings"
)

func main() {
	fmt.Printf("%s, World!\n", greetings.Greeting())
}
