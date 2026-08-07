package main

import "fmt"

func printSlice[T any, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"Hello", "World", "Generics"}

	printSlice(slice, "John")
	printSlice(stringSlice, "John")

	myStack := stack[string]{elements: stringSlice}
	fmt.Println(myStack)
}
