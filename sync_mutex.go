package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(n int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	balance += n
	lock.Unlock()
}

func WithDraw(n int, wg *sync.WaitGroup, lock *sync.Mutex) bool {
	defer wg.Done()
	lock.Lock()
	if n > balance {
		fmt.Println("Insufficient funds")
		lock.Unlock()
		return false
	}
	balance -= n
	lock.Unlock()
	return true
}

func Balance() (b int) {
	b = balance
	return
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go WithDraw(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())
}
