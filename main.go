package main

import "fmt"

func main() {
	greeting := greet("fr")
	fmt.Println(greeting)
}

// language represents language's code
type language string

// greet returns greeting to world
func greet(l language) string {
	switch l {
	case "en":
		return "Hello World"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
