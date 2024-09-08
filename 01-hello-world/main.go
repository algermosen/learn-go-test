package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(s string) string {
	return fmt.Sprintf("Hello, %s", s)
}
