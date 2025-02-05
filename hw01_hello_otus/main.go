package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func GetReversedString(s string) string {
	return reverse.String(s)
}

func main() {
	const phrase = "Hello, OTUS!"
	fmt.Print(GetReversedString(phrase))
}
