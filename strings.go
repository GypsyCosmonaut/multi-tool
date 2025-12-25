package main

import (
	"fmt"
	"regexp"
)

func Reverse(text string) string {
    n := len(text)
    runes := make([]rune, n)
    for _, rune := range text {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}

func main() {
	text := "Big fan of Elon"

        // string indexing -> bytes
	// rune indexing -> characters

	text_rune := []rune(text)

	if len(text_rune) < 3 {
		fmt.Println("String must be at least 6 characters long")
		return
	}

	inverted_suffix := "ksuM "

	suffix := Reverse(inverted_suffix)

	suffix_rune := []rune(suffix)

	statement := append(text_rune, suffix_rune...)

	//Concatenate 2 strings
	text = text + " Musk"

	fmt.Println("Concatenated String:", text)

	//Extract first n characters
	first := statement[:3]

	//Extract last n characters
	last := statement[len(statement)-3:]

	fmt.Println("First 3:", string(first))
	fmt.Println("Last 3:", string(last))

	//Extract regex match from string
        regex := regexp.MustCompile(`\p{Lu}`)
        matches := regex.FindAllString(string(statement), -1)

	fmt.Println("Regex Match:", matches)

	for _, match := range matches {
		fmt.Println(match + "!")
	}
}

