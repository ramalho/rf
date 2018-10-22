package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

// Find prints characters with names matching all of the given words.
func find(words ...string) {
	query := []string{}
	for _, word := range words {
		query = append(query, strings.ToUpper(word))
	}
	for char := ' '; char <= unicode.MaxRune; char++ {
		name := runenames.Name(char)
		if ContainsAll(strings.Fields(name), query...) {
			fmt.Printf("%U\t%c\t%v\n", char, char, name)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide one or more words to find.")
		return
	}
	find(os.Args[1:]...)
}
