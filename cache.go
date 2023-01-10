package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f,
		make(map[int]FunctionResult),
		sync.Mutex{},
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, ok := m.cache[key]
	m.lock.Unlock()

	if !ok {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}

	return result.value, result.err
}

func GetFib(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	var wg sync.WaitGroup
	cache := NewCache(GetFib)
	fibo := []int{10, 20, 30, 40, 40, 30, 20, 10}
	for _, n := range fibo {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(i)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d, %s, %d\n", i, time.Since(start), value)
		}(n)
		wg.Wait()
	}
}
