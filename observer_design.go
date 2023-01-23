package main

import "fmt"

type Topic interface {
	Register(observer Observer)
}

type Observer interface {
	GetId() string
	UpdateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is now available\n", i.name)
	i.available = true
	i.Broadcast()
}

func (i *Item) Register(ob Observer) {
	i.observers = append(i.observers, ob)
}

func (i *Item) Broadcast() {
	for _, observer := range i.observers {
		observer.UpdateValue(i.name)
	}
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) GetId() string {
	return ec.id
}

func (ec *EmailClient) UpdateValue(name string) {
	fmt.Printf("Sending Email: Item %s is now available from client %s\n", name, ec.id)
}
func main() {
	nvidiaItem := NewItem("Nvidia RTX 3080")
	firstObserver := &EmailClient{id: "12ab"}

	secondObserver := &EmailClient{id: "34cd"}

	nvidiaItem.Register(firstObserver)
	nvidiaItem.Register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
