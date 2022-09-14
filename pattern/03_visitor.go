package main

import (
	"fmt"
)

type pizzaOrder struct {
	orderId      uint32
	orderedItems map[string]float32
	address      string
	customerName string
}

func (po *pizzaOrder) Accept(v Visitor) {
	v.visit(po)
}

func NewPizzaOrder(orderedItems map[string]float32, address, customerName string) *pizzaOrder {
	return &pizzaOrder{orderedItems: orderedItems, address: address, customerName: customerName}
}

type Visitor interface {
	visit(*pizzaOrder)
}

type pizzaOrderVisitor struct{}

func (pov *pizzaOrderVisitor) visit(po *pizzaOrder) {
	var totalOrderCost float32
	for _, v := range po.orderedItems {
		totalOrderCost += v
	}
	fmt.Printf(
		"Dear %s, the order %d at the cost of %.2f will be delivered to %s!\n",
		po.customerName, po.orderId, totalOrderCost, po.address,
	)
}

func NewPizzaOrderVisitor() *pizzaOrderVisitor {
	return &pizzaOrderVisitor{}
}

//func main() {
//	orderedItems := map[string]float32{
//		"Pizza pepperoni": 1499.99,
//		"Coca Cola":       64.99,
//	}
//	address := "Baker st. 221b"
//	name := "Sherlock Holmes"
//
//	order := NewPizzaOrder(orderedItems, address, name)
//	visitor := NewPizzaOrderVisitor()
//	order.Accept(visitor)
//}
