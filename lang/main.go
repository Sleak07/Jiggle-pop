package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language,e.g. en,ur")
	flag.Parse()
	greeting := greet(language(lang))
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
