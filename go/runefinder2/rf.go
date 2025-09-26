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

// hasSpecialChar reports whether the text contains any character that is not
// A-Z, a-z, 0-9, space, or hyphen.
func hasSpecialChar(text string) bool {
	// Regex that matches only allowed characters: A-Z, a-z, 0-9, space, hyphen
	allowedPattern := regexp.MustCompile(`^[A-Za-z0-9 -]*$`)
	return !allowedPattern.MatchString(text)
}

// listChars lists each character in the text using the same format as search output.
func listChars(text string) {
	count := 0
	for _, char := range text {
		name := runenames.Name(char)
		if len(name) == 0 {
			// For characters without official Unicode names, show a placeholder
			fmt.Printf("%U\t%c\t<no name>\n", char, char)
		} else {
			fmt.Printf("%U\t%c\t%v\n", char, char, name)
		}
		count++
	}
	fmt.Printf("(%d found)\n", count)
}

func find(text string, firstLast ...rune) {
	first, last := default_first, default_last
	switch len(firstLast) {
	case 0:
		// done
	case 1:
		first = firstLast[0]
	case 2:
		first, last = firstLast[0], firstLast[1]
	default:
		panic("find: firstLast must be 0, 1 or 2 runes")
	}
	words := tokenize(text)
	query := []string{}
	for _, word := range words {
		query = append(query, strings.ToUpper(word))
	}
	count := 0
	for char := first; char <= last; char++ {
		name := runenames.Name(char)
		if len(name) == 0 {
			continue
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
		if hasSpecialChar(text) {
			listChars(text)
		} else {
			find(text)
		}
	} else {
		fmt.Println("Please provide words to find.\nExample:")
		fmt.Println(os.Args[0], "cat eyes")
	}
}
