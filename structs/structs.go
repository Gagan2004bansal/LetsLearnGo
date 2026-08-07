package main

import (
	"fmt"
	"time"
)

// customer struct
type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Nanosecond precision
	customer
}

// constructor function to create a new order
func newOrder(id string, amount float32, status string, customer customer) *order {
	order := order{
		id:       id,
		amount:   amount,
		status:   status,
		customer: customer,
	}

	return &order
}

// changeStatus method to change the status of the order
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {

	o1 := order{
		id:     "1",
		amount: 50.00,
		status: "pending",
	}

	o1.createdAt = time.Now()

	fmt.Println("Order :", o1)
	o1.changeStatus("delivered")
	fmt.Println("Order status :", o1.status)

	// o2 := newOrder("2", 100.00, "pending")
	// fmt.Println("Order 2 :", o2)

	newCustomer := customer{
		name:  "Gagan",
		phone: "1234567890",
	}

	o3 := newOrder("3", 150.00, "pending", newCustomer)
	fmt.Println(o3)

}
