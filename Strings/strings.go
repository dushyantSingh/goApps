package main

import (
	"fmt"
	"strconv"
)

func main() {
	book := "The color of magic"
	println(book)
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	println("Slice", book[4:11])
	println("Slices", book[4:])
	println("Slice33", book[:11])

	some := book + "Library"
	println(some)

	u, err := strconv.ParseInt("-42", 10, 64)
	println("String and err", u, err)

	a, erra := strconv.Atoi("44444c")
	fmt.Printf("String = %v (type %T) error %v\n", a, a, erra)
	s := strconv.Itoa(44444)
	fmt.Printf("String = %v (type %T)\n", s, s)
}
