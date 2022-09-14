package main

import "fmt"

type Command interface {
	Execute()
}

type Pizza struct{}

func NewPizza() *Pizza {
	return &Pizza{}
}

func (p *Pizza) cook() {
	fmt.Println("Pizza is being cooked!")
}

type DeliveryMan struct{}

func NewDeliveryMan() *DeliveryMan {
	return &DeliveryMan{}
}

func (dm *DeliveryMan) deliver() {
	fmt.Println("Delivery man is performing the delivery!")
}

type DeliveryManDeliverCommand struct {
	deliveryMan *DeliveryMan
}

func NewDeliveryManDeliverCommand(deliveryMan *DeliveryMan) *DeliveryManDeliverCommand {
	return &DeliveryManDeliverCommand{deliveryMan: deliveryMan}
}

func (gdo *DeliveryManDeliverCommand) Execute() {
	gdo.deliveryMan.deliver()
}

type PizzaCookCommand struct {
	pizza *Pizza
}

func (loc *PizzaCookCommand) Execute() {
	loc.pizza.cook()
}

func NewPizzaCookCommand(pizza *Pizza) *PizzaCookCommand {
	return &PizzaCookCommand{pizza: pizza}
}

type SimpleRemoteControl struct {
	slot Command
}

func NewSimpleRemoteControl(slot Command) *SimpleRemoteControl {
	return &SimpleRemoteControl{slot: slot}
}

func (src *SimpleRemoteControl) SetCommand(command Command) {
	src.slot = command
}

func (src *SimpleRemoteControl) ButtonPressed() {
	src.slot.Execute()
}

//func main() {
//	remote := NewSimpleRemoteControl(nil)
//
//	pizza := NewPizza()
//	deliveryMan := NewDeliveryMan()
//
//	pizzaCookCommand := NewPizzaCookCommand(pizza)
//	deliveryManDeliverCommand := NewDeliveryManDeliverCommand(deliveryMan)
//
//	remote.SetCommand(pizzaCookCommand)
//	remote.ButtonPressed()
//
//	remote.SetCommand(deliveryManDeliverCommand)
//	remote.ButtonPressed()
//}
