package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Some 
	test to do something here and there
	 What about you
	 How are you doing here
	`
	words := strings.Fields(text)
	fmt.Println(words)

	println("-------")
	counts := map[string]int{}

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
	
}
