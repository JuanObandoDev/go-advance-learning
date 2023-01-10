package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(n int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	balance += n
	lock.Unlock()
}

func WithDraw(n int, wg *sync.WaitGroup, lock *sync.RWMutex) bool {
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

func Balance(lock *sync.RWMutex) (b int) {
	lock.RLock()
	b = balance
	lock.RUnlock()
	return
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go WithDraw(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}
