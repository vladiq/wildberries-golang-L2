package main

import (
	"errors"
	"fmt"
)

type State interface {
	requestPizza() error
	cookPizza(int) error
	payMoney(int) error
	giveOutPizza() error
}

type PizzaShop struct {
	hasPizza       State
	pizzaRequested State
	hasMoney       State
	noPizza        State

	currentState State

	pizzaCount int
	pizzaPrice int
}

func NewPizzaShop(pizzaCount int, pizzaPrice int) *PizzaShop {
	shop := &PizzaShop{
		pizzaCount: pizzaCount,
		pizzaPrice: pizzaPrice,
	}
	hasPizza := &hasPizzaState{
		pizzaShop: shop,
	}
	pizzaRequested := &pizzaRequestedState{
		pizzaShop: shop,
	}
	hasMoney := &hasMoneyState{
		pizzaShop: shop,
	}
	noPizza := &noPizzaState{
		pizzaShop: shop,
	}

	shop.setState(hasPizza)
	shop.hasPizza = hasPizza
	shop.pizzaRequested = pizzaRequested
	shop.hasMoney = hasMoney
	shop.noPizza = noPizza

	return shop
}

func (ps *PizzaShop) requestPizza() error {
	return ps.currentState.requestPizza()
}

func (ps *PizzaShop) cookPizza(count int) error {
	return ps.currentState.cookPizza(count)
}

func (ps *PizzaShop) payMoney(money int) error {
	return ps.currentState.payMoney(money)
}

func (ps *PizzaShop) giveOutPizza() error {
	return ps.currentState.giveOutPizza()
}

func (ps *PizzaShop) setState(s State) {
	ps.currentState = s
}

func (ps *PizzaShop) incrementPizzaCount(count int) {
	fmt.Printf("Cooking %d pizzas\n", count)
	ps.pizzaCount += count
}

type noPizzaState struct {
	pizzaShop *PizzaShop
}

func (hps *noPizzaState) requestPizza() error {
	return errors.New("no pizza available")
}

func (hps *noPizzaState) cookPizza(count int) error {
	hps.pizzaShop.incrementPizzaCount(count)
	hps.pizzaShop.setState(hps.pizzaShop.hasPizza)
	return nil
}

func (hps *noPizzaState) payMoney(money int) error {
	return errors.New("no pizza available")
}

func (hps *noPizzaState) giveOutPizza() error {
	return errors.New("no pizza available")
}

type hasPizzaState struct {
	pizzaShop *PizzaShop
}

func (hps *hasPizzaState) requestPizza() error {
	if hps.pizzaShop.pizzaCount == 0 {
		hps.pizzaShop.setState(hps.pizzaShop.noPizza)
		return errors.New("no pizzas cooked yet")
	}
	fmt.Println("Pizza requested")
	hps.pizzaShop.setState(hps.pizzaShop.pizzaRequested)
	return nil
}

func (hps *hasPizzaState) cookPizza(count int) error {
	fmt.Println(count, "pizzas cooked")
	hps.pizzaShop.incrementPizzaCount(count)
	return nil
}

func (hps *hasPizzaState) payMoney(money int) error {
	return errors.New("order your pizza first")
}

func (hps *hasPizzaState) giveOutPizza() error {
	return errors.New("order your pizza first")
}

type pizzaRequestedState struct {
	pizzaShop *PizzaShop
}

func (prs *pizzaRequestedState) requestPizza() error {
	return errors.New("pizza already requested")
}

func (prs *pizzaRequestedState) cookPizza(count int) error {
	return errors.New("order preparation in progress")
}

func (prs *pizzaRequestedState) payMoney(money int) error {
	if money < prs.pizzaShop.pizzaPrice {
		return fmt.Errorf("not enough money, you must pay $%d", prs.pizzaShop.pizzaPrice)
	}
	fmt.Println("Received money")
	prs.pizzaShop.setState(prs.pizzaShop.hasMoney)
	return nil
}

func (prs *pizzaRequestedState) giveOutPizza() error {
	return errors.New("you have to pay first")
}

type hasMoneyState struct {
	pizzaShop *PizzaShop
}

func (hms *hasMoneyState) requestPizza() error {
	return fmt.Errorf("order preparation in progress")
}

func (hms *hasMoneyState) cookPizza(count int) error {
	return fmt.Errorf("order preparation in progress")
}

func (hms *hasMoneyState) payMoney(money int) error {
	return fmt.Errorf("order preparation in progress")
}

func (hms *hasMoneyState) giveOutPizza() error {
	fmt.Println("Giving pizza out...")
	hms.pizzaShop.pizzaCount--
	if hms.pizzaShop.pizzaCount == 0 {
		hms.pizzaShop.setState(hms.pizzaShop.noPizza)
	} else {
		hms.pizzaShop.setState(hms.pizzaShop.hasPizza)
	}
	return nil
}

//func main() {
//	pizzaShop := NewPizzaShop(1, 10)
//
//	err := pizzaShop.requestPizza()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = pizzaShop.payMoney(10)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = pizzaShop.giveOutPizza()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	fmt.Println()
//
//	err = pizzaShop.cookPizza(2)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//	fmt.Println()
//
//	err = pizzaShop.requestPizza()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = pizzaShop.payMoney(10)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = pizzaShop.giveOutPizza()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//}
