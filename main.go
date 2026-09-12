package main

import "fmt"

func main() {
	greeting := greet("fr")
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
	"fr": "Bonjour le monde", // French
	"ur": "‫دﻧﯿﺎ‬ ‫ﯿﻠﻮ‬ ",
	// Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
