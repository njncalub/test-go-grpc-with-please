package greetings

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestGreeting(t *testing.T) {
	assert.NotEqual(t, Greeting(), "")
}
