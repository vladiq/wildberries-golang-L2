package main

import "fmt"

type Stage interface {
	Execute(*Customer)
	SetNext(Stage)
}

type Login struct {
	next Stage
}

func (l *Login) Execute(c *Customer) {
	if c.loggedIn {
		fmt.Println("Customer is already logged in")
	} else {
		fmt.Println("Logging in")
		c.loggedIn = true
	}
	l.next.Execute(c)
	return
}

func (l *Login) SetNext(next Stage) {
	l.next = next
}

type Payment struct {
	next Stage
}

func (p *Payment) Execute(c *Customer) {
	if c.paymentProcessed {
		fmt.Println("The payment is already processed")
	} else {
		fmt.Println("Processing the payment")
		c.paymentProcessed = true
	}
	p.next.Execute(c)
	return
}

func (p *Payment) SetNext(next Stage) {
	p.next = next
}

type Kitchen struct {
	next Stage
}

func (k *Kitchen) Execute(c *Customer) {
	if c.pizzaCooked {
		fmt.Println("The pizza is already cooked")
	} else {
		fmt.Println("Cooking the pizza")
		c.pizzaCooked = true
	}
	k.next.Execute(c)
	return
}

func (k *Kitchen) SetNext(next Stage) {
	k.next = next
}

type Delivery struct {
	next Stage
}

func (d *Delivery) Execute(c *Customer) {
	if c.deliveryDone {
		fmt.Println("The order is already delivered")
	} else {
		fmt.Println("Delivering the order")
		c.deliveryDone = true
	}
}

func (d *Delivery) SetNext(next Stage) {
	d.next = next
}

type Customer struct {
	loggedIn         bool
	paymentProcessed bool
	pizzaCooked      bool
	deliveryDone     bool
}

//func main() {
//	login := new(Login)
//	payment := new(Payment)
//	kitchen := new(Kitchen)
//	delivery := new(Delivery)
//
//	login.SetNext(payment)
//	payment.SetNext(kitchen)
//	kitchen.SetNext(delivery)
//
//	customer := new(Customer)
//	login.Execute(customer)
//	fmt.Println()
//
//	customer = &Customer{loggedIn: true, paymentProcessed: true}
//	login.Execute(customer)
//}
