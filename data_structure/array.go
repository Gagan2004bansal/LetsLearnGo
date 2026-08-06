package main

import "fmt"

func array() {
	var nums [4]int

	nums[0] = 1

	fmt.Println("Length of array :", len(nums))         // Print the length/size of array
	fmt.Println("Value of array :", nums)               // Print the value of array
	fmt.Println("Value of array at index 0 :", nums[0]) // Print the value of array at index 0

	arr := [3]int{1, 2, 3}               // Declare and initialize an array
	fmt.Println("Value of array :", arr) // Print the value of array

	arr2d := [2][2]int{{1, 2}, {3, 4}}        // Declare and initialize a 2D array
	fmt.Println("Value of 2D array :", arr2d) // Print the value of 2D array
}
