package main

import (
	"fmt"
	"sync"
)

// globar shared variable
var balance int

// func deposit using waitgroup
func Deposit(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	balance += n
	fmt.Printf("Transaction Done, balance: %d\n", balance)
}

// func withdraw using waitgroup
func WithDraw(n int, wg *sync.WaitGroup) bool {
	defer wg.Done()
	// validating n is less than balance
	if n > balance {
		fmt.Println("Insufficient funds")
		return false
	}
	balance -= n
	return true
}

func main() {
	var wg sync.WaitGroup // create waitgroup
	balance = 500         // initial balance
	fmt.Println("Balance before deposit: ", balance)
	wg.Add(2)             // adding 2 waitgroups
	go Deposit(200, &wg)  // deposit 200
	go WithDraw(700, &wg) // withdraw 700 (it doesn't have enough balance at this point)
	wg.Wait()
	fmt.Println("Balance after deposit: ", balance)
}
