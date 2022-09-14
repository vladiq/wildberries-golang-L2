package main

import "fmt"

///*
//	Реализовать паттерн «фасад».
//Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
//	https://en.wikipedia.org/wiki/Facade_pattern
//*/

type account struct {
	phoneNumber string
	address     string
	cardNumber  string
}

func (a *account) setPhoneNumber(phoneNumber string) {
	a.phoneNumber = phoneNumber
}

type paymentSystem struct{}

func (ps *paymentSystem) processPayment(cardNumber string) {
	fmt.Println("Payment successful!")
}

type cooker struct{}

func (c *cooker) cook(orderItemList []string) {
	fmt.Println("Started cooking: ", orderItemList)
}

type deliveryService struct {
	phoneForNotifications string
	deliveryAddress       string
}

func (ds *deliveryService) setPhoneForNotification(phone string) {
	ds.phoneForNotifications = phone
}

func (ds *deliveryService) setDeliveryAddress(address string) {
	ds.deliveryAddress = address
}

func (ds *deliveryService) notify() {
	fmt.Printf(
		"Dear customer! Your order will be delivered to %s. Messages will be sent to %s.\n",
		ds.deliveryAddress,
		ds.phoneForNotifications,
	)
}

type pizzaDeliveryFacade struct {
	account         *account
	paymentSystem   *paymentSystem
	cooker          *cooker
	deliveryService *deliveryService
}

func NewPizzaDeliveryFacade() *pizzaDeliveryFacade {
	return &pizzaDeliveryFacade{
		account:         &account{},
		paymentSystem:   &paymentSystem{},
		cooker:          &cooker{},
		deliveryService: &deliveryService{},
	}
}

func (pdf *pizzaDeliveryFacade) logIn(phoneNumber string) {
	pdf.account.setPhoneNumber(phoneNumber)
	pdf.deliveryService.setPhoneForNotification(phoneNumber)

	fmt.Println("Login successful!")
}

func (pdf *pizzaDeliveryFacade) order(orderItemList []string, cardNumber, address string) {
	pdf.paymentSystem.processPayment(cardNumber)

	pdf.cooker.cook(orderItemList)

	pdf.deliveryService.setDeliveryAddress(address)
	pdf.deliveryService.notify()
}

//func main() {
//	pizzaDeliveryService := NewPizzaDeliveryFacade()
//
//	phoneNumber := "+79998887766"
//	pizzaDeliveryService.logIn(phoneNumber)
//
//	orderItemList := []string{"Pepperoni", "Spicy", "Chicken and Pork", "Chernogolovka Fizzy Drink"}
//	cardNumber := "0000 0000 0000 0000"
//	address := "Pushkin st. 15, 34"
//	pizzaDeliveryService.order(orderItemList, cardNumber, address)
//}
