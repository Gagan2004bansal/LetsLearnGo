package main

import "fmt"

// enumerated type
// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeStatus(status OrderStatus) {
	fmt.Println("Status changed to:", status)
}

func main() {
	changeStatus(Prepared)
}
