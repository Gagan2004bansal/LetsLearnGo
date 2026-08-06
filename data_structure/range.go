package main

import "fmt"

// iterating over data structures using range
func ranges() {

	nums := []int{1, 2, 3, 4, 5}

	for index, num := range nums {
		fmt.Println(index, "index ->", num)
	}

	m := map[string]string{"name": "gagan", "age": "22"}
	for key, value := range m {
		fmt.Println(key, "key ->", value)
	}

	// c -> unicode code point rune
	// i -> starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, "index ->", string(c))
	}

}
