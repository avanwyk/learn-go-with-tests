package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("should repeat character n times", func(t *testing.T) {
		actual := Repeat("a", 6)
		expected := "aaaaaa"

		if actual != expected {
			t.Errorf("Expected %q but got: %q", expected, actual)
		}
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("b", 3))
	// Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
