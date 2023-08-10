package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	english_greet = "Hello, "
	french_greet  = "Bonjour, "
	spanish_greet = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = french_greet
	case spanish:
		prefix = spanish_greet
	default:
		prefix = english_greet
	}
	return
}
func main() {
	fmt.Print(Hello("aman", "English"))
}
