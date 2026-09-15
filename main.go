package main

import "fmt"

func main() {
	greeting := greet("el")
	fmt.Println(greeting)
}

// language represents language's code
type language string

// phrasebook holds each greeting
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	// Greek
	"en": "Hello world",
	// English
	"fr": "Bonjour le monde",  // French
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
