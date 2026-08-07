package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguage() (string, string, bool) {
	return "Go", "Python", true
}

func functions() {

	// result := add(2, 3)
	// fmt.Println("The result is:", result)

	lang1, lang2, _ := getLanguage()
	fmt.Println("Languages:", lang1, lang2)
}
