// TODO: Take user input and count the number of words and find the longest word
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a word of your Choice: ")
	user := bufio.NewScanner(os.Stdin)
	user.Scan()

	userInput := strings.TrimSpace(user.Text())

	// splitting the sentence to words
	word := strings.Fields(userInput)

	// Getting the count
	wordCount := len(word)

	// Getting longest word

	longestWord := ""
	for _, wod := range word {
		if len(wod) > len(longestWord) {
			longestWord = wod
		}
	}
	fmt.Printf("Word count   : %d\n", wordCount)
	fmt.Printf("Longest word : %s\n", longestWord)
}
