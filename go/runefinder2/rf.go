package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

const default_first, default_last = ' ', unicode.MaxRune

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

// Regex to match characters used in Unicode names:
// case-insensitive A-Z, 0-9, space, hyphen
var unicodeNameRe = regexp.MustCompile(`(?i)^[A-Z0-9 -]+$`)

// likeUnicodeName reports whether the text contains
// only characters used in Unicode character names.
func likeUnicodeName(text string) bool {
	return unicodeNameRe.MatchString(text)
}

// listChars lists each rune in the text.
func listChars(text string) {
	count := 0
	for _, char := range text {
		name := runenames.Name(char)
		fmt.Printf("%U\t%c\t%v\n", char, char, name)
		count++
	}
	fmt.Printf("(%d found)\n", count)
}

// find lists runes whose Unicode names have all words from
// given text, searching within optional firstLast rune range.
func find(text string, firstLast ...rune) {
	first, last := default_first, default_last
	switch len(firstLast) {
	case 0:
		// use default values
	case 1:
		first = firstLast[0]
	case 2:
		first, last = firstLast[0], firstLast[1]
	default:
		panic("find: firstLast must be 0, 1 or 2 runes")
	}
	query := tokenize(strings.ToUpper(text))
	count := 0
	for char := first; char <= last; char++ {
		name := runenames.Name(char)
		if len(name) == 0 {
			name = "NAMELESS"
		}
		if containsAll(tokenize(name), query) {
			fmt.Printf("%U\t%c\t%v\n", char, char, name)
			count++
		}
	}
	fmt.Printf("(%d found)\n", count)
}

func main() {
	if len(os.Args) > 1 {
		text := strings.Join(os.Args[1:], " ")
		if likeUnicodeName(text) {
			find(text)
		} else {
			listChars(text)
		}
	} else {
		fmt.Println("Please give words to find or characters to get their names.")
		fmt.Print("Example:\n\t")
		fmt.Println(os.Args[0], "cat eyes")
		fmt.Println("Use . or special characters in arguments to force character mode.")
		fmt.Print("Example:\n\t")
		fmt.Println(os.Args[0], ". 👋")
	}
}
