package main

import (
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

type RuneName struct {
	Rune rune
	Name string
}

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

func find(text string, first rune, last rune) iter.Seq[RuneName] {
	return func(yield func(RuneName) bool) {
		query := tokenize(strings.ToUpper(text))
		for char := first; char <= last; char++ {
			name := runenames.Name(char)
			if len(name) > 0 && containsAll(tokenize(name), query) {
				if !yield(RuneName{Rune: char, Name: name}) {
					return
				}
			}
		}
	}
}

func report(text string, first rune, last rune) {
	count := 0
	for r := range find(text, first, last) {
		fmt.Printf("%U\t%c\t%v\n", r.Rune, r.Rune, r.Name)
		count++
	}
	fmt.Printf("(%d found)\n", count)
}

func main() {
	if len(os.Args) > 1 {
		report(strings.Join(os.Args[1:], " "), ' ', unicode.MaxRune)
	} else {
		fmt.Println("Please provide words to find.\nExample:")
		fmt.Println(os.Args[0], "cat eyes")
	}
}
