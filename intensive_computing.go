package main

import (
	"fmt"
	"sync"
)

func ExpensiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return ExpensiveFibonacci(n-1) + ExpensiveFibonacci(n-2)
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	ok := s.InProgress[job]
	if ok {
		s.Lock.RUnlock()
		res := make(chan int)
		defer close(res)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], res)
		s.Lock.Unlock()
		fmt.Printf("Waiting for response job: %d\n", job)
		response := <-res
		fmt.Printf("Response done, response: %d\n", response)
		return
	}
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculating fibonacci for %d\n", job)
	result := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendingWorkers, ok := s.IsPending[job]
	s.Lock.RUnlock()

	if ok {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result
		}
		fmt.Printf("Result send - all pending workers are ready, job: %d\n", job)
	}

	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: map[int]bool{},
		IsPending:  map[int][]chan int{},
		Lock:       sync.RWMutex{},
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}
