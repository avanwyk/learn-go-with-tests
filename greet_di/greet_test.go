package greet_di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("should greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := Greet(&buffer, "Chris")

		expected := "Hello, Chris"
		actual := buffer.String()

		if err != nil {
			t.Fatalf("Unexpected error: %q", err.Error())
		}

		if expected != actual {
			t.Errorf("Expected greeting %q, actual greeting %q", expected, actual)
		}
	})
}
