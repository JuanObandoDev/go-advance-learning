package main

import "fmt"

type IPayment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Paying with cash")
}

func ProcessPayment(p IPayment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying with bank account %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	*BankPayment
	bankAccount int
}

func (bpa BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	// bankAccount := &BankPayment{}
	// ProcessPayment(bankAccount)

	bankAccount := &BankPaymentAdapter{
		&BankPayment{},
		123456,
	}
	ProcessPayment(bankAccount)
}
