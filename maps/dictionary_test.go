package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	definition := "this is just a test"
	dictionary := Dictionary{"test": definition}

	t.Run("should search for a word", func(t *testing.T) {

		actual, err := dictionary.Search("test")
		expected := definition

		assertNoError(t, err)

		assertEqualStrings(t, expected, actual)
	})

	t.Run("should return error if word was not found", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, DefinitionNotFound, err)
	})

	t.Run("should add word to dictionary", func(t *testing.T) {
		_, err := dictionary.Search("new")
		assertError(t, DefinitionNotFound, err)

		expected := "something that did not exist before"
		_ = dictionary.Add("new", expected)

		definition, searchErr := dictionary.Search("new")

		assertNoError(t, searchErr)
		assertEqualStrings(t, expected, definition)
	})

	t.Run("should not overwrite existing word on Add", func(t *testing.T) {
		expected := "something that did not exist before"
		err := dictionary.Add("new", expected)

		assertError(t, DefinitionExists, err)
	})

	t.Run("should update definition", func(t *testing.T) {
		expected := "a predicate"
		err := dictionary.Update("test", expected)

		definition, _ := dictionary.Search("test")
		assertNoError(t, err)
		assertEqualStrings(t, expected, definition)

		newError := dictionary.Update("unknown", "not known")
		assertError(t, DefinitionNotFound, newError)
	})

	t.Run("should delete definition", func(t *testing.T) {
		dictionary.Delete("test")

		_, err := dictionary.Search("test")
		assertError(t, DefinitionNotFound, err)
	})
}

func assertEqualStrings(t *testing.T, expected, actual string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected %q, but got: %q", expected, actual)
	}
}

func assertError(t *testing.T, expected, actual error) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected error %q, but got error %q", expected, actual)
		return
	}
	if actual == nil {
		if expected != nil {
			t.Errorf("Expected an error")
		}
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Error occurred but was not expected: %q", err)
	}
}
