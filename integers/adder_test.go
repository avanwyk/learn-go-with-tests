package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertAddEquals := func(t *testing.T, expected int, actual int) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected %d, but got actual %d", expected, actual)
		}
	}

	t.Run("should add", func(t *testing.T) {
		actual := Add(2, 3)
		expected := 5
		assertAddEquals(t, expected, actual)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
