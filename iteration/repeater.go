package iteration

import "strings"

func Repeat(input string, n int) string {
	// strings.Repeat(input, n)
	var buffer strings.Builder
	for i := 0; i < n; i++ {
		buffer.WriteString(input)
	}
	return buffer.String()
}
