package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

const default_first, default_last = ' ', unicode.MaxRune

func tokenize(text string) []string {
	return strings.Fields(strings.ReplaceAll(text, "-", " "))
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
		find(strings.Join(os.Args[1:], " "))
	} else {
		fmt.Println("Please provide words to find.")
	}
}
