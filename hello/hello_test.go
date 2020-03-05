package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected string, actual string) {
		t.Helper()
		if actual != expected {
			t.Errorf("Expected message %q, actual message %q", expected, actual)
		}
	}

	t.Run("should greet people", func(t *testing.T) {
		name := "Chris"
		actual := Hello(name, "")
		expected := "Hello, " + name
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("should greet world if no name is given", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("should greet in Spanish", func(t *testing.T) {
		actual := Hello("Julia", "Spanish")
		expected := "Hola, Julia"
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("should greet in French", func(t *testing.T) {
		actual := Hello("Andre", "French")
		expected := "Bonjour, Andre"
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("should greet in Afrikaans", func(t *testing.T) {
		actual := Hello("Pieter", "Afrikaans")
		expected := "Goeiedag, Pieter"
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("should run main", func(t *testing.T) {
		main()
	})
}
