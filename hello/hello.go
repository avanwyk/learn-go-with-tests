package main

import "fmt"

func Hello(entity string) string {
	return fmt.Sprintf("Hello, %s", entity)
}

func main() {
	fmt.Println(Hello("world"))
}
