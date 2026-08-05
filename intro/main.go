package main

// package import from Go standard library
import "fmt"

// Entry Point for the Go Program for execution
func main() {
	fmt.Println("Let's Learn GoLang!")

	// simple values
	// int
	fmt.Println(1 + 1)
	// string
	fmt.Println("Hello" + " " + "GoLang")
	// bool
	fmt.Println(true)
	// float
	fmt.Println(3.14)

	// Variable declaration and initialization
	var name string
	name = "Gagan Bansal" // is same as var name = "Gagan Bansal" (automatically infers the type)
	fmt.Println(name)

	var age int = 22
	fmt.Println(age)

	// Shorthand syntax
	nickname := "Gagan"
	fmt.Println(nickname)

	// Constants
	const pi = 3.14

	// multiple variable declaration
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

	// for -> only construct in Go for looping
	for i := 0; i < 3; i++ {
		fmt.Printf("Num : %d\n", i)
	}

	// while loop using for
	j := 0
	for j < 3 {
		fmt.Printf("Num : %d\n", j)
		j++
	}

	// infinite loop using for
	for {
		fmt.Println("Infinite Loop")
		break // break the infinite loop
	}

	for i := range 3 {
		fmt.Printf("Num : %d\n", i)
	}

	// if else
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 12 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a minor")
	}
	// Go Does not have ternary, you will have to use if else statements instead

	// Switch
	switch age {
	case 18:
		fmt.Println("You are an adult")
	case 12:
		fmt.Println("You are a teenager")
	default: // It is optional
		fmt.Println("You are a minor")
	}
}
