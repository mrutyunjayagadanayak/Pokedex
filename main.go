package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var result []string
	if text == "" {
		return result
	}
	text = strings.ToLower(text)
	result = strings.Fields(text)
	return result
}
