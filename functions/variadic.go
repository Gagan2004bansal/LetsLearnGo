package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func variadic() {

	fmt.Println("Variadic Functions")

	arr := []int{2, 3, 5, 5}
	fmt.Println(sum(arr...))

}
