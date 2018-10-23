package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

const first, last = ' ', unicode.MaxRune

func find(words ...string) {
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
		if containsAll(strings.Fields(name), query) {
			fmt.Printf("%U\t%c\t%v\n", char, char, name)
			count++
		}
	}
	fmt.Printf("(%d found)\n", count)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide words to find.")
		return
	}
	find(os.Args[1:]...)
}
