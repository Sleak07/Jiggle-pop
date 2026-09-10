package main

import "fmt"

func main() {
	greeting := greet()
	fmt.Println(greeting)
}

// greet returns greeting to world
func greet() string {
	return "Hello World"
}
