package main

import "fmt"

type PaymentMethods interface {
	payAmount()
}

func payment(p PaymentMethods) {
	p.payAmount()
}

type creditCard struct {
}

func (c creditCard) payAmount() {
	fmt.Println("xxxxx.xx amount got paid through credit card")
}

type bharatPay struct {
}

func (c bharatPay) payAmount() {
	fmt.Println("xxxxx.xx amount got paid through bharatPay")
}

type upi struct {
}

func (c upi) payAmount() {
	fmt.Println("xxxxx.xx amount got paid through upi")
}
