package iteration

import (
	"fmt"
	"strings"
	"unicode"
)

func StringIterations() {
	text := "Hello World"

	// Rune Based Iteration
	// Can handle Unicode Characters Correctly
	for i, v := range text {
		fmt.Printf("Index: %d Value: %c \n", i, v)

	}

	// Byte Based Iteration
	// Only Works with ASCII Characters
	for i := 0; i < len(text); i++ {
		fmt.Println(text[i])

	}

	// Rune Iteration
	runes := []rune(text)
	for _, r := range runes {

		fmt.Printf("Value is: %c \n", r)
	}

	processString(text)

	// Conditonal String Processing
	for _, char := range text {

		switch {
		case unicode.IsLetter(char):
			fmt.Printf("Letter: %c", char)
		case unicode.IsDigit(char):
			fmt.Printf("Letter: %d", char)

		}

	}

	fmt.Println()

	fmt.Printf("Word Count: %d", countWords(text))

}

func processString(text string) {
	for _, r := range text {

		// Safe Unicode Processing
		if r > 127 {
			fmt.Println("Non-ASCII character detected")

		}

	}
}

func countWords(text string) int {

	wordCount := 0
	inWord := false

	for _, v := range text {

		if unicode.IsLetter(v) && !inWord {

			wordCount++
			inWord = true
		}

		if !unicode.IsLetter(v) {
			inWord = false

		}
	}

	return wordCount

}
func transformString(input string) string {
	var result strings.Builder

	for _, char := range input {
		switch {
		case unicode.IsLower(char):
			result.WriteRune(unicode.ToUpper(char))
		case unicode.IsUpper(char):
			result.WriteRune(unicode.ToLower(char))
		default:
			result.WriteRune(char)
		}
	}

	return result.String()
}
