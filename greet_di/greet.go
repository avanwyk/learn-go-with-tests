package greet_di

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	return err
}

func main() {
	Greet(os.Stdout, "User")
}
