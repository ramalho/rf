package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

func tokenize(text string) []string {
	return strings.Fields(strings.ReplaceAll(text, "-", " "))
}

// containsAll reports whether list contains all items.
func containsAll(list []string, items []string) bool {
	for _, s := range items {
		if !slices.Contains(list, s) {
			return false
		}
	}
	return true
}

func find(text string, first rune, last rune) {
	query := tokenize(strings.ToUpper(text))
	count := 0
	for char := first; char <= last; char++ {
		name := runenames.Name(char)
		if len(name) > 0 && containsAll(tokenize(name), query) {
			fmt.Printf("%U\t%c\t%v\n", char, char, name)
			count++
		}
	}
	fmt.Printf("(%d found)\n", count)
}

func main() {
	if len(os.Args) > 1 {
		find(strings.Join(os.Args[1:], " "), ' ', unicode.MaxRune)
	} else {
		fmt.Println("Please provide words to find.\nExample:")
		fmt.Println(os.Args[0], "cat eyes")
	}
}
