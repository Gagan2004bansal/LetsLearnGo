package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Razorpay.")
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Stripe.")
}

func main() {

	payment := payment{gateway: stripe{}}
	payment.makePayment(1000)

}
