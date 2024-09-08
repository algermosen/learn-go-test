package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	spanishLang = "Spanish"
	frenchLang  = "French"
	italianLang = "Italian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	italianHelloPrefix = "Ciao, "
)

func Hello(s string, language string) string {
	if s == "" {
		s = "world"
	}

	return greetingPrefix(language) + s
}

func greetingPrefix(language string) string {
	switch language {
	case spanishLang:
		return spanishHelloPrefix
	case frenchLang:
		return frenchHelloPrefix
	case italianLang:
		return italianHelloPrefix
	default:
		return englishHelloPrefix
	}
}
