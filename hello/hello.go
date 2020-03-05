package main

import "fmt"

const es = "Spanish"
const fr = "French"
const af = "Afrikaans"
const enGreetingPrefix = "Hello, "
const esGreetingPrefix = "Hola, "
const frGreetingPrefix = "Bonjour, "
const afGreetingPrefix = "Goeiedag, "

func Hello(entity string, language string) string {
	if entity == "" {
		entity = "World"
	}
	return greetingPrefix(language) + entity
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case es:
		prefix = esGreetingPrefix
	case fr:
		prefix = frGreetingPrefix
	case af:
		prefix = afGreetingPrefix
	default:
		prefix = enGreetingPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
