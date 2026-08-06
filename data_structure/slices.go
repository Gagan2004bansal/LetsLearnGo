package main

import (
	"fmt"
	"slices"
)

// slice : means dynamic array
// most used construct in Go
func slice() {

	// unintialized slice is nil
	var nums []int
	fmt.Println(nums)
	fmt.Println("Length of slice :", len(nums)) // Print the length/size of slice
	fmt.Println(nums == nil)

	var arr = make([]int, 2, 5) // create a slice with length 2 and capacity 5
	// capacity -> the number of elements the slice can hold before it needs to be resized
	fmt.Println(arr)
	fmt.Println(cap(arr)) // Print the capacity of slice before append anything
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3, 4, 5) // append multiple values to the slice
	arr[0] = 10
	fmt.Println(arr)
	fmt.Println(cap(arr)) // Print the capacity of slice after append

	nums2 := []int{}
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	fmt.Println(nums2)

	// copy slice
	var arr2 = make([]int, 2, 5)
	arr2[0] = 1
	var arr3 = make([]int, len(arr2))

	copy(arr3, arr2) // copy the contents of arr2 to arr3
	fmt.Println(arr3)

	// slice operator
	var arr4 = []int{1, 2, 3, 4, 5}
	fmt.Println(arr4[1:3]) // slice from index 1 to 3 (not including 3)

	// slice
	var arr5 = []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Equal(arr5, arr4)) // true
}
