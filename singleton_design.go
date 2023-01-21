package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) GetConnection() {
	fmt.Println("Getting connection")
	time.Sleep(2 * time.Second)
	fmt.Println("Got connection")
}

var db *Database
var lock sync.Mutex

func GetDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		db = &Database{}
		db.GetConnection()
	} else {
		fmt.Println("Database instance already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}
	wg.Wait()
}
