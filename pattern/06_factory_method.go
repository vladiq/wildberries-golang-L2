package main

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(float32) string
}

func GetPaymentMethod(method string) (PaymentMethod, error) {
	switch method {
	case "cash":
		return new(PaymentInCash), nil
	case "debit card":
		return new(PaymentViaDebitCard), nil
	default:
		return nil, errors.New(fmt.Sprintf("payment method %s not recognized", method))
	}
}

type PaymentInCash struct{}

func (pic *PaymentInCash) Pay(amount float32) string {
	return fmt.Sprintf("$%.2f paid in cash", amount)
}

type PaymentViaDebitCard struct{}

func (pvdc *PaymentViaDebitCard) Pay(amount float32) string {
	return fmt.Sprintf("$%.2f paid via debit card", amount)
}

//func main() {
//	if payment, err := GetPaymentMethod("cash"); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(payment.Pay(150))
//	}
//
//	if payment, err := GetPaymentMethod("debit card"); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(payment.Pay(150.00))
//	}
//
//	if payment, err := GetPaymentMethod("credit card"); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(payment.Pay(150.00))
//	}
//}
