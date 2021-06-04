// Package greetings sends greetings.
package greetings

import (
	"math/rand"
)

var greetings = []string{
	"Hello",
	"Bonjour",
	"Marhabaan",
}

// Greeting returns a random greeting.
func Greeting() string {
	return greetings[rand.Intn(len(greetings))]
}
