package main

import "fmt"

// by value <- copy of the value is passed to the function
func changeNum(num int) {
	num = 5
	fmt.Println("Inside changeNum function, num:", num)
}

// By reference <- address of the value is passed to the function
func changeNumRef(num *int) {
	*num = 5
	fmt.Println("Inside changeNumRef function, num:", *num)
}

func main() {
	num := 10
	changeNum(num)
	fmt.Println("Outside changeNum function, num:", num)

	changeNumRef(&num)
	fmt.Println("Outside changeNumRef function, num:", num)
}
