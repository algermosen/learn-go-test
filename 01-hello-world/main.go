package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

const englishHelloPrefix = "Hello, "

func Hello(s string) string {
	if s == "" {
		s = "world"
	}
	return englishHelloPrefix + s
}
